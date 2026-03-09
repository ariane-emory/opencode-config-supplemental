---
description: Edit an existing branch and then return. 
agent: plan
ignored: false
---
!`baseone expand ~/ocs/md/note-branch.md`
!`baseone expand ~/ocs/md/edit-branch--body.md change "$ARGUMENTS"`
!`baseone expand ~/ocs/md/noted-branch-return.md`