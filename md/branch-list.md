### From the upstream remote:
- split-config-fixes (**MERGE ADVICE:** Note that you **MUST** use only the local copy of this branch when merging, you **MUST NOT** pull new versions from the upstream remote!)

### From origin:

- feat/base-one-rebrand
- feat/sinister-quotes (**MERGE ADVICE:** the placeholders used **MUST** be the SINISTER_PLACEHOLDERS array in this branch's packages/ui/src/constants/placeholders.ts file, **NO OTHER PLACEHOLDER SOURCE/LOCATION IS PERMISSIBLE!**)

### From foreign remotes:

- feat/markdown-renderer (from the gignit remote)
- feat/thinking-indicator-hidden (from the rcdailey remote)
- fix/session-new-prompt-handoff (from the AksharP5 remote)

### More from origin:

- feat/session-grouping
- feat/session-bookmarks
- fix/dialog-datetime-alignment (**MERGE ADVICE:** for best results, merge this one immediately after feat/session-bookmarks. This feature **MUST** not be clobbered; if there is a conflict, it **MUST** be combined with the other feature with which it is conflicting!)
- feat/keybindable-commands
- feat/automatic-list-continuation
- feat/continue-command
- feat/configurable-snapshot-lifespan
- feat/configurable-new-plan-mode
- feat/canceled-prompts-in-history (**MERGE ADVICE:**Careful not to clobber this while merging! Merging this branch **MUST** add the new item to the command palette.)
- feat/permission-spinner
- feat/permission-indicator-in-sidebar
- feat/opencode-expand
- feat/argument-range-syntax
- feat/default-arguments (**MERGE ADVICE:** When merging this branch, make sure that you don't accidentally reintroduce the swallowing behaviour that the feat/argument-range-syntax branch was meant to eliminate.)
- fix/preserve-quotes-in-arguments
- fix/history-navigation-key-commands
- fix/build-with-short-version (**MERGE ADVICE:** Automatically uses short timestamp version for integration branches without requiring OPENCODE_VERSION to be set)
- fix/autocompletion-filtered-order
- fix/modal-menus-filtered-order
- fix/config-package-json-pollution (**MERGE ADVICE:** This branch **MUST** be included in integration branches to prevent package.json pollution with non-SemVer versions)
- fix/session-list-viewport-jumping
- fix/merging-multiple-configs
- fix/markdown-codeblock-theme-property
- fix/persist-sidebar (**MERGE ADVICE:** This branch is meant not only to make the sidebar display state persistent across bestarts if the progran but also  to remove the normal behaviour where the sidebar is hidden when the terminal is not wide enough! There **MUST NOT** be a way to return to the auto state after transitioning to the "show" or "hide" state. This change in the sidebar behaviour **MUST NOT** be clobbered while merging! )
- feat/command-palette-consistency (**MERGE ADVICE:** to prevent recurence of a past mistake: this branch is meant to **MOVE** several items from the Session category to the System category in the command palette. You **MUST NOT** duplicate them into both categories when resolving merge conflicts! Additionally, if fix/persist-sidebar was merged previously, be sue to properly move the new logic for the sidebar that it added: no return to "auto" after leaving, et cetera) 
- refactor/shared-substitute
- feat/session-id-in-status
- feat/edit-tool-description
- feat/opeoginni--display-message-tps
- feat/kv-diff-style-clean
- feat/global-compaction-threshold
- feat/configurable-message-and-session-limit (**MERGE ADVICE:** Don't forget that both the `experimental._message__limit` and `experimental.session_list_limit` settings should accept either positive integers or the string value "none"!) 
- feat/experimental-dont-cache-markdown
- feat/jsonc-user-themes
- feat/persist-sidebar-group-folding-states
- feat/persistant-sidebar-overlay-behaviour
- feat/shell-advice
- feat/renaming-doesnt-close-session-list
- feat/session-child-toggle-key
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
- feat/configurable-maximum-prompt-input-size
- feat/clickable-sidebar-mcps
- feat/clickable-dialogue-mcps
- feat/clickable-status-mcps
- feat/ignored-commands
- feat/dialogue-background-overlay-setting
- fix/no-split-database
- feat/elapsed-timer
- fix/system-prompt-directories
- fix/rfc2119-question-tool
- feat/sidebar-header-accent-colours
- feat/distinct-title-colour
- feat/tool-output-colour
- feat/improve-experimental-plan-mode-prompt
- fix/input-enter-keybindings
- fix/escape-from-status
- fix/restore-footer (**MERGE ADVICE:** As its name suggests, this feature restores the footer it was removed in a previous version; it must not be allowed to be clobbered by other branches when merging!)
- feat/kimi-with-claude-system-prompt
- feat/remove-canned-jokes

**NOTE**: Any lines beginning with a `#` caracter are comments and you **MUST** ignore any such branches!