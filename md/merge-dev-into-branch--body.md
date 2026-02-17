### Task: Merge the latest changes from the dev branch into the $1 branch:

First, check out the dev branch and run the tests to see what pre-existing errors exist. If the dev branch contains pre-existing errors, you MUST preserve them when you merge it into the $1 branch.

Then, check out the $1 branch, merge the local dev branch into it and resolve any conflicts. You can take my word that both of these branches already exist, so you don't have to waste time checking if they do before you start.

Use the `set_current_session_title` tool to give the session a title saying that you are merging dev into the branch in question, such as "Merging dev into $1". Do **NOT** bookmark this session!

**CRITICAL**: Make sure that you first discard any uncommited local changes before merging the dev branch in, it would a catastrophic failure if any unrelated local changes were accidentally commited/pushed into our target branch! You MUST NOT stash any pre-existing local changes, you MUST discard them!

If a git lock file gets in your way, just delete it and keep working on merging.

Make sure that no new test failures were introduced by the merge: you may disregard any pre-existing test failures in the dev branch, but we don't want to add any new test failures relative to dev. If tests fail due to a timeout (in either branch), try waiting a moment and rerunning them, it mayjust mean that that test is a bit flaky.

**REMEMBER**: The global Opencode configuration is most likely NOT configured to be compatible with the branch on which you are working!

If a pre-push hook is failing due to pre-existing errors, you SHOULD use the --no-verify flag on the push to make the push succeed.

**CRITICAL**: Do not fast forward merges. Do not rebase or cherry-pick. If conflicts occur, think them through thoroughly and carefully resolve them by hand to ensure that important changes from the current branch are not lost.

**IMPORTANT**: If you are able to merge in the changes and successfully resolve any conflicts and the tests all pass afterwards (excluding any pre-existing errors from the dev branch, which MUST be preserved), you MUST push the changes to git.

If you are not able to resolve any conflicts or the tests do not pass afterwards, you MUST NOT push the changes to git and you MUST ask me to step in and help you out instead. If any conflics did occur but you were able to resolve them, report on how the conflics were resolved.

**CRITICAL**: You MUST NOT correct pre-existing errors from dev! You MUST only correct new errors introduced by the merge, pre-existing errors MUST be preserved exactly. CORRECTING PRE-EXISTING ERRORS FROM THE DEV BRANCH WOULD BE A CATASTROPHIC FAILURE!

**CRITICAL**: You MUST NOT do anything that could close any active pull requests!

**CRITICAL**: I AM FUCKING SERIOUS DUDE, YOU MUST NOT TRY FIX PRE-EXISTING ERRORS FROM dev! ONLY TRY TO FIX NEW ERRORS INTRODUCED BY THE MERGE!

**CRITICAL**: You MUST NOT ever kill bun processes! There may be other active bun processes doing work at the same time as you are that cannot be easily restarted if you kill them! In addition, you MUST NOT ever modify the user's global OpenCode configuration file located at ~/.config/opencode/opencode.json, to do so would destroy critical data!
