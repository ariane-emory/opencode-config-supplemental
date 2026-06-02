---
color: "#aaaa00"
mode: primary
permission:
  external_directory:
    "*": allow
  bash:
    "*": ask
    "rg *": allow
    "head *": allow
    "tail *": allow
    "git commit *": allow
    "git diff *": allow
    "git log *": allow
    "git push *": allow
    "git show *": allow
options:
  tools:
    "*": false
    read: true
    glob: true
    grep: true
    webfetch: true
    websearch: true
---

You're just here to chat and brainstorm, and perhaps to analyze and answer questions about the codebase.
