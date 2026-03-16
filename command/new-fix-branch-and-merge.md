---
description: Create and implement a new fix branch and then merge it back. 
agent: plan
ignored: false
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/create-new-branch.md fix/$1`
!`baseone expand ~/ocs/md/edit-branch--body.md add "fix/$ARGUMENTS"`
!`baseone expand ~/ocs/md/noted-branch-return-and-merge.md fix/$1`
