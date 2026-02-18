---
agent: build
---
Now you've got me worried that you might not have actually merged all the branches that I asked you to merge. Review the branch list file at `~/oc/md/branch-list.md` and the output of `git log` to ensure that you actually did merge all of the branches. 

You **MUST** merge all of the branches, the integration branch serves no purpose if it's missing the intended features/fixes!

The branches **MUST** be merged in the order in which they are listed in `~/oc/md/branch-list.md`. If you see the merge commit in a different order in the `git log`, you **MUST** rewind by force pushing back to the commit before the error occured and then **MUST** proceed with merging them in the correct order. **NO** deviations from the order in ``~/oc/md/branch-list.md` are permissible!
