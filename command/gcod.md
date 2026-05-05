---
agent: build
ignored: true
---
!`git reset --hard HEAD && echo "Reset."; 
  git clean -fd && echo "Cleaned."
  git checkout dev && echo "Checked out dev."`