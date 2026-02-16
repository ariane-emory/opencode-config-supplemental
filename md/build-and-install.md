**BUILD STEPS**:
1. Build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

**COME IN AND PUSH**: Ensure that the built binary has been commited and pushed to this integration branch on origin. 

**INSTALLATION**: Install the binary globally as 'baseone' (NOT 'base-one'):
- Copy the binary to `~/.local/bin/baseone`
- Ensure it's executable: `chmod +x ~/.local/bin/baseone`

**POST-INSTALL VERIFICATION**: 
- Run `which baseone` to verify it's using ~/.local/bin/baseone
- Run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, check your PATH and remove conflicting binaries
