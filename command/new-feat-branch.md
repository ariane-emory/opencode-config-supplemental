---
description: Create and implement a new feature branch and then merge it back. 
agent: plan
ignored: false
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/create-new-branch.md feat/$1`
!`baseone expand ~/ocs/md/edit-branch--body.md add "feat/$ARGUMENTS"`
!`baseone expand ~/ocs/md/noted-branch-return.md feat/$1`
