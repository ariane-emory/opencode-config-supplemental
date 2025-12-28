package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestArgumentParsing(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		expectedObject string
		expectedModel  string
		expectedPrompt string
	}{
		{
			name:           "Simple feature command",
			args:           []string{"feature", "Add", "a", "new", "feature"},
			expectedObject: "feature",
			expectedModel:  defaultEnhancerModel,
			expectedPrompt: "Add a new feature",
		},
		{
			name:           "question command without model",
			args:           []string{"question", "What is testing?"},
			expectedObject: "question",
			expectedModel:  defaultEnhancerModel,
			expectedPrompt: "What is testing?",
		},
		{
			name:           "ask command alias without model",
			args:           []string{"ask", "What is testing?"},
			expectedObject: "question",
			expectedModel:  defaultEnhancerModel,
			expectedPrompt: "What is testing?",
		},
		{
			name:           "question command with model",
			args:           []string{"question", "claude-3", "What is Go?"},
			expectedObject: "question",
			expectedModel:  "claude-3",
			expectedPrompt: "What is Go?",
		},
		{
			name:           "ask command alias with model",
			args:           []string{"ask", "claude-3", "What is Go?"},
			expectedObject: "question",
			expectedModel:  "claude-3",
			expectedPrompt: "What is Go?",
		},
		{
			name:           "Tests command",
			args:           []string{"tests", "Create", "tests"},
			expectedObject: "tests",
			expectedModel:  defaultEnhancerModel,
			expectedPrompt: "Create tests",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				args: tt.args,
				debug: DebugConfig{
					LogArguments:   false,
					SkipEnhancing:  true, // Skip actual enhancement for testing
					OutputFences:   false,
					IncludeRFC2119: 0,
				},
			}

			app.parseArguments()

			if app.promptObjectArg != tt.expectedObject {
				t.Errorf("Expected promptObjectArg %q, got %q", tt.expectedObject, app.promptObjectArg)
			}

			if app.model != tt.expectedModel {
				t.Errorf("Expected model %q, got %q", tt.expectedModel, app.model)
			}

			if app.userPrompt != tt.expectedPrompt {
				t.Errorf("Expected userPrompt %q, got %q", tt.expectedPrompt, app.userPrompt)
			}
		})
	}
}

func TestCommandTypeConfiguration(t *testing.T) {
	tests := []struct {
		name                string
		commandType         CommandType
		expectedTitle       string
		expectedObject      string
		expectedAction      string
		expectedPunctuation string
	}{
		{
			name:                "Feature command",
			commandType:         CommandFeature,
			expectedTitle:       "New Feature Request",
			expectedObject:      "feature",
			expectedAction:      "Implement",
			expectedPunctuation: ".",
		},
		{
			name:                "Question command",
			commandType:         CommandQuestion,
			expectedTitle:       "Do not edit the code! Just answer this question",
			expectedObject:      "question",
			expectedAction:      "Do not edit the code, just answer",
			expectedPunctuation: "?",
		},
		{
			name:                "Tests command",
			commandType:         CommandTests,
			expectedTitle:       "New Test Request",
			expectedObject:      "test(s)",
			expectedAction:      "Implement",
			expectedPunctuation: ".",
		},
		{
			name:                "BugFix command",
			commandType:         CommandBugFix,
			expectedTitle:       "Critical Bug Fix Request",
			expectedObject:      "problem",
			expectedAction:      "Implement",
			expectedPunctuation: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				promptObjectArg: string(tt.commandType),
			}

			app.configureForCommandType()

			if app.config.TitleHeading != tt.expectedTitle {
				t.Errorf("Expected TitleHeading %q, got %q", tt.expectedTitle, app.config.TitleHeading)
			}

			if app.config.Object != tt.expectedObject {
				t.Errorf("Expected Object %q, got %q", tt.expectedObject, app.config.Object)
			}

			if app.config.Action != tt.expectedAction {
				t.Errorf("Expected Action %q, got %q", tt.expectedAction, app.config.Action)
			}

			if app.config.Punctuation != tt.expectedPunctuation {
				t.Errorf("Expected Punctuation %q, got %q", tt.expectedPunctuation, app.config.Punctuation)
			}
		})
	}
}

