### From origin:

- feat/base-one-rebrand (**MERGE ADVICE:** the program **MUST** continue to properly load configurations (and other content) from`~/.config/opencode` if if exists! `~/.config/baseone` should be preferred **IF** it exists, but that directory **MUST NEVER** be created by the program! The user **MUST** explicitly create the new `baseone` directory themself if they want it to be used! If both directories exist, then the `baseone` directory **MUST** be used and the `opencode` directory **MUST** be ignored. If both directories exist, then their contents **MUST NEVER** be merged, it's one directory **OR** the other!)
- feat/sinister-quotes (**MERGE ADVICE:** the placeholders used **MUST** be the SINISTER_PLACEHOLDERS array in this branch's packages/ui/src/constants/placeholders.ts file, **NO OTHER PLACEHOLDER SOURCE/LOCATION IS PERMISSIBLE!**)

### From foreign remotes:

- feat/markdown-renderer (from the gignit remote)
- feat/thinking-indicator-hidden (from the rcdailey remote)

### More from origin:

- feat/session-grouping
- feat/session-bookmarks
- fix/dialog-datetime-alignment (**MERGE ADVICE:** for best results, merge this one immediately after feat/session-bookmarks. This feature **MUST** not be clobbered; if there is a conflict, it **MUST** be combined with the other feature with which it is conflicting!)
- feat/keybindable-commands
- feat/automatic-list-continuation
- feat/continue-command
- feat/configurable-snapshot-lifespan
- feat/configurable-new-plan-mode
- feat/enable-exa-setting
- feat/canceled-prompts-in-history (**MERGE ADVICE:**Careful not to clobber this while merging! Merging this branch **MUST** add the new item to the command palette.)
- feat/permission-spinner
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
- fix/persist-sidebar (**MERGE ADVICE:** This branch is meant not only to make the sidebar display state persistent across bestarts if the progran but also  to remove the normal behaviour where the sidebar is hidden when the terminal is not wide enough! There **MUST NOT** be a way to return to the auto state after transitioning to the "show" or "hide" state. This change in the sidebar behaviour **MUST NOT** be clobbered while merging!)
- feat/persist-sidebar-group-folding-states (**MERGE ADVICE**: Be sure not to let this feature get clobbered by subsequent merges!)
- feat/permission-indicator-in-sidebar
- feat/command-palette-consistency (**MERGE ADVICE:** to prevent recurence of a past mistake: this branch is meant to **MOVE** several items from the Session category to the System category in the command palette. You **MUST NOT** duplicate them into both categories when resolving merge conflicts! Additionally, if fix/persist-sidebar was merged previously, be sue to properly move the new logic for the sidebar that it added: no return to "auto" after leaving, et cetera) 
- feat/persistant-sidebar-overlay-behaviour
- refactor/shared-substitute
- feat/session-id-in-status
- feat/opeoginni--display-message-tps
- feat/kv-diff-style-clean
- feat/global-compaction-threshold
- feat/configurable-message-and-session-limit (**MERGE ADVICE:** Don't forget that both the `experimental._message__limit` and `experimental.session_list_limit` settings should accept either positive integers or the string value "none"!) 
- feat/experimental-dont-cache-command-markdown
- feat/jsonc-user-themes
- feat/shell-advice (**MERGE ADVICE**: Make sure to combine this properly with the changes to the bash tool's description that are made in the feat/improve-bash-tool-git-advice branch, both sets of changes must be synthesized!) 
- feat/improve-bash-tool-git-advice (**MERGE ADVICE**: Make sure to combine this properly with the changes to the bash tool's description that are made in the feat/shell-advice branch, both sets of changes must be synthesized!)
- feat/edit-tool-description
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
- feat/alphabetical-message-modal
- feat/toggle-sidebar-scrollbar
- feat/full-datetimes-in-fork-and-timeline-dialogues
- feat/configurable-maximum-prompt-input-size
- fix/always-allow-folding-sidebar-mcps
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
- feat/distinct-title-colour (**MERGE ADVICE:** Make sure that this change in the title's colouring is made compatible with the reformatting in grouped session titles that comes from the feat/session-grouping branch, **BOTH** the distinct colour for the titles **AND** the formatting of grouped sessions' titles)
- feat/tool-output-colour
- feat/improve-experimental-plan-mode-prompt
- fix/input-enter-keybindings
- fix/escape-from-status
- fix/restore-footer (**MERGE ADVICE:** As its name suggests, this feature restores the footer it was removed in a previous version; it must not be allowed to be clobbered by other branches when merging!)
- feat/remove-canned-jokes
- fix/session-list-delete-selection
- feat/allow-variant_list-keybinding
- feat/kimi-with-claude-system-prompt

**NOTE**: Any lines beginning with a `#` character are comments and you **MUST** ignore any such branches!