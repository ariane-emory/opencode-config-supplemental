!`{ git reset --hard HEAD && echo "Reset."; } || echo "No reset!";`
!`git clean -fd && echo "Cleaned."`
!`git checkout dev && echo "Checked out dev."`
!`baseone expand ~/.config/opencode/md/integrate-branches--body.md $(date +%Y-%m-%d-%H-%M)`