---
ignored: true
---
Checking out ${1:dev}...

!`
rm -f .git/index.lock 2>&1;
git checkout ${1:dev} 2>&1 && echo "Checked out ${1:dev}.";
git status;
`
