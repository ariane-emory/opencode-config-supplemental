---
agent: build
---
I need you to build the version of OpenCode/BaseOne in the latest integration branch and install it locally in the `~/.local/bin` folder.

!`cat ~/.config/opencode/md/latest-integration.md`

**PRE-INSTALL CLEANUP**: Before building, check for and remove any existing 'baseone' or 'baseone' binaries that could conflict with the new installation. Check common locations like:
- ~/.local/bin (our target folder)
- ~/.bun/bin/baseone or base-one
- /usr/local/bin/baseone
- Any other directories in PATH

**BUILD STEPS**:
1. Build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

**INSTALLATION**: Install the binary globally as 'baseone' (NOT 'base-one'):
- Copy the binary to `~/.local/bin/baseone`
- Ensure it's executable: `chmod +x ~/.local/bin/baseone`

**POST-INSTALL VERIFICATION**: 
- Run `which baseone` to verify it's using /usr/local/bin/baseone
- Run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, check your PATH and remove conflicting binaries