func TestTextProcessing(t *testing.T) {
	app := &App{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Remove Enhanced Prompt headers",
			input:    "# Enhanced Prompt\nSome content\n## Enhanced Prompt\nMore content",
			expected: "Some content\nMore content",
		},
		{
			name:     "Deepen headings",
			input:    "# Title\n## Subtitle\nContent",
			expected: "## Title\n### Subtitle\nContent",
		},
		{
			name:     "Compress blank lines",
			input:    "Content\n\n\n\nMore content",
			expected: "Content\n\nMore content",
		},
		{
			name:     "Remove leading newlines",
			input:    "\n\n\nContent",
			expected: "Content",
		},
		{
			name:     "Complex processing",
			input:    "\n\n# Enhanced Prompt\n\n# Title\n\n\n\nContent",
			expected: "## Title\n\nContent",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := app.processText(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestExpandPath(t *testing.T) {
	// Test with ~ expansion
	input := "~/test/file.txt"
	result := expandPath(input)

	if !strings.Contains(result, "test/file.txt") {
		t.Errorf("Expected expanded path to contain 'test/file.txt', got %q", result)
	}

	// Test with regular path (no expansion)
	input2 := "/absolute/path/file.txt"
	result2 := expandPath(input2)

	if result2 != input2 {
		t.Errorf("Expected unchanged path %q, got %q", input2, result2)
	}
}

func TestValidateArguments(t *testing.T) {
	tests := []struct {
		name        string
		app         *App
		expectError bool
	}{
		{
			name: "Valid arguments",
			app: &App{
				config: PromptConfig{
					TitleHeading: "Valid Title",
				},
				userPrompt: "Valid prompt",
			},
			expectError: false,
		},
		{
			name: "Empty title heading",
			app: &App{
				config: PromptConfig{
					TitleHeading: "",
				},
				userPrompt: "Valid prompt",
			},
			expectError: true,
		},
		{
			name: "BARE command title",
			app: &App{
				config: PromptConfig{
					TitleHeading: "ENHANCED_PROMPT_TITLE_HEADING",
					Object:       "BARE",
				},
				userPrompt: "Valid prompt",
			},
			expectError: false,
		},
		{
			name: "Empty user prompt",
			app: &App{
				config: PromptConfig{
					TitleHeading: "Valid Title",
				},
				userPrompt: "",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.app.validateArguments()
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}

func TestGetRFC2119Level(t *testing.T) {
	// Test default value
	os.Unsetenv("RFC2119")
	level := getRFC2119Level()
	if level != 0 {
		t.Errorf("Expected default level 0, got %d", level)
	}

	// Test with RFC2119=1
	os.Setenv("RFC2119", "1")
	level = getRFC2119Level()
	if level != 1 {
		t.Errorf("Expected level 1, got %d", level)
	}

	// Test with RFC2119=2
	os.Setenv("RFC2119", "2")
	level = getRFC2119Level()
	if level != 2 {
		t.Errorf("Expected level 2, got %d", level)
	}

	// Test with invalid value
	os.Setenv("RFC2119", "invalid")
	level = getRFC2119Level()
	if level != 0 {
		t.Errorf("Expected level 0 for invalid value, got %d", level)
	}

	// Cleanup
	os.Unsetenv("RFC2119")
}

func TestReadPrelude(t *testing.T) {
	// Create a temporary prelude file
	tempDir := t.TempDir()
	preludeContent := "## Test Prelude\nThis is a test prelude file."
	preludePath := filepath.Join(tempDir, "prelude.md")

	err := os.WriteFile(preludePath, []byte(preludeContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test prelude file: %v", err)
	}

	// Test reading the file
	content, err := os.ReadFile(preludePath)
	if err != nil {
		t.Fatalf("Failed to read test prelude file: %v", err)
	}

	if string(content) != preludeContent {
		t.Errorf("Expected content %q, got %q", preludeContent, string(content))
	}
}

func TestAddBoilerplate(t *testing.T) {
	app := &App{
		config: PromptConfig{
			Object:                      "feature",
			DemonstrativeObjectFragment: "this feature",
			EpilogueIndependentClause:   "thinking through the implementation",
		},
		debug: DebugConfig{
			IncludeRFC2119: 0,
		},
	}

	input := "# Test Title\nSome content"
	result := app.addBoilerplate(input)

	// Check that the original content is preserved
	if !strings.Contains(result, input) {
		t.Errorf("Expected original content to be preserved in result")
	}

	// Check that boilerplate sections are added
	expectedSections := []string{
		"## IMPORTANT: Employ our standard practices",
		"Start by thinking through the implementation",
		"break the implementation of this feature down into small steps",
		"write the plan into an appropriately named new Markdown file",
		"check off each step you've completed",
	}

	for _, section := range expectedSections {
		if !strings.Contains(result, section) {
			t.Errorf("Expected boilerplate to contain %q", section)
		}
	}
}

func TestRFCArgumentParsing(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		expectedModel    string
		expectedPrompt   string
		expectedRFCLevel int
	}{
		{
			name:             "No RFC argument",
			args:             []string{"feature", "Add", "a", "feature"},
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "Add a feature",
			expectedRFCLevel: 0,
		},
		{
			name:             "Lowercase RFC argument without model",
			args:             []string{"feature", "rfc", "Add", "feature"},
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "Add feature",
			expectedRFCLevel: 1,
		},
		{
			name:             "Uppercase RFC argument without model",
			args:             []string{"feature", "RFC", "Add", "feature"},
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "Add feature",
			expectedRFCLevel: 2,
		},
		{
			name:             "Lowercase RFC argument with model",
			args:             []string{"feature", "qwen3-coder-flash", "rfc", "Add", "feature"},
			expectedModel:    "qwen3-coder-flash",
			expectedPrompt:   "Add feature",
			expectedRFCLevel: 1,
		},
		{
			name:             "Uppercase RFC argument with model",
			args:             []string{"feature", "qwen3-coder-flash", "RFC", "Add", "feature"},
			expectedModel:    "qwen3-coder-flash",
			expectedPrompt:   "Add feature",
			expectedRFCLevel: 2,
		},
		{
			name:             "ask command alias without model",
			args:             []string{"ask", "What", "is", "testing?"},
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "What is testing?",
			expectedRFCLevel: 0,
		},
		{
			name:             "ask command alias with model",
			args:             []string{"ask", "claude-3", "What", "is", "Go?"},
			expectedModel:    "claude-3",
			expectedPrompt:   "What is Go?",
			expectedRFCLevel: 0,
		},
		{
			name:             "ask command alias with rfc",
			args:             []string{"ask", "rfc", "What", "is", "testing?"},
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "What is testing?",
			expectedRFCLevel: 1,
		},
		{
			name:             "ask command alias with model and rfc",
			args:             []string{"ask", "gpt-4", "rfc", "What", "is", "Go?"},
			expectedModel:    "gpt-4",
			expectedPrompt:   "What is Go?",
			expectedRFCLevel: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				args: tt.args,
				debug: DebugConfig{
					LogArguments:   false,
					SkipEnhancing:  true,
					OutputFences:   false,
					IncludeRFC2119: 0,
				},
			}

			app.parseArguments()

			if app.model != tt.expectedModel {
				t.Errorf("Expected model %q, got %q", tt.expectedModel, app.model)
			}

			if app.userPrompt != tt.expectedPrompt {
				t.Errorf("Expected userPrompt %q, got %q", tt.expectedPrompt, app.userPrompt)
			}

			if app.debug.IncludeRFC2119 != tt.expectedRFCLevel {
				t.Errorf("Expected RFC level %d, got %d", tt.expectedRFCLevel, app.debug.IncludeRFC2119)
			}
		})
	}
}

func TestParseRestArguments(t *testing.T) {
	tests := []struct {
		name             string
		rest             string
		expectedModel    string
		expectedPrompt   string
		expectedRFCLevel int
	}{
		{
			name:             "Single argument - invalid model",
			rest:             "prompt",
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "prompt",
			expectedRFCLevel: 0,
		},
		{
			name:             "Two arguments - model + prompt",
			rest:             "qwen3-coder-flash prompt text",
			expectedModel:    "qwen3-coder-flash",
			expectedPrompt:   "prompt text",
			expectedRFCLevel: 0,
		},
		{
			name:             "RFC as first argument",
			rest:             "rfc prompt text",
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "prompt text",
			expectedRFCLevel: 1,
		},
		{
			name:             "RFC as second argument with valid model",
			rest:             "qwen3-coder-flash rfc prompt text",
			expectedModel:    "qwen3-coder-flash",
			expectedPrompt:   "prompt text",
			expectedRFCLevel: 1,
		},
		{
			name:             "RFC as second argument with invalid model",
			rest:             "invalid-model rfc prompt text",
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "invalid-model rfc prompt text",
			expectedRFCLevel: 0,
		},
		{
			name:             "RFC beyond second position",
			rest:             "some text rfc more text",
			expectedModel:    defaultEnhancerModel,
			expectedPrompt:   "some text rfc more text",
			expectedRFCLevel: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				debug: DebugConfig{
					IncludeRFC2119: 0,
				},
			}

			app.parseRestArguments(tt.rest)

			if app.model != tt.expectedModel {
				t.Errorf("Expected model %q, got %q", tt.expectedModel, app.model)
			}

			if app.userPrompt != tt.expectedPrompt {
				t.Errorf("Expected userPrompt %q, got %q", tt.expectedPrompt, app.userPrompt)
			}

			if app.debug.IncludeRFC2119 != tt.expectedRFCLevel {
				t.Errorf("Expected RFC level %d, got %d", tt.expectedRFCLevel, app.debug.IncludeRFC2119)
			}
		})
	}
}

func TestLogInvocation(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, logFileName)

	app := &App{
		args: []string{"feature", "Add", "new", "feature"},
	}

	err := app.logInvocationWithPath(logPath)
	if err != nil {
		t.Errorf("logInvocation failed: %v", err)
	}

	// Check that log file was created
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	contentStr := string(content)
	expectedArgs := "feature Add new feature"
	if !strings.Contains(contentStr, expectedArgs) {
		t.Errorf("Expected log to contain %q, got %q", expectedArgs, contentStr)
	}

	if !strings.Contains(contentStr, "- [") {
		t.Errorf("Expected log to start with markdown list format, got %q", contentStr)
	}
}

func TestLogInvocationAppend(t *testing.T) {
	tempDir := t.TempDir()
	logPath := filepath.Join(tempDir, logFileName)

	// First invocation
	app1 := &App{
		args: []string{"feature", "First", "feature"},
	}
	err1 := app1.logInvocationWithPath(logPath)
	if err1 != nil {
		t.Errorf("First logInvocation failed: %v", err1)
	}

	// Second invocation
	app2 := &App{
		args: []string{"question", "What", "is", "this?"},
	}
	err2 := app2.logInvocationWithPath(logPath)
	if err2 != nil {
		t.Errorf("Second logInvocation failed: %v", err2)
	}

	// Check that both entries are in the log
	content, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	contentStr := string(content)
	if !strings.Contains(contentStr, "feature First feature") {
		t.Errorf("Expected log to contain first invocation")
	}
	if !strings.Contains(contentStr, "question What is this?") {
		t.Errorf("Expected log to contain second invocation")
	}

	// Count lines to ensure both entries exist
	lines := strings.Split(strings.TrimSpace(contentStr), "\n")
	if len(lines) < 2 {
		t.Errorf("Expected at least 2 log entries, got %d", len(lines))
	}
}

func TestFormatLogEntry(t *testing.T) {
	args := []string{"feature", "Add", "new", "feature"}
	entry := formatLogEntry(args)

	// Check format
	if !strings.HasPrefix(entry, "- [") {
		t.Errorf("Expected entry to start with '- [', got %q", entry)
	}

	if !strings.Contains(entry, "feature Add new feature") {
		t.Errorf("Expected entry to contain 'feature Add new feature', got %q", entry)
	}

	if !strings.Contains(entry, "]") {
		t.Errorf("Expected entry to contain closing bracket, got %q", entry)
	}
}

func TestGetLogFilePath(t *testing.T) {
	logPath := getLogFilePath()

	if !strings.Contains(logPath, "log.md") {
		t.Errorf("Expected log path to contain 'log.md', got %q", logPath)
	}

	if !strings.Contains(logPath, ".config/opencode/prompt-enhancer") {
		t.Errorf("Expected log path to contain '.config/opencode/prompt-enhancer', got %q", logPath)
	}
}

func TestHasTerminalPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Ends with period",
			input:    "This is a sentence.",
			expected: true,
		},
		{
			name:     "Ends with question mark",
			input:    "Is this a question?",
			expected: true,
		},
		{
			name:     "Ends with exclamation",
			input:    "This is exciting!",
			expected: true,
		},
		{
			name:     "Ends with ellipsis",
			input:    "This continues…",
			expected: true,
		},
		{
			name:     "No terminal punctuation",
			input:    "This has no punctuation",
			expected: false,
		},
		{
			name:     "Ends with comma",
			input:    "This ends with a comma,",
			expected: false,
		},
		{
			name:     "Question mark with trailing whitespace",
			input:    "Is this a question?   ",
			expected: true,
		},
		{
			name:     "Period with trailing whitespace",
			input:    "This is a sentence.  \n\t",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
		{
			name:     "Only whitespace",
			input:    "   \n\t  ",
			expected: false,
		},
		{
			name:     "Exclamation with spaces",
			input:    "Wow!     ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasTerminalPunctuation(tt.input)
			if result != tt.expected {
				t.Errorf("hasTerminalPunctuation(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPunctuationHandlingInPrompts(t *testing.T) {
	tests := []struct {
		name              string
		args              []string
		expectedInOutput  string
		notExpectedOutput string
	}{
		{
			name:              "Question command with question mark",
			args:              []string{"question", "How does it work?"},
			expectedInOutput:  "How does it work?",
			notExpectedOutput: "How does it work??",
		},
		{
			name:              "Question command without punctuation",
			args:              []string{"question", "How does it work"},
			expectedInOutput:  "How does it work?",
			notExpectedOutput: "How does it work.",
		},
		{
			name:              "Ask alias with question mark",
			args:              []string{"ask", "What is this?"},
			expectedInOutput:  "What is this?",
			notExpectedOutput: "What is this??",
		},
		{
			name:              "Ask alias without punctuation",
			args:              []string{"ask", "What is this"},
			expectedInOutput:  "What is this?",
			notExpectedOutput: "What is this.",
		},
		{
			name:              "Question with period",
			args:              []string{"question", "Tell me about this."},
			expectedInOutput:  "Tell me about this.",
			notExpectedOutput: "Tell me about this.?",
		},
		{
			name:              "Question with exclamation",
			args:              []string{"question", "Why is this failing!"},
			expectedInOutput:  "Why is this failing!",
			notExpectedOutput: "Why is this failing!?",
		},
		{
			name:              "Question with ellipsis",
			args:              []string{"question", "I wonder why…"},
			expectedInOutput:  "I wonder why…",
			notExpectedOutput: "I wonder why…?",
		},
		{
			name:              "Question with trailing space and question mark",
			args:              []string{"question", "How does it work?   "},
			expectedInOutput:  "How does it work?",
			notExpectedOutput: "How does it work??",
		},
		{
			name:              "Change command with period",
			args:              []string{"change", "Implement this feature."},
			expectedInOutput:  "Implement this feature.",
			notExpectedOutput: "Implement this feature..",
		},
		{
			name:              "Change command without punctuation",
			args:              []string{"change", "Implement this feature"},
			expectedInOutput:  "Implement this feature.",
			notExpectedOutput: "Implement this feature?",
		},
		{
			name:              "Feature command with exclamation",
			args:              []string{"feature", "Add this now!"},
			expectedInOutput:  "Add this now!",
			notExpectedOutput: "Add this now!.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				args: tt.args,
				debug: DebugConfig{
					LogArguments:   false,
					SkipEnhancing:  true,
					OutputFences:   false,
					IncludeRFC2119: 0,
				},
			}

			app.parseArguments()

			if app.userPrompt == "" {
				t.Fatal("userPrompt is empty")
			}

			// Build the prompt section to test punctuation handling
			var basePrompt bytes.Buffer
			if app.config.Object != "BARE" {
				basePrompt.WriteString(app.config.Action + " " + app.config.IndependentClauseFragment + ": ")
			}

			// Apply the same logic as generateEnhancedPrompt
			if hasTerminalPunctuation(app.userPrompt) {
				basePrompt.WriteString(app.userPrompt)
			} else {
				basePrompt.WriteString(app.userPrompt + app.config.Punctuation)
			}

			result := basePrompt.String()

			if !strings.Contains(result, tt.expectedInOutput) {
				t.Errorf("Expected output to contain %q, but got %q", tt.expectedInOutput, result)
			}

			if strings.Contains(result, tt.notExpectedOutput) {
				t.Errorf("Expected output NOT to contain %q, but got %q", tt.notExpectedOutput, result)
			}
		})
	}
}
