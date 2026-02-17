**BUILD STEPS**:
1. You **MUST** build the opencode package with `bun run build --single` from packages/opencode/
2. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

**BINARY PATH**: We will be installing the binary at this path:
`~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss` (substituting in the timestamp from this integration branch's name)

**BACKUP IF NEEDED**: If a file already exists at the binary path, use `mv` to add the suffix `.bak` end of its name. If an older backup file with this name exists, it may be deleted. 

**INSTALLATION**: You **MUST** perfrm these steps to install the binary globally:
- Copy the binary to `~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss` (substituting in the timestamp from this integration branch's name).
- Ensure that it's executable: `chmod +x ~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss`
- You **MUST** create a symbolic link to the binary at `~/.local/bin/baseone`: `ln -s ~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss ~/.local/bin/baseone`.

**POST-INSTALL VERIFICATION**: 
- You **MUST** run `which baseone` to verify it's using ~/.local/bin/baseone
- You **MUST** run `baseone --version` to confirm the correct version is installed
- If an old version is still being used, You **SHOULD** check your PATH and remove conflicting binaries
