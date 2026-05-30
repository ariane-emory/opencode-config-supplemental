Each branch in the branch list that was provided has a PR number besides the branch name indicating which PR on the http://github.com/ariane-emory/opencode repository corresponds to that branch. 

You **MUST** run the `gh --repo ariane-emory/opencode pr view 1234` command (using 1234 as a placeholder PR number in that example) to examine each PR's description to see what feature/fix/change the branch we are merging is meant to implement.

You **SHOULD** also run `gh --repo ariane-emory/opencode pr diff 1234` command (again using 1234 as a placeholder PR number in that example) to see how this branch differs from the `dev` branch prior to performing the merge.

While merging this branch into the integration branch, you **MUST** ensure that you preserve the feature/fix described in the PR.