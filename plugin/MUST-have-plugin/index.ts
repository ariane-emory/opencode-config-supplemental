/**
 * Replacer Plugin
 *
 * Performs case-insensitive string replacements on user-typed prompts
 * before they're sent to the LLM.
 *
 * Config: ~/.config/opencode/MUST-have-plugin.jsonc (JSONC format)
 * Debug log: /tmp/opencode-MUST-have-plugin-debug.log
 *
 * Features:
 * - Case-insensitive matching with word boundaries
 * - Multi-word phrase support (longest-first matching)
 * - Hot reload: config re-read on every message
 * - Auto-generates RFC2119 defaults if no config exists
 * - JSONC support (comments allowed in config)
 *
 * Scope:
 * - Only replaces in user-typed prompts
 * - Does NOT modify file content attached via @ mentions
 * - Does NOT modify slash command output
 */
import type { Plugin } from "@opencode-ai/plugin"
import { readFileSync, writeFileSync, existsSync, appendFileSync } from "fs"
import { resolve } from "path"

const CONFIG_PATH = resolve(process.env.HOME || "", ".config", "opencode", "MUST-have-plugin.jsonc")
const DEBUG_LOG = "/tmp/opencode-MUST-have-plugin-debug.log"

// Track debug state (loaded from config)
let debugEnabled = false

/**
 * RFC2119 default replacements
 * These keywords are used in technical specifications to indicate requirement levels.
 * See: https://datatracker.ietf.org/doc/html/rfc2119
 */
const RFC2119_DEFAULTS: Record<string, string> = {
  "must": "MUST",
  "must not": "MUST NOT",
  "required": "REQUIRED",
  "shall": "SHALL",
  "shall not": "SHALL NOT",
  "should": "SHOULD",
  "should not": "SHOULD NOT",
  "recommended": "RECOMMENDED",
  "not recommended": "NOT RECOMMENDED",
  "may": "MAY",
  "optional": "OPTIONAL",
}

/**
 * Default config file content (JSONC format)
 */
const DEFAULT_CONFIG = `{
  // Uncomment to enable debug logging (view with: tail -f /tmp/opencode-MUST-have-plugin-debug.log)
  // "debug": true,

  "replacements": {
    "must": "MUST",
    "must not": "MUST NOT",
    "required": "REQUIRED",
    "shall": "SHALL",
    "shall not": "SHALL NOT",
    "should": "SHOULD",
    "should not": "SHOULD NOT",
    "recommended": "RECOMMENDED",
    "not recommended": "NOT RECOMMENDED",
    "may": "MAY",
    "optional": "OPTIONAL"
  }
}
`

/**
 * Write a timestamped message to the debug log file
 */
function log(message: string): void {
  if (!debugEnabled) return
  try {
    const timestamp = new Date().toISOString()
    appendFileSync(DEBUG_LOG, `${timestamp} ${message}\n`)
  } catch {
    // Silently fail - logging should never break the plugin
  }
}

/**
 * Parse JSONC (JSON with Comments) by stripping comments before parsing
 */
