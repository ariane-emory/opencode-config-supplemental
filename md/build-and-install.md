**BUILD STEPS**:
1. You **MUST** build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

**BACKUP**: You **MUST** backup the old installed binary (if it exists):
- If there is an existing binary at `~/.local/bin/baseone`, move it to a timestamped backup at a path like `~/.local/bin/baseone.bak.YYYY-MM-DD-HH-mm-ss` (substituting in the timestamp from this integration branch's name).

**INSTALLATION**: You **MUST** install the binary globally as 'baseone' (NOT 'base-one'):
- Copy the binary to `~/.local/bin/baseone.bak.YYYY-MM-DD-HH-mm-ss` (substituting in the timestamp from this integration branch's name).
- Ensure it's executable: `chmod +x ~/.local/bin/baseone.bak.YYYY-MM-DD-HH-mm-ss`
- You **MUST** create a symbolic link to the binary at `~/.local/bin/baseone`: `ln -s ~/.local/bin/baseone.bak.YYYY-MM-DD-HH-mm-ss ~/.local/bin/baseone`.

**POST-INSTALL VERIFICATION**: 
- You **MUST** run `which baseone` to verify it's using ~/.local/bin/baseone
- You **MUST** run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, You **SHOULD** check your PATH and remove conflicting binaries
