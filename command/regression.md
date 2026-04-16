---
agent: plan
ignored: true
---
There is a problem with the feature/fix from the $1 branch. 

${2..}

As per usual, we need to examine the branches involved and determine whether this is a regression that occurred during the merging process or whether there was a problem with the original feature/fix branch so that we can make a plan and fix it in the appropriate location.

If the problem was in the original feature/fix branch, we'll need to check it out, fix the problem there, push that fixed version to origin,   then return to this integration branch and merge that updated branch back into this integration branch. If instead the regression occurred during the merging process (perhaps a bad conflict resolution or the like), then we can simply fix here it in this integration branch. 

Please determine the correct branch to fix the problem in, make a plan to fix it, and proceed to do so. 