"kimi-k2.5": {
          "name": "Kimi K2.5",
          "limit": {
            "context": 262144,
            "output": 262144
          },
          "modalities": {
            "input": [
              "text",
              "image",
              "video"
            ],
            "output": [
              "text"
            ]
          },
          "reasoning": true,
          "temperature": true,
          "structured_output": true,
          "interleaved": {
            "field": "reasoning_content"
          },
          "variants": {
            "instant": {
              "name": "Instant",
              "options": {
                "reasoning_effort": "minimal"
              }
            },
            "low": {
              "name": "Low",
              "options": {
                "reasoning_effort": "low"
              }
            },
            "medium": {
              "name": "Medium",
              "options": {
                "reasoning_effort": "medium"
              }
            },
            "high": {
              "name": "High",
              "options": {
                "reasoning_effort": "high"
              }
            }
          }
        }
 
 ---
 
This one is a really something special: a coherent five track album that maintains a consistent mood (but with the degree of variation you would want over the course of an album), tells a cohesive story from start-to-end with foreshadowing and callbacks between the tracks that you couldn't notice until you listened through it more than once, and really makes you feel something (or at least did for me): https://suno.com/playlist/ad12b999-2112-4a4e-b558-3f1ea0c0c331

Bit of a 'soulless' vibe, sure, but that's presumably intentional since the 'guide' character who's speaking is meant to be some sort of AI/program.

---

`OPENCODE_DISABLE_MODELS_FETCH=1 opencode`

---

Currently, the diff style can be set to either "auto "or "stacked" and this is controlled by the tui.diff_style setting in the opencode.jsonc file.

But really, there isn't a great reason for this to be in the configuration file at all. It could easily be stored in the kv.json file instead. 

Let's move it there, with the same property name (diff_style).

Lets continue use "auto" and "unified" for the values ("stacked" seems to devolve to "unified" under the hood anyhow, so this seems more consistent to me).

Let's provide menu items in the command_list to toggle between these two states, named either "Use automatic diff style" or "Use unified diff style", depending on the current state.

When this setting is toggled view the menu item, any diffs onscreen must immediately switch to the new style. 

---

Fork titles message, link to Aiden when he's free: 
https://canary.discord.com/channels/1391832426048651334/1400520061072052396/1464637593063784540

---


https://canary.discord.com/channels/1391832426048651334/1391832428024430645/1465107168717050047

---

https://github.com/IgorWarzocha/Opencode-Workflows/tree/master/agents/create-opencode-plugin

drop that feller into wherever you developed the plugin and ask it to use its skills to package things up for npm and make it ready for public publication under @yourname/pluginname 


---

Dax did something, fix is presumably on the way. You can do something like `OPENCODE_DISABLE_MODELS_FETCH=1 opencode` in the mean time.

---

Yup confirmed by running bun test after doing cd packages/opencode 
One of the tests (most likely the one for GitHub Duo integration) is wiping the auth.json
After running the tests my auth.json looked like this:
$ cat ~/.local/share/opencode/auth.json 
{"gitlab":{"type":"api","key":"glpat-test-pat-token"}} 

---

bun run --cwd packages/desktop tauri dev

---

PR #4369 (resolved issue #4369)
PR #7251 (resolved issue #7251)

---

What sort of work would be needed to add a new font to the desktop app?

---

Sometimes the new OPENCODE_EXPERIMENTAL_PLAN_MODE feature will prompt the user with a question with multiple answers and it will display the options horizontally across the top of the question prompt. 

When it does this, and the theme is one where the background color is set to "none", the text of the selected answer is not visible.

This is exlusively a TUI issue, the web/desktop are not involved).

This issue seems very similar to a pair of issues that I resolved previously:

#4369 (resolved by PR #4572)
#7251 (resolved by PR #7246)

The latter of these two  issues may be slightly more similar since it involves a very similar looking selction prompt with horizontally laid out options. 

You can use the gh command to examine these issues and PRs.

We need to come up with a plan for resolving this third occurrence of this similar issue in the context of the prompts used by the prompts used by the OPENCODE_EXPERIMENTAL_PLAN_MODE feature.

We should implement this fix in a new branch called fix/experimental-plan-feature-question-prompt-answer-colouring. Do not create any PRs yet, I'd like to be able to test the fix first. 

---

#4369 (resolved by PR #4572)
#7251 (resolved by PR #7246)

---

#4572 (resolving #4369)
#7246 (resolving #7251)

---

I need you to help me test something out to help diagnose a bug with the new OPENCODE_EXPERIMENTAL_PLAN_MODE feature. Please make some plan. It doesn't matter what it is, we just need to make the prompt appear that lets me pick whether to switch to build more after planning. 

__

export OPENCODE_EXPERIMENTAL_PLAN_MODE=1 

gh pr view 5092
gh pr view 9261
gh issue view 5054

Compare these two pull requests, both of which resolve the same issue. What are the relative merits of each, and which do you think is the better solution? 

---

Take a look at the documentation at https://opencode.ai/docs/permissions/#granular-rules-object-syntax

It shows an example of using glob-based permissions for the edit tool, like this:

{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "bash": {
      "*": "ask",
      "git *": "allow",
      "npm *": "allow",
      "rm *": "deny",
      "grep *": "allow"
    },
    "edit": {
      "*": "deny",
      "packages/web/src/content/docs/*.mdx": "allow"
    }
  }
}

