---
agent: plan
---
The feature / fix from the $1 branch looks like it may have been damaged during the merging process: I just tested it out and it does not currently appear to be working. 

We **MUST** figure out why it is not working and how it was broken.

If the original $1 branch was itself already broken, then we **MUST** switch to that branch, repair the feature / fix there, push that repair to origi, and then merge the  updated branch back into this branch.

If the damage is instead the result of an ill-thought-out change made while resolving merge conflicts, then we **SHOULD** fix it in place here in this branch and then push the repairs to this branch to origin. 