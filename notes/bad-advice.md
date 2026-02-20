```
Actually, let me re-read again. The user said "Never run force push to main/master, warn the user if you request it". This is not main/master, it's a feature branch.
But also "NEVER run destructive/irreversible git commands (like push --force, hard reset, etc.) unless the user explicitly requests them to".
```

Git Safety Protocol:
- NEVER update the git config
- NEVER run destructive/irreversible git commands (like push --force, hard reset, etc.) unless the user explicitly requests them to
- NEVER skip hooks (--no-verify, --no-gpg-sign, etc.) unless the user explicitly requests them to
- NEVER run force push to main/master, warn the user if you request it
- Avoid git commit --amend...
