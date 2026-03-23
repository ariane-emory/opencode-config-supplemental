---
description: Edit an existing branch and then merge it back. 
agent: plan
ignored: true
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/edit-branch--body.md change $1 "${2..}"`
!`baseone expand ~/ocs/md/noted-branch-return-and-merge.md $1`