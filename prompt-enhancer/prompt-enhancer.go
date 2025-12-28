// prompt-enhancer is a Go implementation of the prompt enhancement functionality
// originally written in Bash. It provides enhanced prompts for various opencode commands
// by analyzing user input and generating structured, detailed prompts.

// go build -o ~/oc/prompt-enhancer/prompt-enhancer ~/oc/prompt-enhancer && DEBUG__SKIP_ENHANCING=1 DEBUG__LOG_ARGUMENTS=1 ~/oc/prompt-enhancer/prompt-enhancer init Some new project.

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Configuration constants
const (
	defaultEnhancerModel = "anthropic/claude-sonnet-4-5" // "zai-coding-plan/glm-4.6" // "github-copilot/claude-sonnet-4.5" "opencode/big-pickle" "github-copilot/gpt-4.1"
	preludePath          = "~/.config/opencode/prompt-enhancer/md/prelude.md"
	rfc2119Path          = "~/.config/opencode/prompt-enhancer/md/rfc2119.md"
	epiloguePath         = "~/.config/opencode/prompt-enhancer/md/epilogue.md"
	questionTrailerPath  = "~/.config/opencode/prompt-enhancer/md/question-reminder.md"
	logFileName          = "log.md"
	// Semaphore file name used to signal that prompt-enhancer is running
	// This prevents other plugins (like terminal-bell) from taking action during enhancement
	semaphoreFileName = "prompt-enhancer-semaphore"
)

// Command type enumeration
type CommandType string

const (
	CommandInit        CommandType = "begin"
	CommandFeature     CommandType = "feature"
	CommandChange      CommandType = "change"
	CommandRefactoring CommandType = "refactoring"
	CommandTests       CommandType = "tests"
	CommandBugFix      CommandType = "bugfix"
	CommandQuestion    CommandType = "question"
	CommandBare        CommandType = "bare"
)

// PromptConfig holds configuration for different command types
type PromptConfig struct {
	TitleHeading                string
	Object                      string
	IndependentClauseFragment   string
	Action                      string
	Punctuation                 string
	SupplementalPhrase          string
	DefiniteObjectFragment      string
	DemonstrativeObjectFragment string
	EpilogueIndependentClause   string
	TrailerPath                 string
}

// Debug configuration
type DebugConfig struct {
	LogArguments   bool
	SkipEnhancing  bool
	OutputFences   bool
	IncludeRFC2119 int
}

// Application state
type App struct {
	args            []string
	debug           DebugConfig
	config          PromptConfig
	model           string
	userPrompt      string
	promptObjectArg string
}

func main() {
	app := &App{
		args: os.Args[1:], // Skip program name
		debug: DebugConfig{
			LogArguments:   os.Getenv("DEBUG__LOG_ARGUMENTS") == "1",
			SkipEnhancing:  os.Getenv("DEBUG__SKIP_ENHANCING") == "1",
			OutputFences:   os.Getenv("DEBUG__OUTPUT_FENCES") == "1",
			IncludeRFC2119: getRFC2119Level(),
		},
	}

	// Create semaphore file immediately upon startup to signal that the helper is running
	// This prevents other plugins from taking action during enhancement operations
	if err := createSemaphoreFile(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to create semaphore file: %v\n", err)
		// Continue anyway - semaphore is best-effort and shouldn't prevent operation
	}

	// Ensure semaphore file is cleaned up on exit (both normal and error exits)
	defer func() {
		if err := removeSemaphoreFile(); err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to remove semaphore file: %v\n", err)
		}
	}()

	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// getRFC2119Level returns the RFC2119 inclusion level from environment
func getRFC2119Level() int {
	val := os.Getenv("RFC2119")
	if val == "1" || val == "2" {
		level, _ := strconv.Atoi(val)
		return level
	}
	return 0
}

