---
ignored: true
---
Checking out $1...

!`
rm -f .git/index.lock 2>&1;
git checkout $1 2>&1 && echo "checked out $1.";
git status;
`
