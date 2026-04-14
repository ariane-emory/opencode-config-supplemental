---
agent: plan
---
$ARGUMENTS

As per usual, we need to examine the branches involved and determine whether this is a regression that occurred during the merging process or whether there was a problem with one of the original feature/fix branches involved so that we can make a plan and fix it in the appropriate location.

If the problem was in one of the original branches, we'll need to check it out, fix ther problem there, and push that fixed version to origin,   then return to this integration brranch and merge the fixed branch back into this integration branch. If the regression occurred during the merging process, then we can simply fix here it in this integration branch. 