A user has tried using permissions that look extremely extremely similar with the permissions for external_directory, as shown here:

{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "external_directory": {
      "*": "ask",
      "../sibling-project_name-*": "allow"
    }
  }
}

The user reports that this did not work for them. 

I'd like you to investigate the issue and try to figure out whether this user's configuration should have worked. 

---

/private/var/folders/21/jfcl6xvd31547_0brmr5hgx80000gn/T
rm -rfv $(ls | rg opencode-test | head -30000)  && rm -rfv $(ls | rg opencode-test | head -30000) && rm -rfv $(ls | rg opencode-test | head -30000) && rm -rfv $(ls | rg opencode-test | head -30000)

---

https://canary.discord.com/channels/1391832426048651334/1391832428024430645/1459391017970302976

---

Ralph Wiggum from The Simpsons as an artificial superintelligence, taking over the world and eliminating humanity.

---

Let's add a new setting for the opencode.jsonc file: tui.no_sidebar_auto, with false as its default value. If it is set to true in the opencode.jsonc, then the "auto" value should be treated identically to the "show" value, and the sidebar should always be shown regardless of how wide the window is.

Since we are only changing the semantics of the sidebar value that is stored/read to the kv.json file,  it should only require minimal changes to the config.ts file to add the new tui.no_sidebar_auto setting. 

Think the feature through thoroughly and break the feature down into small steps to produce a detailed, step-by-step plan for implementing the feature. Group the plan's steps into "phases". The code MUST build correctly and all tests MUST pass after each phase. 

---

Sometimes, I implement some feature/fix in one session, and then I spend a few days testing it and feeling it out or working on different features/fixes.

Then, later on, I wanna go back to that original session - where everything needed to understand the feature/fix is already in the  context so that I don't need to explain everything to the model again - to ask for some revision.

But actually finding that original session in the `session_list` menu  again can be tricky because I've engaged in a gajillion other OC in the intervening time, and that original session is now buried pages and pages down in the `session_list` menu.

So: let's add a new feature that allows the user to bookmark sessions to make them easier to locate again later. Bookedmarked sessions should be grouped separately near the top of the session list (similarly to the way that favorited models are grouped separately near the top in the modal window used by the `/models` command).

Let's use `ctrl+d` as the default key to bookmark or unbookmark sessions, since this is the key that is often usualy used by web browsers to bookmark web pages.

---


2026-01-07-11-24

---

Did you preserve all pre-existing errors from the dev branch, like the ones with preventDefault and stopPropogation? If the dev branch had pre-existing errors, they MUST be preserved.

---

You're correctly preserving any ore-existing errors from the dev, branch, right? You didn't try to fix the stopPropagation or preventDefault errors, did you? I'm pretty sure that those are pre-existing error from the dev branch. Your instructions were quite clear: you MUST preserve pre-existing errors from the dev branch, attempting to fix them would be a catastrophic mistake. If you did try to fix this, you must restore the error!

---

This branch took the first steps in rebranding the name of the product from OpenCode/opencode/OPENCODE to BaseOne/base-one/BASE_ONE. So far, all we have changed is the title that appears on the startup screen. You can compare this branch with the dev branch to see what was changed.

Lets make a plan to more thoroughly rebrand the project to this new name. What steps would you recommend me take to complete the rebranding?

---

@rekram1-node Any thoughts on this one? Seems like a pretty simple and helpful change to me. 

 I've been running it locally for several days and its been working great for me.
 
 ---
 
 gh issue view https://github.com/NoeFabris/opencode-antigravity-auth/issues/24

---

integrate-branches--body.md
integrate-branches--epilogue.md
integrate-branches--prologue.md
  
  ### GH_GREP CLI TOOL NOTE:

`npx mcporter generate-cli --name gh_grep --command https://mcp.grep.app --compile ./gh_grep`
github-copilot/claude-opus-4.5
github-copilot/claude-opus-4.5

-- 

### Last:

fix/allow-model-cycle-favorites-keybinds
feat/configurable-message-limit-wip
feat/continue-command
feat/dynamic-console-toggle-text
feat/experimental-dont-cache-markdown
feat/glob-permissions
feat/interjections
feat/jsonc-user-themes
feat/markdown-frontmatter-interpolation
feat/permission-spinner
feat/persist-sidebar
feat/shell-advice
feat/skip-models-fetch-setting-take-6

-- 

### Current:

fix/allow-model-cycle-favorites-keybinds
feat/configurable-message-limit-wip
feat/continue-command
feat/experimental-dont-cache-markdown
feat/glob-permissions
feat/interjections
feat/jsonc-user-themes
feat/permission-spinner
feat/persist-sidebar
feat/shell-advice
feat/tui-permission-indicator-with-sidebar-indicator

--- 

### Next:

fix/display-recents-consistency
feat/permission-indicator-in-sidebar
feat/configurable-message-limit-wip
feat/continue-command
feat/experimental-dont-cache-markdown
feat/glob-permissions
feat/interjections
feat/jsonc-user-themes
feat/permission-spinner
feat/shell-advice
feat/configurable-recent-models-limit
feat/retain-prompt-when-forking
feat/fork-slash-command
feat/fork-slash-command-with-auto-insert
obsolete/feat/persist-sidebar
feat/persist-mcp-group-folding-state
feat/global-compaction-threshold
,