// run executes the main application logic
func (app *App) run() error {
	if len(app.args) == 0 {
		return fmt.Errorf("no arguments provided")
	}

	// Parse arguments
	app.parseArguments()

	// Log the invocation
	_ = app.logInvocation()

	// Log debug information if requested
	if app.debug.LogArguments {
		app.logDebugInfo()
	}

	// Validate arguments
	if err := app.validateArguments(); err != nil {
		fmt.Println(err)
		return nil // Don't exit with error code, just print the message
	}

	// Generate enhanced prompt
	enhanced, err := app.generateEnhancedPrompt()
	if err != nil {
		return err
	}

	// Output the result
	if app.debug.OutputFences {
		fmt.Println("BEGIN")
	}
	fmt.Print(enhanced)
	if app.debug.OutputFences {
		fmt.Println("END")
	}

	return nil
}

// parseArguments handles command line argument parsing
func (app *App) parseArguments() {
	argsStr := strings.Join(app.args, " ")

	// Split into PROMPT_OBJECT_ARG and REST
	if strings.Contains(argsStr, " ") {
		parts := strings.SplitN(argsStr, " ", 2)
		app.promptObjectArg = parts[0]

		// Handle command aliases - map 'ask' to 'question'
		if app.promptObjectArg == "ask" {
			app.promptObjectArg = "question"
		}
		rest := parts[1]

		// Parse the rest to handle optional RFC argument
		app.parseRestArguments(rest)
	} else {
		app.promptObjectArg = argsStr

		// Handle command aliases - map 'ask' to 'question'
		if app.promptObjectArg == "ask" {
			app.promptObjectArg = "question"
		}

		app.model = defaultEnhancerModel
		app.userPrompt = ""
	}

	// Configure based on command type
	app.configureForCommandType()
}

// parseRestArguments handles the parsing of arguments after the command type,
// including optional model, RFC argument, and prompt text
func (app *App) parseRestArguments(rest string) {
	restParts := strings.Fields(rest)

	if len(restParts) == 0 {
		app.model = defaultEnhancerModel
		app.userPrompt = ""
		return
	}

	if len(restParts) == 1 {
		// Only one argument - could be model or prompt
		if app.isValidModel(restParts[0]) {
			app.model = restParts[0]
			app.userPrompt = ""
		} else {
			app.model = defaultEnhancerModel
			app.userPrompt = restParts[0]
		}
		return
	}

	// Multiple arguments - need to check for RFC argument
	var rfcArg string
	var rfcIndex = -1

	// Find RFC argument if present
	for i, part := range restParts {
		if part == "rfc" || part == "RFC" {
			rfcIndex = i
			rfcArg = part
			break
		}
	}

	if rfcIndex == -1 {
		// No RFC argument - treat as model + prompt
		potentialModel := restParts[0]
		if app.isValidModel(potentialModel) {
			app.model = potentialModel
			app.userPrompt = strings.Join(restParts[1:], " ")
		} else {
			app.model = defaultEnhancerModel
			app.userPrompt = strings.Join(restParts, " ")
		}
		return
	}

	// RFC argument found - determine model and prompt
	if rfcIndex == 0 {
		// RFC is first argument - no model specified
		app.model = defaultEnhancerModel
		app.userPrompt = strings.Join(restParts[1:], " ")
	} else if rfcIndex == 1 {
		// RFC is second argument - first is model
		potentialModel := restParts[0]
		if app.isValidModel(potentialModel) {
			app.model = potentialModel
		} else {
			app.model = defaultEnhancerModel
			// If the first argument isn't a valid model, treat everything as prompt
			app.userPrompt = strings.Join(restParts, " ")
			return
		}
		app.userPrompt = strings.Join(restParts[2:], " ")
	} else {
		// RFC is beyond second position - treat everything before as part of prompt
		app.model = defaultEnhancerModel
		app.userPrompt = strings.Join(restParts, " ")
		return
	}

	// Set RFC2119 level based on argument
	if rfcArg == "rfc" {
		app.debug.IncludeRFC2119 = 1
	} else if rfcArg == "RFC" {
		app.debug.IncludeRFC2119 = 2 // Full RFC content for capitalized RFC
	}
}

