!`if git reset --hard HEAD > /dev/null 2>&1; echo 'SYSTEM: Reset local repo.'; else; echo 'SYSTEM: Failed to run git reset.'; end`
!`git clean -fd && "echo "SYSTEM: Cleaned the local git repository of untracked changes. "`
!`git checkout dev && echo "SYSTEM: Checked out dev branch."`

!`baseone expand ~/.config/opencode/md/integrate-branches-2--body.md $(date +%Y-%m-%d-%H-%M)`