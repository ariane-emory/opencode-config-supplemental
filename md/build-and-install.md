**BUILD STEPS**:
1. You **MUST** set `OPENCODE_VERSION` to the timestamp from this integration branch's name (e.g., `2026-02-17-17-11`) when building, otherwise the version will display as `0.0.0-integration/...-YYYYMMDDHHMM` instead of just the timestamp.
2. Build from packages/opencode/ with: `OPENCODE_VERSION="YYYY-MM-DD-HH-mm" bun run build --single`
3. The built binary will be at either `dist/opencode-darwin-arm64/bin/baseone` or `dist/opencode-linux-arm64/bin/baseone` (depending on your platform)

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

**MACOS BINARY CACHING ISSUE**:
On macOS, copying a new binary to a path where a previous binary existed can cause SIGKILL (exit status 137) due to LaunchServices or code signing cache. If `baseone --version` returns status 137:
1. Test the binary directly from the build directory to confirm it works: `packages/opencode/dist/opencode-darwin-arm64/bin/baseone --version`
2. Remove the installed binary: `rm ~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss`
3. Wait a moment, then copy it fresh: `cp packages/opencode/dist/opencode-darwin-arm64/bin/baseone ~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss`
4. Recreate the symlink: `ln -sf ~/.local/bin/baseone.YYYY-MM-DD-HH-mm-ss ~/.local/bin/baseone`
5. Verify again