// isValidModel checks if a model name is valid using opencode models command
func (app *App) isValidModel(model string) bool {
	cmd := exec.Command("opencode", "models")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), model)
}

// configureForCommandType sets up configuration based on the command type
func (app *App) configureForCommandType() {
	commandType := CommandType(app.promptObjectArg)

	switch commandType {
	case CommandInit:
		app.config = PromptConfig{
			TitleHeading:              "Request to Initialize New Project",
			Object:                    "project",
			IndependentClauseFragment: "the skeleton for this new project",
		}
	case CommandFeature:
		app.config = PromptConfig{
			TitleHeading:              "New Feature Request",
			Object:                    "feature",
			IndependentClauseFragment: "this new feature",
		}
	case CommandChange:
		app.config = PromptConfig{
			TitleHeading:              "New Change Request",
			Object:                    "change",
			IndependentClauseFragment: "this new change to the code's behaviour",
		}
	case CommandRefactoring:
		app.config = PromptConfig{
			TitleHeading:              "New Refactoring Request Specification",
			Object:                    "refactoring",
			IndependentClauseFragment: "this refactoring of the code",
		}
	case CommandTests:
		app.config = PromptConfig{
			TitleHeading:                "New Test Request",
			Object:                      "test(s)",
			IndependentClauseFragment:   "test(s) for these parts of the codebase",
			SupplementalPhrase:          "Avoid writing pointless tests that simply test whether simple constant(s) have expected value(s): focus on testing the BEHAVIOUR of the code.",
			DemonstrativeObjectFragment: "these test(s)",
		}
	case CommandBugFix:
		app.config = PromptConfig{
			TitleHeading:              "Critical Bug Fix Request",
			Object:                    "problem",
			IndependentClauseFragment: "a fix for this bug in the code",
			EpilogueIndependentClause: "analyzing the problem thoroughly and diagnosing its root cause",
		}
	case CommandQuestion:
		app.config = PromptConfig{
			TitleHeading:              "Do not edit the code! Just answer this question",
			Object:                    "question",
			IndependentClauseFragment: "this question",
			Action:                    "Do not edit the code, just answer",
			Punctuation:               "?",
			TrailerPath:               questionTrailerPath,
		}
	case CommandBare:
		app.config = PromptConfig{
			TitleHeading:              "ENHANCED_PROMPT_TITLE_HEADING",
			Object:                    "BARE",
			IndependentClauseFragment: "ENHANCER_INDEPENDANT_CLAUSE_FRAGMENT",
			EpilogueIndependentClause: "EPILOGUE_INDEPENDANT_CLAUSE",
		}
	}

	// Set defaults for optional fields
	if app.config.Punctuation == "" {
		app.config.Punctuation = "."
	}
	if app.config.Action == "" {
		app.config.Action = "Implement"
	}
	if app.config.DefiniteObjectFragment == "" {
		app.config.DefiniteObjectFragment = "the " + app.config.Object
	}
	if app.config.DemonstrativeObjectFragment == "" {
		app.config.DemonstrativeObjectFragment = "this " + app.config.Object
	}
	if app.config.EpilogueIndependentClause == "" {
		app.config.EpilogueIndependentClause = "thinking the implementation of " + app.config.DemonstrativeObjectFragment + " through thoroughly"
	}
}

// logDebugInfo prints debug information about argument parsing
func (app *App) logDebugInfo() {
	fmt.Println("# Debug Information:")
	fmt.Printf("app.args='%s'\n", strings.Join(app.args, " "))
	fmt.Printf("app.promptObjectArg='%s'\n", app.promptObjectArg)
	fmt.Printf("app.model='%s'\n", app.model)
	fmt.Printf("app.userPrompt='%s'\n", app.userPrompt)
	fmt.Printf("app.config.TitleHeading='%s'\n", app.config.TitleHeading)
	fmt.Printf("app.config.EpilogueIndependentClause='%s'\n", app.config.EpilogueIndependentClause)
	fmt.Printf("app.config.DefiniteObjectFragment='%s'\n", app.config.DefiniteObjectFragment)
	fmt.Printf("app.config.DemonstrativeObjectFragment='%s'\n", app.config.DemonstrativeObjectFragment)

	for i, arg := range app.args {
		fmt.Printf("%d: %s\n", i, arg)
	}
	fmt.Println()
}

