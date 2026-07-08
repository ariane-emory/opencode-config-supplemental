**SYSTEM:** *Branch to work on: "$1"*
**SYSTEM:** *PR number for this branch: "$2"*

### Your mission: Merge the latest changes from the dev branch into the $1 branch (branch for PR #$2) while resolving any conflicts that occur

### Preparation:

You **SHOULD** take my word for it that the dev branch and the $1 branch already exist locally, so you don't have to waste time checking if they do before you start.

**CRITICAL**: First, you **MUST** run `bun install`. You **MUST** be sure to do this before you run any tests!

Then, you **MUST** run the tests to see what pre-existing errors exist. If the dev branch contains pre-existing errors, you **MUST** preserve them when you merge it into the $1 branch. If you just ran these same tests on the dev branch recently, you **SHOULD** skip re-running them and reuse the results that you saw a moment ago the last time that you ran them. The tests can take a while to run, so you **SHOULD** be sure to use a long timeout value, if any — using no time at value at all is very likely the safest option!

**NOTE**: If you happen to discover pre-existing errors in the $1 branch while you're worrking, then fixing those is alright so long as the central feature of the branch described in the pull request is preserved. If you're not sure, you can pause and ask me using the `question` tool. 

Then, you **MUST** check out the $1 branch, merge the local dev branch into it and resolve any conflicts. Bear in mind that earlier versions of the dev branch are likely to have been merged into the $1 branch many times previously! 

**CRITICAL**: You **MUST** make sure that you first discard any uncommited local changes before merging the dev branch in, it would a catastrophic failure if any unrelated local changes were accidentally commited/pushed into our target branch! You **MUST NOT** stash any pre-existing local changes, you **MUST** discard them!

If a git lock file gets in your way, you **SHOULD** delete it and keep working on merging.

### Notes:

You may need to run `bun install` again after performing the merge, if any new packages were added, but you **MAY** be able to skip the step if no packages were added or removed in the feature/fix branch.You **MUST** sure that no new test failures were introduced by the merge: you **SHOULD** disregard any pre-existing test failures in the dev branch, but we don't want to add any new test failures relative to dev. If tests fail due to a timeout (in either branch), try waiting a moment and rerunning them, it may just mean that that test is a bit flaky.

**NOTE**: If you get an error about a bun version mismatch, you may need to first install the correct version of bun! 

**REMEMBER**: The global OpenCode configuration is most likely **NOT** configured to be compatible with the branch on which you are working!

**REMEMBER**: You **SHOULD** use the `todowrite` tool to maintain a to-do list to keep track of your progress and to ensure that the user can easily monitor it as you proceed. 

### Merge procedure rules:

**CRITICAL:** You **MUST** use only the locally available copies of the branch and the dev branch. You **MUST NOT** ever pull changes from remote branches!

If a pre-existing failure was occurring on the dev branch, then the same error will likely recur in a pre-push when you try to push the merged branch to origin. If a pre-push hook is failing due to pre-existing errors, you **SHOULD** use the --no-verify flag on the push to make the push succeed.

**CRITICAL**: Do not fast forward merges. Do not rebase or cherry-pick. If conflicts occur, think them through thoroughly and carefully resolve them by hand to ensure that important changes from the current branch are not lost. The $1 branch is a tightly-scoped branch that implements a single featur and/or fix, and so you **MUST NOT** break the feature/fix that the $1 branch is meant to implement! You **MUST NOT** use the `--ours` or `--theirs` switches! When resolving conflicts you **MUST** be sure to pay special attention to any agent-directed defensive comments that begin with strings like `// AGENTS: ` or `// **CRITICAL** `since these are often used to guide agent behaviour during merging.

**IMPORTANT**: If you are able to merge in the changes and successfully resolve any conflicts and the tests all pass afterwards (excluding any pre-existing errors from the dev branch, which **MUST** be preserved), you **MUST** push the changes to git.

If you are not able to resolve any conflicts or the tests do not pass afterwards, you **MUST NOT** push the changes to git and you **MUST** ask me to step in and help you out instead. If any conflics did occur but you were able to resolve them, report on how the conflics were resolved.

**CRITICAL**: You **MUST NOT** correct pre-existing errors from the dev branch! You **MUST** only correct new errors introduced by the merge, pre-existing errors **MUST** be preserved exactly. CORRECTING PRE-EXISTING ERRORS FROM THE dev BRANCH WOULD BE A CATASTROPHIC FAILURE THAT **MUST** BE AVOIDED AT ALL COSTS!

**CRITICAL**: You **MUST NOT** do anything that could close any active pull requests!

**CRITICAL**: I AM FUCKING SERIOUS DUDE, YOU **MUST NOT** TRY FIX PRE-EXISTING ERRORS FROM dev! **ONLY** TRY TO FIX NEW ERRORS INTRODUCED BY THE MERGE!

**CRITICAL**: You **MUST NOT** ever kill bun processes! There may be other active bun processes doing work at the same time as you are that cannot be easily restarted if you kill them! In addition, you **MUST NOT** ever modify the user's global OpenCode configuration file located at ~/.config/opencode/opencode.json, to do so would destroy critical data!

There is no particular time constraint here, so don't rush, take the time that you need to do the job properly. 

**CRITICAL**: You **MUST NOT** forget to push the merged $1 branch to origin once you have finished. 

**IMPORTANT**: Focus on preserving the central feature or fix that this branch provides, as described in its pull request! The structure in the newer dev branch HEAD may have changed, in which case it is more important to adapt the intended feature or fix to the new structure rather than to preserve the exact structure in the original branch! We want to be sure to make sure that the central fix or feature provided by the branch is made to work with the current structure of the dev branch! 

### Remember: 

The $1 branch is meant to implement PR #$2 on the http://github.com/ariane-emory/opencode repository..

You **SHOULD** run the `gh --repo ariane-emory/opencode pr view $2` command to examine PR #$2's description to see what feature/fix/change the branch we are updating is meant to implement.

You **SHOULD** also run `gh --repo ariane-emory/opencode pr diff $2` command to see how this branch differs from the `dev` branch prior to performing the merge.

While updating this branch with the newest changes, you **MUST** ensure that you preserve the feature/fix described in PR #$2.

### CRITICAL: Setting the session's title:

You **MUST** use the `set_current_session_title` tool to give the session a title maching this format:
`merging dev|Merging dev into $1` 

**CRITICAL**: You **MUST** follow this format exactly! The prefixed part prior prior to the pipe character is **CRUCIAL** for our other tools to function proparly, If you do not properly add this prefix to the session title, then **NONE** of this work can be used, and the task is a total failure!

**CRITICAL**: Do **NOT** forget the pipe character shown in the title format!

You **MUST NOT** bookmark this session!

${3..}
