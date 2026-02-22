### From origin:

- fix/remove-dot-true
- fix/rfc2119-question-tool
- fix/restore-footer
- fix/persist-sidebar
- fix/autocompletion-filtered-order
- fix/modal-menus-filtered-order
- fix/bad-plugin-errors
- fix/config-package-json-pollution (MUST be included in integration branches to prevent package.json pollution with non-SemVer versions)
- fix/session-list-viewport-jumping
- fix/merging-multiple-configs
- fix/dialog-datetime-alignment (for best results, merge this one after feat/session-bookmarks)
- refactor/shared-substitute
- feat/argument-range-syntax
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
- feat/keybindable-commands
- feat/configurable-snapshot-lifespan
- feat/configurable-new-plan-mode
- feat/config-imports
- feat/canceled-prompts-in-history
- feat/command-palette-consistecy
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
- feat/sinister-quotes (the placeholders used MUST be the SINISTER_PLACEHOLDERS array in this branch's packages/ui/src/constants/placeholders.ts file, NO OTHER PLACEHOLDER SOURCE/LOCATION IS PERMISSIBLE!)

###  Afterwards, from foreign remotes:

- feature/markdown-renderer (from the gignit remote)
- feature/compaction-model (from the gignit remote)
- add-bash-env-parameter (from the taxilian remote)
- feat/thinking-indicator-hidden (from the rcdailey remote)
- session-new-prompt-handoff (from the AksharP5 remote)

### And finally, finish up with this branch:

- feat/base-one-rebrand

(skip any branches on lines which are commented out with a "#" charcter)
