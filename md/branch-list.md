### From origin:

- split-config-fixes
- fix/system-prompt-directories
- fix/remove-dot-true
- fix/rfc2119-question-tool
- fix/restore-footer
- fix/persist-sidebar
- fix/autocompletion-filtered-order
- fix/modal-menus-filtered-order
#fix/bad-plugin-errors
- fix/config-package-json-pollution (MUST be included in integration branches to prevent package.json pollution with non-SemVer versions)
- fix/session-list-viewport-jumping
- fix/merging-multiple-configs
- refactor/shared-substitute
- feat/command-palette-consistecy (NOTE to prevent recurence of a past mistake: this branch is meant to MOVE several items from the Session category to the System category in the command palette. You MUST NOT duplicate them into both categories when resolving merge conflicts!) 
- feat/session-id-in-status
- feat/argument-range-syntax
- feat/default-arguments
- feat/opencode-expand
- feat/edit-tool-description
- feat/opeoginni--display-message-tps
- feat/kv-diff-style-clean (not yet merged upstream, but hopefully soon)
- feat/global-compaction-threshold
- feat/configurable-message-and-session-limit
- feat/experimental-dont-cache-markdown
- feat/jsonc-user-themes
- feat/permission-indicator-in-sidebar
- feat/permission-spinner
- feat/persist-sidebar-group-folding-states
- feat/persistant-sidebar-overlay-behaviour
- feat/shell-advice
- feat/elapsed-timer
- feat/sidebar-no-auto-setting
- feat/set-session-title
- feat/get-session-title
- feat/automatic-list-continuation
- feat/continue-command
- feat/session-bookmarks
- fix/dialog-datetime-alignment (for best results, merge this one immediately after feat/session-bookmarks)
- feat/keybindable-commands
- feat/configurable-snapshot-lifespan
- feat/configurable-new-plan-mode
- feat/config-imports
- feat/canceled-prompts-in-history
- feat/no-disabled-lsps-in-sidebar
- feat/agent-timestamps
- feat/rewind-modal-option
- feat/session-grouping
- feat/alphabetize-command-palette-groups
- feat/taller-dialogs
- feat/full-datetimes 
- feat/add-arianes-themes
- feat/aspiers--readline-additions
- feat/sidebar-clock
- feat/ignored-commands
- feat/improve-bash-tool-git-advice
- feat/alphabetical-message-modal
- feat/sinister-quotes (the placeholders used MUST be the SINISTER_PLACEHOLDERS array in this branch's packages/ui/src/constants/placeholders.ts file, NO OTHER PLACEHOLDER SOURCE/LOCATION IS PERMISSIBLE!)

###  Afterwards, from foreign remotes:

- feature/markdown-renderer (from the gignit remote)
- feature/compaction-model (from the gignit remote)
- add-bash-env-parameter (from the taxilian remote)
- feat/thinking-indicator-hidden (from the rcdailey remote)
- fix/session-new-prompt-handoff (from the AksharP5 remote)

### And finally, finish up with this branch:

- feat/base-one-rebrand

**NOTE**: Any lines beginning with a `#` caractter are comments and you **MUST** ignore any such branches!