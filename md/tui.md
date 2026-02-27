Recent changes in dev (that we are now in sync with) mean that all the settings in the `tui` group in the `opencode.json` / `opencode.jsonc` file are now beaing being migrated into a new `tui.json` file, and this is complicating things for this feature branch, since after the settings that it adds get migrated to the `tui.json` file they are then not visible in the expected location to the feature.

To avoid this problem, lets move the `tui.$1` setting to `experimental.$1`. Find all the usages/references to this setting and come up with a plan to move it to the `experimental` group. Look at the existing examples of settings in the `experimental` group so see how these settings should be accessed. 

We **SHOULD** focus on making sure that the feature work with the setting stored at `experimental .$1`, we **MUST NOT** worry about fixing pre-existing type errors in def. 

When you have finished making your changes, you **MUST** commit and push the changes to origin with an appropriate commit message (using `--no-verify` if needed) - you **MUST NOT** leave any unstaged changes that have not been pushed to origin in the directory!