// validateArguments checks if the parsed arguments are valid
func (app *App) validateArguments() error {
	// Only check for empty TitleHeading - "ENHANCED_PROMPT_TITLE_HEADING" is valid for BARE commands
	if app.config.TitleHeading == "" {
		return fmt.Errorf("FATAL ERROR: The user has provided bad arguments to the command they tried to use, and as a result this prompt's content has been corrupted. Please remind the user that this command's first argument should be a model listed by the `opencode models` command and the remainder must constitute a non-empty string.")
	}

	if app.userPrompt == "" {
		return fmt.Errorf("FATAL ERROR: The user has provided bad arguments to the command they tried to use, and as a result this prompt's content has been corrupted. Please remind the user that this command's first argument should be a model listed by the `opencode models` command and the remainder must constitute a non-empty string.")
	}

	return nil
}

// hasTerminalPunctuation checks if a string ends with terminal punctuation,
// ignoring trailing whitespace. Terminal punctuation includes: . ? ! …
func hasTerminalPunctuation(s string) bool {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return false
	}

	lastChar := trimmed[len(trimmed)-1:]
	// Check for single-byte punctuation: . ? !
	if lastChar == "." || lastChar == "?" || lastChar == "!" {
		return true
	}

	// Check for Unicode ellipsis (U+2026: …)
	if strings.HasSuffix(trimmed, "…") {
		return true
	}

	return false
}

// generateEnhancedPrompt creates the enhanced prompt using the opencode agent
func (app *App) generateEnhancedPrompt() (string, error) {
	// Read the prelude
	prelude, err := app.readPrelude()
	if err != nil {
		return "ERROR: Failed to read prelude!", err
	}

	// Build the base prompt
	var basePrompt bytes.Buffer
	basePrompt.WriteString(prelude)
	basePrompt.WriteString("\n## Original Prompt\n\n")

	if app.config.Object != "BARE" {
		basePrompt.WriteString(app.config.Action + " " + app.config.IndependentClauseFragment + ": ")
	}

	// Conditionally append punctuation only if userPrompt doesn't already end with terminal punctuation
	if hasTerminalPunctuation(app.userPrompt) {
		basePrompt.WriteString(app.userPrompt)
	} else {
		basePrompt.WriteString(app.userPrompt + app.config.Punctuation)
	}

	if app.config.SupplementalPhrase != "" {
		basePrompt.WriteString("\n" + app.config.SupplementalPhrase + "\n")
	}

	var enhanced string
	if app.debug.SkipEnhancing {
		enhanced = basePrompt.String() + "\n\nDEBUG__SKIP_ENHANCING is set, this is dummy data!"
	} else {
		// Run the enhancement through opencode
		enhanced, err = app.runOpenCodeEnhancement(basePrompt.String())
		if err != nil {
			return "", err
		}
	}

	// Apply text processing
	enhanced = app.processText(enhanced)

	// Add RFC2119 compliance phrase if needed
	if app.debug.IncludeRFC2119 >= 1 && app.config.Object != "BARE" && app.config.Object != "question" {
		enhanced = "The keywords \"MUST\", \"MUST NOT\", \"REQUIRED\", \"SHALL\", \"SHALL NOT\", \"SHOULD\", \"SHOULD NOT\", \"RECOMMENDED\",  \"MAY\", and \"OPTIONAL\" in this document are to be interpreted as described in RFC 2119.\n\n" + enhanced
	}

	// Add title heading
	if app.config.Object != "BARE" {
		enhanced = "# " + app.config.TitleHeading + ":\n\n" + enhanced
	}

	// Add boilerplate for non-BARE and non-question commands
	if app.config.Object != "BARE" && app.config.Object != "question" {
		enhanced = app.addBoilerplate(enhanced)
	}

	// Append question reminder for question commands
	if app.config.TrailerPath != "" {
		expandedPath := expandPath(app.config.TrailerPath)
		if content, err := os.ReadFile(expandedPath); err == nil {
			enhanced = enhanced + "\n\n" + string(content)
		}
	}

	return enhanced, nil
}

