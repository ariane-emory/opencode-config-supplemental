---
description: Minimize changes for a GH issue.
agent: build
model: opencode/big-pickle
---

This branch is meant to resolve this GH issue: 

gh --repo anomalyco/opencode issue view $ARGUMENTS;
gh --repo anomalyco/opencode issue view $ARGUMENTS --comments;

!`cat ~/.config/opencode/md/deslop-for-gh-issue--epilogue.md`
