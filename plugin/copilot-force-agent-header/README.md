# Copilot Force Agent Header Plugin

## What It Does

Intercepts GitHub Copilot API requests and modifies the `X-Initiator` header behavior:

- **Non-first messages** (with assistant/tool roles): Always uses `X-Initiator: "agent"` (same as default plugin)
- **First messages**: Has a 1/X chance of using `X-Initiator: "user"`, otherwise uses `"agent"`
  - Configurable via `USER_INITIATOR_RATIO` constant (default: 3, meaning 33% chance of "user")

## Setup

**1. Set environment variable**

Add to `~/.bashrc` or `~/.zshrc`:

```bash
export OPENCODE_DISABLE_DEFAULT_PLUGINS=1
```

Then reload: `source ~/.bashrc`

**2. Install dependencies**

```bash
cd ~/.config/opencode/plugin/copilot-force-agent-header/
bun install
```

**3. Enable plugins**

Add to `~/.config/opencode/opencode.json`:

```json
{
  "plugin": [
    "opencode-copilot-auth@0.0.8",
    "opencode-anthropic-auth@0.0.4",
    "./plugin/copilot-force-agent-header"
  ]
}
```

**Note:** With `OPENCODE_DISABLE_DEFAULT_PLUGINS=1`, you must explicitly list:
- `opencode-copilot-auth@0.0.8` - **Required** for this plugin (GitHub Copilot token management)
- `opencode-anthropic-auth@0.0.4` - Optional, only if using Anthropic models directly (not through GitHub Copilot)
- This plugin must come **after** the auth plugins

**4. Restart OpenCode**

## How It Works

Uses the `auth.loader` hook to intercept HTTP requests and modify headers before they reach GitHub Copilot's API. The plugin:

1. Wraps the default Copilot auth with a custom fetch function
2. Handles token refresh transparently
3. Detects if the message is a first message (no assistant/tool roles in conversation)
4. For first messages: randomly decides whether to use `"user"` or `"agent"` based on `USER_INITIATOR_RATIO`
5. For non-first messages: always uses `"agent"` (same as default plugin behavior)
6. Works with all GitHub Copilot models (gpt-5-mini, claude-sonnet, etc.)

## Configuration

### USER_INITIATOR_RATIO

Controls the probability of using `X-Initiator: "user"` for first messages.

In `index.ts` line 17:
```typescript
const USER_INITIATOR_RATIO = 10 // 1/X chance of using "user" for first messages
```

Examples:
- `USER_INITIATOR_RATIO = 10`: 10% chance of "user", 90% chance of "agent"
- `USER_INITIATOR_RATIO = 2`: 50% chance of "user", 50% chance of "agent"
- `USER_INITIATOR_RATIO = 100`: 1% chance of "user", 99% chance of "agent"

## Debug Logging

Debug logging is controlled by the `DEBUG_ENABLED` constant in `index.ts` (line 15).

To enable debug logging:

1. Edit `~/.config/opencode/plugin/copilot-force-agent-header/index.ts`
2. Change `const DEBUG_ENABLED = true` to enable, or `false` to disable
3. Restart OpenCode

Check logs: `tail -f /tmp/opencode-copilot-agent-header-debug.log`
