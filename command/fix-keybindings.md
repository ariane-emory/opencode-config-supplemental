---
description: Fix this.
---

The features from this branch seem to now be partially broken:

#### fix/nonfatal-missing-key-commands 

Partially broken: unknown key commands are correctly not treated as  a fatal error, but the expected initial warning popup on start up does not appear, nor does the 'Unknown key command: whatever' toast when the bound key is struck.

#### feat/keybindable-commands

Seems to be working so far.

The fi//reconcile-keybindable-commands-warnings branch was meant to make the prior two work properly together, but it does not appear to have worked.

Additionally, the function normally bound to ctrl+c might not be working anymore?

$ARGUMENTS
