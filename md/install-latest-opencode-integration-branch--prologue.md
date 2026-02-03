**CRITICAL**: You MUST begin by fetching origin, or you're going to end up trying to use the wrong integration branch!

First, you MUST fetch origin and then check out the latest integration branch available on that remote. 

The integration branches all have the date/time in their branch names, so it should be easy for you to figure out which one is latest. If that branch already exists locally, do a 'git pull' just to make sure it's up to date. Make sure that the branch that you have selected to install actually exists on origin before you install it!

**CRITICAL**: You MUST NOT install an older integration branch! You MUST fetch first so that you can be sure you have correctly identified the latest one!

Once you have identified the newest integration branch, set the session title to "Installing $BRANCH_NAME", where $BRANCH_NAME is the name of the newest integration branch, and bookmark it.

Then, proceed to install it so I can run it globally from any directory.
