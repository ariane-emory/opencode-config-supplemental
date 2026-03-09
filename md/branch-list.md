### From the upstream remote:
- split-config-fixes (Note that you **MUST** use only the local copy of this branch, you **MUST NOT** pull new versions from the upstream remote!)

### From origin:

- feat/base-one-rebrand
- feat/sinister-quotes (**MERGE ADVICE**: the placeholders used **MUST** be the SINISTER_PLACEHOLDERS array in this branch's packages/ui/src/constants/placeholders.ts file, **NO OTHER PLACEHOLDER SOURCE/LOCATION IS PERMISSIBLE!**)
-
### From foreign remotes:

- feat/markdown-renderer (from the gignit remote)
- feat/thinking-indicator-hidden (from the rcdailey remote)
- fix/session-new-prompt-handoff (from the AksharP5 remote)

### More from origin:

- feat/session-grouping
- feat/session-bookmarks
- fix/dialog-datetime-alignment (**MERGE ADVICE:** for best results, merge this one immediately after feat/session-bookmarks. This feature **MUST** not be clobbered; if there is a conflict, it **MUST** be combined with the other feature with which it is conflicting!)
- feat/keybindable-commands
- feat/opencode-expand
- feat/automatic-list-continuation
- feat/continue-command
- feat/configurable-snapshot-lifespan
- feat/configurable-new-plan-mode
- feat/config-imports
- feat/canceled-prompts-in-history
- feat/permission-spinner
- feat/permission-indicator-in-sidebar
- fix/no-split-database
- fix/system-prompt-directories
- fix/remove-dot-true
- fix/rfc2119-question-tool
- fix/autocompletion-filtered-order
- fix/modal-menus-filtered-order
- fix/config-package-json-pollution (**MERGE ADVICE:** This branch **MUST** be included in integration branches to prevent package.json pollution with non-SemVer versions)
- fix/session-list-viewport-jumping
- fix/merging-multiple-configs
- fix/markdown-codeblock-theme-property
- fix/persist-sidebar
- feat/command-palette-consistecy (**MERGE ADVICE:** to prevent recurence of a past mistake: this branch is meant to **MOVE** several items from the Session category to the System category in the command palette. You **MUST NOT** duplicate them into both categories when resolving merge conflicts! Additionally, if fix/persist-sidebar was merged previously, be sue to properly move the new logic for the sidebar that it added: no return to "auto" after leaving, et cetera) 
- refactor/shared-substitute
- feat/argument-range-syntax
- feat/default-arguments (**MERGE ADVICE:** When merging this branch, make sure that you don't accidentally reintroduce the swallowing behaviour that the feat/argument-range-syntax branch was meant to eliminate.)
- feat/session-id-in-status
- feat/edit-tool-description
- feat/opeoginni--display-message-tps
- feat/kv-diff-style-clean
- feat/global-compaction-threshold
- feat/configurable-message-and-session-limit
- feat/experimental-dont-cache-markdown
- feat/jsonc-user-themes
- feat/persist-sidebar-group-folding-states
- feat/persistant-sidebar-overlay-behaviour
- feat/shell-advice
- feat/elapsed-timer
- feat/set-session-title
- feat/get-session-title
- feat/no-disabled-lsps-in-sidebar
- feat/agent-timestamps
- feat/rewind-modal-option
- feat/alphabetize-command-palette-groups
- feat/taller-dialogs
- feat/add-arianes-themes
- feat/aspiers--readline-additions
- feat/sidebar-clock
- feat/improve-bash-tool-git-advice
- feat/alphabetical-message-modal
- feat/toggle-sidebar-scrollbar
- feat/full-datetimes-in-fork-and-timeline-dialogues
- feat/clickable-sidebar-mcps
- feat/clickable-dialogue-mcps
- feat/ignored-commands
- feat/sidebar-header-accent-colours
- feat/distinct-title-colour
- feat/tool-output-colour
- feat/improve-experimental-plan-mode-prompt
- feat/renaming-doesnt-close-session-list
- fix/restore-footer (MERGE ADVICE:** As its name suggests, this feature restores the footer it was removed in a previous version; it must not be allowed to be clobbered by other branches when merging!)

**NOTE**: Any lines beginning with a `#` caracter are comments and you ****MUST**** ignore any such branches!