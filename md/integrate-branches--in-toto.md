!`{ git reset --hard HEAD && echo "Reset."; } || echo "Failed to run git reset. ";`
!`git clean -fd && "Cleaned the local git repository of untracked changes. "`
!`git checkout dev && echo "Checked out dev branch."`
!`baseone expand ~/.config/opencode/md/integrate-branches--body.md $(date +%Y-%m-%d-%H-%M)`