function parseJSONC(content: string): any {
  // Strip single-line comments (// ...) - but not inside strings
  // This is a simplified approach that works for typical config files
  const lines = content.split("\n")
  const strippedLines = lines.map((line) => {
    // Find // that's not inside a string
    let inString = false
    let escapeNext = false
    for (let i = 0; i < line.length; i++) {
      const char = line[i]
      if (escapeNext) {
        escapeNext = false
        continue
      }
      if (char === "\\") {
        escapeNext = true
        continue
      }
      if (char === '"') {
        inString = !inString
        continue
      }
      if (!inString && char === "/" && line[i + 1] === "/") {
        return line.substring(0, i)
      }
    }
    return line
  })

  const stripped = strippedLines.join("\n")
  // Also strip multi-line comments /* ... */
  const noMultiLine = stripped.replace(/\/\*[\s\S]*?\*\//g, "")

  return JSON.parse(noMultiLine)
}

/**
 * Create default config file with RFC2119 keywords if it doesn't exist
 */
function ensureConfigExists(): void {
  if (!existsSync(CONFIG_PATH)) {
    try {
      writeFileSync(CONFIG_PATH, DEFAULT_CONFIG, "utf-8")
      log(`[INIT] Created default config at ${CONFIG_PATH}`)
    } catch (error) {
      log(`[ERROR] Failed to create default config: ${error}`)
    }
  }
}

interface Config {
  debug: boolean
  replacements: Record<string, string>
}

/**
 * Load and parse the config file
 * Returns debug flag and replacements map
 */
function loadConfig(): Config {
  const defaultConfig: Config = {
    debug: false,
    replacements: RFC2119_DEFAULTS,
  }

  try {
    if (!existsSync(CONFIG_PATH)) {
      return defaultConfig
    }

    const content = readFileSync(CONFIG_PATH, "utf-8")
    const parsed = parseJSONC(content)

    return {
      debug: parsed.debug === true,
      replacements: parsed.replacements || {},
    }
  } catch (error) {
    // Log to stderr since debug logging might not be enabled yet
    process.stderr.write(`Replacer: [WARN] Failed to parse config: ${error}\n`)
    return defaultConfig
  }
}

/**
 * Escape special regex characters in a string
 */
function escapeRegex(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, "\\$&")
}

/**
 * Apply all replacements to the given text
 * - Case-insensitive matching
 * - Word boundary aware (won't replace "must" inside "customer")
 * - Longest phrases matched first (to handle "must not" before "must")
 */
function applyReplacements(
  text: string,
  replacements: Record<string, string>,
): { result: string; counts: Map<string, number> } {
  // Sort keys by length descending (longest first)
  const sortedKeys = Object.keys(replacements).sort((a, b) => b.length - a.length)

  let result = text
  const counts = new Map<string, number>()

  for (const key of sortedKeys) {
    const value = replacements[key]
    // Create case-insensitive regex with word boundaries
    const pattern = new RegExp(`\\b${escapeRegex(key)}\\b`, "gi")

    // Count occurrences before replacement
    const matches = result.match(pattern)
    if (matches && matches.length > 0) {
      counts.set(key, matches.length)
      result = result.replace(pattern, value)
    }
  }

  return { result, counts }
}

/**
 * Replacer Plugin
 */
const Replacer: Plugin = async () => {
  // Ensure default config exists
  ensureConfigExists()

  // Initial config load to get debug state
  const initialConfig = loadConfig()
  debugEnabled = initialConfig.debug

  log("[INIT] Replacer plugin loaded")
  log(`[INIT] Config path: ${CONFIG_PATH}`)
  log(`[INIT] Loaded ${Object.keys(initialConfig.replacements).length} replacement pairs`)

  return {
    "chat.message": async (input, output) => {
      // Hot reload: re-read config on every message
      const config = loadConfig()
      debugEnabled = config.debug

      if (Object.keys(config.replacements).length === 0) {
        log("[SKIP] No replacements configured")
        return
      }

      let totalReplacements = 0
      const allCounts = new Map<string, number>()

      // Apply to text-type parts only (not file content, not other part types)
      // User-typed content comes through as TextPart with type === "text"
      for (const part of output.parts) {
        if (part.type === "text" && "text" in part && typeof part.text === "string") {
          const { result, counts } = applyReplacements(part.text, config.replacements)
          ;(part as { text: string }).text = result

          // Merge counts
          for (const [key, count] of counts) {
            allCounts.set(key, (allCounts.get(key) || 0) + count)
            totalReplacements += count
          }
        }
      }

      // Log replacements made
      if (debugEnabled && totalReplacements > 0) {
        for (const [key, count] of allCounts) {
          const value = config.replacements[key]
          const plural = count === 1 ? "occurrence" : "occurrences"
          log(`[REPLACE] '${key}' -> '${value}' (${count} ${plural})`)
        }
      }
    },
  }
}

export default Replacer
