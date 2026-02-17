---
description: Verify solution to a GH issue.
agent: plan
model: zai-coding-plan/glm-4.7
---

This branch is meant to implement a resolution to this GH issue: 

gh --repo anomalyco/opencode issue view $ARGUMENTS;
gh --repo anomalyco/opencode issue view $ARGUMENTS --comments;

First, analyze the changes in this branch relative to the upstream/dev branch:

!`git diff dev..HEAD`
!`cat ~/.config/opencode/md/verify-gh-issue-resolution--epilogue.md`
