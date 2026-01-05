---
description: Fix this.
---

The features from this branch seem to now be partially broken:

#### fix/nonfatal-missing-key-commands 

Partially broken: unknown key commands are correctly not treated as a fatal error, but the expected initial warning popup on start up does not appear, nor does the 'Unknown key command: whatever' toast when the bound key is struck.

#### feat/keybindable-commands

Not working at all.

Additionally, the function normally bound to ctrl+c might not be working anymore?

$ARGUMENTS
