# A MUST-have plugin

Automatically replaces text patterns in your prompts before they're sent to the LLM.

## What It Does

Performs case-insensitive string replacements on user-typed prompts. The primary use case is auto-capitalizing RFC2119 keywords (MUST, SHOULD, MAY, etc.) in technical specifications.

**Example**: Typing `"the system must validate input"` becomes `"the system MUST validate input"`.

### Features

- **Case-insensitive matching**: `must`, `Must`, and `MUST` all match
- **Word boundary aware**: Won't replace `must` inside `customer`
- **Multi-word phrases**: `must not` is matched as a unit (before `must` alone)
- **Hot reload**: Config changes take effect immediately (no restart needed)
- **JSONC support**: Comments allowed in config file

### Scope

- Only replaces text in user-typed prompts
- Does NOT modify file content attached via `@` mentions
- Does NOT modify slash command output

## Configuration

**Config file**: `~/.config/opencode/must-have-plugin.json`

If the config file doesn't exist, it's automatically created with RFC2119 defaults.

### Default Configuration

```jsonc
{
  // Uncomment to enable debug logging (view with: tail -f /tmp/opencode-must-have-plugin-debug.log)
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
```

### Custom Replacements

Add your own replacement pairs to the `replacements` object:

```jsonc
{
  "replacements": {
    // RFC2119 keywords
    "must": "MUST",
    "should": "SHOULD",
    
    // Custom replacements
    "todo": "TODO",
    "fixme": "FIXME",
    "api": "API",
    "url": "URL"
  }
}
```

### Configuration Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `debug` | boolean | `false` | Enable debug logging to `/tmp/opencode-must-have-plugin-debug.log` |
| `replacements` | object | RFC2119 keywords | Key-value pairs for text replacement |

## Debug Logging

Enable debug mode to see what replacements are being made:

1. Edit `~/.config/opencode/must-have-plugin.json`
2. Uncomment or add `"debug": true`
3. View logs in real-time:

```bash
tail -f /tmp/opencode-must-have-plugin-debug.log
```

### Log Format

```
2026-01-19T15:30:42.123Z [INIT] Replacer plugin loaded
2026-01-19T15:30:42.124Z [INIT] Loaded 11 replacement pairs
2026-01-19T15:31:05.456Z [REPLACE] 'must' -> 'MUST' (2 occurrences)
2026-01-19T15:31:05.457Z [REPLACE] 'should not' -> 'SHOULD NOT' (1 occurrence)
```

## RFC2119 Keywords

The default configuration includes all keywords from [RFC 2119](https://datatracker.ietf.org/doc/html/rfc2119), which defines requirement levels for use in technical specifications:

| Keyword | Meaning |
|---------|---------|
| MUST / REQUIRED / SHALL | Absolute requirement |
| MUST NOT / SHALL NOT | Absolute prohibition |
| SHOULD / RECOMMENDED | Recommended, but valid reasons may exist to ignore |
| SHOULD NOT / NOT RECOMMENDED | Not recommended, but may be acceptable in some cases |
| MAY / OPTIONAL | Truly optional |

## Troubleshooting

### Replacements not working

1. Check that the config file exists: `cat ~/.config/opencode/must-have-plugin.json`
2. Verify JSON syntax is valid (remember: JSONC allows `//` comments)
3. Enable debug mode and check the log file

### Unexpected replacements

- Replacements use word boundaries, so `must` won't match inside `customer`
- Multi-word phrases are matched first, so `must not` won't become `MUST not`
- Check for typos in your replacement keys

### Config changes not taking effect

The plugin re-reads the config on every message, so changes should be immediate. If not:

1. Verify you saved the config file
2. Check for JSON syntax errors
3. Restart OpenCode as a last resort
