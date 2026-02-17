**BUILD STEPS**:
1. You **MUST** build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

**COMMIT AND PUSH**: You **MUST** Ensure that the built binary has been commited and pushed to this integration branch on origin. 

**BACKUP**: You **MUST** backup the old installed binary (if it exists):
- If there is an existing binary at `~/.local/bin/baseone`, move it to a timestamped backup at a path like `~/.local/bin/baseone.bak.YYYY-MM-DD-HH-mm-ss` (substituting in it's file ceation timestamp apprropriately).

**INSTALLATION**: You **MUST** install the binary globally as 'baseone' (NOT 'base-one'):
- Copy the binary to `~/.local/bin/baseone`
- Ensure it's executable: `chmod +x ~/.local/bin/baseone`

**POST-INSTALL VERIFICATION**: 
- You **MUST** run `which baseone` to verify it's using ~/.local/bin/baseone
- You **MUST** run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, You **SHOULD** check your PATH and remove conflicting binaries
