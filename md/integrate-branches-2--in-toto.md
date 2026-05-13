!`git reset --hard HEAD >/dev/null 2>/dev/null; and echo 'SYSTEM: Reset local repo.'; or echo 'SYSTEM: Failed to run git reset.'`
!`git clean -fd && "echo "SYSTEM: Cleaned the local git repository of untracked changes. "`
!`git checkout dev && echo "SYSTEM: Checked out dev branch."`

!`baseone expand ~/.config/opencode/md/integrate-branches-2--body.md $(date +%Y-%m-%d-%H-%M)`