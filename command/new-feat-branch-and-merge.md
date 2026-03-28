---
description: Create and implement a new feature branch and then merge it back. 
agent: plan
ignored: true
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/create-new-branch.md feat/$1`
!`baseone expand ~/ocs/md/edit-branch--body.md change feat/$1 "${2..}"`
!`baseone expand ~/ocs/md/noted-branch-return-and-merge.md feat/$1`
