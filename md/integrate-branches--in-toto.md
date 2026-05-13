!`git reset --hard HEAD > /dev/null && echo "- SYSTEM: Reset repo state."; 
 git clean -fd > /dev/null && echo "- SYSTEM: Cleaned untracked files.";
 git checkout dev > /dev/null && echo "- SYSTEM: Checked out the dev branch.";`

!`baseone expand ~/.config/opencode/md/integrate-branches--body.md $(date +%Y-%m-%d-%H-%M) $ARGUMENTS`