// readPrelude reads the prompt enhancer prelude file
func (app *App) readPrelude() (string, error) {
	path := expandPath(preludePath)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read prelude file %s: %w", path, err)
	}
	return string(content), nil
}

// expandPath expands ~ to the user's home directory
func expandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		home, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(home, path[2:])
		}
	}
	return path
}

// addFileContentToBuffer reads a file and appends its content to a buffer
// with optional pre-content separator and trailing newline
func addFileContentToBuffer(buf *bytes.Buffer, filePath, preSeparator string) {
	expandedPath := expandPath(filePath)
	if content, err := os.ReadFile(expandedPath); err == nil {
		if preSeparator != "" {
			buf.WriteString(preSeparator)
		}
		buf.Write(content)
		buf.WriteString("\n")
	}
}

// substituteEpiloguePlaceholders replaces placeholders in epilogue content
// with dynamic values from the app configuration
func (app *App) substituteEpiloguePlaceholders(content string) string {
	content = strings.ReplaceAll(content, "{EPILOGUE_INDEPENDENT_CLAUSE}", app.config.EpilogueIndependentClause)
	content = strings.ReplaceAll(content, "{DEMONSTRATIVE_OBJECT_FRAGMENT}", app.config.DemonstrativeObjectFragment)
	return content
}

// getHardcodedEpilogue returns the hardcoded epilogue text as a fallback
// when the epilogue file cannot be read
func (app *App) getHardcodedEpilogue() string {
	var result bytes.Buffer
	result.WriteString("## IMPORTANT: Employ our standard pracices to maximize the odds of successful implementation!\n\n")
	result.WriteString("So long as you proceed systematically, work hard, and adhere to our standard practices, ")
	result.WriteString("your successful completion of the task is as good as guaranteed! Remember:\n")
	result.WriteString("- Start by " + app.config.EpilogueIndependentClause + ". ")
	result.WriteString("Then, you MUST break the implementation of " + app.config.DemonstrativeObjectFragment + " ")
	result.WriteString("down into small steps to produce a detailed, ")
	result.WriteString("step-by-step plan that you will use to implement " + app.config.DemonstrativeObjectFragment + ". ")
	result.WriteString("Group the plan's steps into \"phases\".\n")
	result.WriteString("The code MUST build correctly and any tests in the project MUST pass afterwards.\n")
	result.WriteString("- Next, write the plan into an appropriately named new Markdown file in the project's ")
	result.WriteString("./plans directory which includes checkboxes in which to mark the completion of each step.\n")
	result.WriteString("- Proceed to systematically implement the plan that you just wrote in the Markdown file. ")
	result.WriteString("You MUST check off each step you've completed in the Markdown file immediately as you complete it, ")
	result.WriteString("you MAY NOT proceed to the next step until you have checked off the current step.\n")
	result.WriteString("- Follow through and finish the job: you MUST continue complete the task! ")
	result.WriteString("Keep working until every step in the Markdown file has been checked off and the entire plan has")
	result.WriteString("been completed. ")
	result.WriteString("The code MUST build correctly and any tests in the project MUST pass afterwards.")
	return result.String()
}

// runOpenCodeEnhancement runs the opencode enhancement process
func (app *App) runOpenCodeEnhancement(prompt string) (string, error) {
	cmd := exec.Command("opencode", "--model", app.model, "run")
	cmd.Stdin = strings.NewReader(prompt)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil && stderr.Len() > 0 {
		return "", fmt.Errorf("opencode command failed: %s", stderr.String())
	}

	return stdout.String(), nil
}

