## Task: Merge the latest changes from the ${2:dev} branch into the $1 branch:

### Preparation:

First, check out the ${2:dev} branch,

You **SHOULD** take my word that both of these branches already exist locally, so you don't have to waste time checking if they do before you start.

Next, run `bun install`.

Then run the tests to see what pre-existing errors exist. If the ${2:dev} branch contains pre-existing errors, you **MUST** preserve them when you merge it into the $1 branch.

Then, check out the $1 branch, merge the local ${2:dev} branch into it and resolve any conflicts. The ${2:dev} branch is likely to have been merged into the $1 branch many times previously using earlier versions of the ${2:dev} branch! 

**CRITICAL**: Make sure that you first discard any uncommited local changes before merging the ${2:dev} branch in, it would a catastrophic failure if any unrelated local changes were accidentally commited/pushed into our target branch! You **MUST NOT** stash any pre-existing local changes, you **MUST** discard them!

If a git lock file gets in your way, just delete it and keep working on merging.

### CRITICAL: Setting the session's title:

You **MUST** use the `set_current_session_title` tool to give the session a title maching this format:
`merging ${2:dev}|Merging ${2:dev} into $1` 

**CRITICAL**: You **MUST** follow this format exactly! The part prior to the \pipe character is **CRUCIAL** for our other tools to function proparly!

**CRITICAL**: Do **NOT** forget the pipe character shown in the title format!

Do **NOT** bookmark this session!

### Notes:

You may need to run `bun install` again after performing the merge. .Make sure that no new test failures were introduced by the merge: you may disregard any pre-existing test failures in the ${2:dev} branch, but we don't want to add any new test failures relative to ${2:dev}. If tests fail due to a timeout (in either branch), try waiting a moment and rerunning them, it may just mean that that test is a bit flaky.

**NOTE**: If you get an error about a bun version mismatch, you may need to first install the correct version of bun! 

**REMEMBER**: The global Opencode configuration is most likely **NOT** configured to be compatible with the branch on which you are working!

### Merge procedure rules:

**CRITICAL:** You **MUST** use only the locally available copies of the branch and the ${2:dev} branch. You **MUST NOT** ever pull changes from remote branches!

If a pre-push hook is failing due to pre-existing errors, you **SHOULD** use the --no-verify flag on the push to make the push succeed.

**CRITICAL**: Do not fast forward merges. Do not rebase or cherry-pick. If conflicts occur, think them through thoroughly and carefully resolve them by hand to ensure that important changes from the current branch are not lost. The $1 branch is a tightly-scoped branch that implements a single featur and/or fix, and so you **MUST NOT** break the feature/fix that the $1 branch is meant to implement! You **MUST NOT** use the `--ours` or `--theirs` switches! When resolving conflicts you **MUST** be sure to pay special attention to any agent-directed defensive comments that begin with strings like `// AGENTS: ` or `// **CRITICAL** `since these are often used to guide agent behaviour during merging.

**IMPORTANT**: If you are able to merge in the changes and successfully resolve any conflicts and the tests all pass afterwards (excluding any pre-existing errors from the ${2:dev} branch, which **MUST** be preserved), you **MUST** push the changes to git.

If you are not able to resolve any conflicts or the tests do not pass afterwards, you **MUST NOT** push the changes to git and you **MUST** ask me to step in and help you out instead. If any conflics did occur but you were able to resolve them, report on how the conflics were resolved.

**CRITICAL**: You **MUST NOT** correct pre-existing errors from ${2:dev}! You **MUST** only correct new errors introduced by the merge, pre-existing errors **MUST** be preserved exactly. CORRECTING PRE-EXISTING ERRORS FROM THE ${2:dev} BRANCH WOULD BE A CATASTROPHIC FAILURE THAT **MUST** BE AVOIDED AT ALL COSTS!

**CRITICAL**: You **MUST NOT** do anything that could close any active pull requests!

**CRITICAL**: I AM FUCKING SERIOUS DUDE, YOU **MUST NOT** TRY FIX PRE-EXISTING ERRORS FROM ${2:dev}! **ONLY** TRY TO FIX NEW ERRORS INTRODUCED BY THE MERGE!

**CRITICAL**: You **MUST NOT** ever kill bun processes! There may be other active bun processes doing work at the same time as you are that cannot be easily restarted if you kill them! In addition, you **MUST NOT** ever modify the user's global OpenCode configuration file located at ~/.config/opencode/opencode.json, to do so would destroy critical data!

There is no particular time constraint here, so don't rush, take the time that you need to do the job properly. 

**CRITICAL**: You **MUST NOT** forget to push the merged branch to origin once you have finished. 

${3..}