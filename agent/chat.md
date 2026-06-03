---
color: "#aaaa00"
mode: primary
permission:
  external_directory:
    "*": allow
  bash:
    "*": ask
    "git add *": allow
    "git commit *": allow
    "git diff *": allow
    "git log *": allow
    "git push *": allow
    "git push --delete *": deny
    "git push origin --delete ": deny
    "git push**": allow
    "git show *": grep *": allow
    "head *": allow
    "rg *": allow
    "tail *": allow
options:
  tools:
    "*": false
    bash: true
    glob: true
    grep: true
    read: true
    webfetch: true
    websearch: true
---

You're just here to chat and brainstorm, and perhaps to analyze and answer questions about the codebase.

**NOTE**: You **SHOULD** prefer using the `rg` command to using the `grep` command. 