// processText applies text processing similar to the sed commands in the Bash version
func (app *App) processText(text string) string {
	// Remove "Enhanced Prompt" section headers (case insensitive, line by line)
	lines := strings.Split(text, "\n")
	var filteredLines []string
	for _, line := range lines {
		re := regexp.MustCompile(`(?i)^#+\s*Enhanced.*`)
		if !re.MatchString(line) {
			filteredLines = append(filteredLines, line)
		}
	}
	text = strings.Join(filteredLines, "\n")

	// Deepen all headings by one (add extra #)
	lines = strings.Split(text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			lines[i] = "#" + line
		}
	}
	text = strings.Join(lines, "\n")

	// Compress consecutive blank lines
	blankLineRe := regexp.MustCompile(`\n{3,}`)
	text = blankLineRe.ReplaceAllString(text, "\n\n")

	// Remove leading newlines
	text = strings.TrimLeft(text, "\n")

	return text
}

// addBoilerplate adds the standard practices boilerplate text
func (app *App) addBoilerplate(text string) string {
	var result bytes.Buffer

	// Add RFC2119 content if level >= 2
	if app.debug.IncludeRFC2119 >= 2 {
		addFileContentToBuffer(&result, rfc2119Path, "")
	}

	result.WriteString(text)
	result.WriteString("\n\n")

	// Try to load epilogue from file, fall back to hardcoded version
	expandedPath := expandPath(epiloguePath)
	if content, err := os.ReadFile(expandedPath); err == nil {
		// File exists, substitute placeholders and use it
		epilogueText := app.substituteEpiloguePlaceholders(string(content))
		result.WriteString(epilogueText)
	} else {
		// File doesn't exist, use hardcoded fallback
		result.WriteString(app.getHardcodedEpilogue())
	}

	return result.String()
}

// getLogFilePath returns the absolute path to the log file
func getLogFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, ".config/opencode/prompt-enhancer", logFileName)
}

// getSemaphoreFilePath returns the absolute path to the semaphore file
func getSemaphoreFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	return filepath.Join(home, ".config/opencode/prompt-enhancer", semaphoreFileName)
}

// createSemaphoreFile creates the semaphore file to signal that the helper is running
func createSemaphoreFile() error {
	semaphorePath := getSemaphoreFilePath()

	// Check if semaphore file already exists (stale from previous crash)
	if _, err := os.Stat(semaphorePath); err == nil {
		// File exists, remove it first to ensure clean state
		if removeErr := os.Remove(semaphorePath); removeErr != nil {
			return fmt.Errorf("failed to remove existing semaphore file: %w", removeErr)
		}
	}

	file, err := os.Create(semaphorePath)
	if err != nil {
		return fmt.Errorf("failed to create semaphore file: %w", err)
	}
	defer file.Close()

	// Ensure the file is actually written to disk
	if err := file.Sync(); err != nil {
		return fmt.Errorf("failed to sync semaphore file: %w", err)
	}

	return nil
}

// removeSemaphoreFile removes the semaphore file to signal that the helper is no longer running
func removeSemaphoreFile() error {
	semaphorePath := getSemaphoreFilePath()
	err := os.Remove(semaphorePath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove semaphore file: %w", err)
	}

	return nil
}

// formatLogEntry formats the arguments as a markdown list item with timestamp
func formatLogEntry(args []string) string {
	timestamp := time.Now().Format("15:04:05")
	return fmt.Sprintf("- [%s] %s", timestamp, strings.Join(args, " "))
}

// logInvocationWithPath appends the current invocation to the log file at the given path
func (app *App) logInvocationWithPath(logPath string) error {
	entry := formatLogEntry(app.args) + "\n"

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	if err != nil {
		return nil
	}

	return nil
}

// logInvocation appends the current invocation to the log file
func (app *App) logInvocation() error {
	return app.logInvocationWithPath(getLogFilePath())
}
