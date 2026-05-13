!`git reset --hard HEAD > /dev/null 2>&1 && echo "SYSTEM: Reset repo state."; 
 git clean -fd > /dev/null 2>&1 && echo "SYSTEM: Cleaned untracked files.";
 git checkout dev > /dev/null 2>&1 && echo "SYSTEM: Checked out the dev branch.";`

