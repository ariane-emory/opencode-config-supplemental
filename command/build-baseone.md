---
agent: build
---
!`cat ~/.config/opencode/md/latest-integration.md`

**PRE-INSTALL CLEANUP**: Before building, check for and remove any existing 'baseone' or 'base-one' binaries that could conflict with the new installation. Check common locations like:
- ~/.bun/bin/baseone or base-one
- /usr/local/bin/baseone
- Any other directories in PATH

**BUILD STEPS**:
1. Build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at `dist/opencode-linux-arm64/bin/base-one` (adjust for your platform)

**INSTALLATION**: Install the binary globally as 'baseone' (NOT 'base-one'):
- Copy the binary to /usr/local/bin/baseone
- Ensure it's executable: `chmod +x /usr/local/bin/baseone`

**POST-INSTALL VERIFICATION**: 
- Run `which baseone` to verify it's using /usr/local/bin/baseone
- Run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, check your PATH and remove conflicting binaries

