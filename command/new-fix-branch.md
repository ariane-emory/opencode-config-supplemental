---
description: Create and implement a new fix branch.
agent: plan
ignored: true
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/create-new-branch.md fix/$1`
!`baseone expand ~/ocs/md/edit-branch--body.md add fix/$1 $2`
!`baseone expand ~/ocs/md/noted-branch-return.md fix/$1`