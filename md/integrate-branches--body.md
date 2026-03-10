!`baseone expand ~/.config/opencode/md/rfc.md`

## Instructions

We need to merge several feature branches into a new integration branch for testing.

The directory that you're in now is used only for assembling single-use integration branches for testing purposes, any local changes that you find here are unimportant and **MUST** be discarded.

Using the dev branch as the starting point, you **MUST** start a new integration branch with explicit tracking to origin to avoid upstream tracking issues. You **MUST** start a new branch, do not try to start from any prior integration branch! Name this new branch: $1

You **MUST** set the title of the current session to "integrations|$1" ("integrations|" followed by the name of the integration branch).

**CRITICAL**: After creating the integration branch, you **MUST** configure it to track origin (not upstream) to prevent push issues:

```fish
# Replace BRANCH-NAME with the actual integration branch name
set BRANCH_NAME integration/(date +%Y-%m-%d-%H-%M)
git config branch.$1.remote origin
git config --unset branch.$1.merge
```

This prevents git from trying to push to upstream when integration branches don't exist there.

**VERIFICATION STEP**: After creating and configuring the integration branch, verify the configuration:

```fish
git config --list | grep branch.integration/
# Should show: 
# branch.$1.remote=origin
# Should NOT show any .merge configuration
```

**CRITICAL**: You **MUST NOT** use git cherry-pick, that always destroys the very feature/fix we're trying to merge.

**CRITICAL - MERGE STRATEGY**: Stay on the integration branch and merge each feature branch into it **ONE BY ONE**:

```fish
# CRITICAL: Verify the branch exists locally before merging
git branch --list feat/branch-name
# If no output, the branch doesn't exist - DO NOT PROCEED without the branch!

# Stay on integration branch - do NOT checkout the feature branch
git merge feat/branch-name --no-ff -m "Merge feat/branch-name"
```

**CRITICAL**: You **MUST** merge only from local copies of the branches, you **MUST NOT** pull or fetch!

Do **NOT** switch away from the integration branch during the merge process.

**CRITICAL**: You **MUST** push the integration branch to origin after merging **EACH** branch into it!

### The branches that you MUST merge:

These are the branches you **MUST** merge into the new integration branch:

!`baseone expand ~/.config/opencode/md/branch-list.md`

### Step 1: Create MERGED-BRANCHES.md Checklist

**CRITICAL**: Before merging **ANY** branches, you **MUST** create MERGED-BRANCHES.md as a living checklist:

**Create the file with this exact format:**

```markdown
m# Integration Branch: integration/YYYY-MM-DD-HH-MM

## Merge Checklist

| Status | # | Branch Name | Remote | Commit Hash | Description |
|--------|---|-------------|--------|-------------|-------------|
| ☐ | 1 | split-config-fixes | upstream | TBD | MUST use only local copy, do NOT pull from upstream |
| ☐ | 2 | feat/base-one-rebrand | origin | TBD | |
| ☐ | 3 | feat/sinister-quotes | origin | TBD | Placeholders MUST be SINISTER_PLACEHOLDERS array |
| ☐ | 4 | feat/markdown-renderer | gignit | TBD | |
| ☐ | 5 | feat/thinking-indicator-hidden | rcdailey | TBD | |
| ☐ | 6 | fix/session-new-prompt-handoff | AksharP5 | TBD | |
| ☐ | 7 | feat/session-grouping | origin | TBD | |
| ☐ | 8 | feat/session-bookmarks | origin | TBD | |
| ☐ | 9 | fix/dialog-datetime-alignment | origin | TBD | Merge immediately after feat/session-bookmarks |
| ... | ... | ... | ... | ... | ... |
```

**REQUIREMENTS:**
- Copy **ALL** branches from branch-list.md into the checklist
- Include merge advice from branch-list.md in the "Description" column
- Use ☐ for unchecked (starts as this for ALL branches)
- Commit Hash starts as "TBD", updated after each merge
- Number sequentially from 1 to $BRANCH_COUNT
- **CRITICAL**: Do NOT mark any as ☑ yet - you're creating this BEFORE merging

**VERIFICATION**: After creating, run:

```fish
# Count unchecked branches - MUST equal branch-list count
grep "^| ☐ |" MERGED-BRANCHES.md | wc -l
# Compare to:
grep "^- " ~/.config/opencode/md/branch-list.md | wc -l
# These two numbers MUST match!
```

**Commit the initial MERGED-BRANCHES.md:**

```fish
git add MERGED-BRANCHES.md
git commit -m "Initial MERGED-BRANCHES.md checklist"
git push
```

**You MUST NOT proceed to merging until the checklist is created and committed!**

### Step 1b: Create Todo List for Progress Tracking

**CRITICAL**: Immediately after creating MERGED-BRANCHES.md, you **MUST** create individual todo list items for **EVERY** branch:

```fish
# Get the branch count
set BRANCH_COUNT (grep "^- " ~/.config/opencode/md/branch-list.md | wc -l)
echo "Creating $BRANCH_COUNT todo items..."
```

Use the todowrite tool to create individual todo items for each branch. This allows the user to monitor progress. Each todo should:

- Include the branch name and remote
- Mark branches with known conflicts as "in_progress" or include a note
- Be updated to "completed" immediately after successful merge AND push

**WHY THIS MATTERS**: The todo list provides real-time visibility into progress. Without it, the user cannot easily see which branches remain to be merged.

**DO NOT start merging branches until the todo list is created!**

### Step 1c: Branch Verification Checklist

Before starting merges, run these verification commands:

```fish
# Verify all branches exist locally
echo "Verifying all branches exist..."
for branch in (grep "^- " ~/.config/opencode/md/branch-list.md | sed 's/^- //')
    if not git branch --list $branch | grep -q $branch
        echo "ERROR: Branch $branch does not exist locally!"
        echo "You MUST fetch all branches before starting."
        exit 1
    end
end
echo "All branches verified!"
```

**CRITICAL**: If any branch is missing, STOP and ask the user for guidance. Do NOT proceed with incomplete branches.

### Step 1d: Pre-Merge Setup Verification

Before merging the first branch, verify:

1. **Session title is set** to "integrations|integration/YYYY-MM-DD-HH-MM"
2. **Integration branch is created** from dev with tracking to origin
3. **MERGED-BRANCHES.md exists** with all branches listed
4. **Todo list is created** with individual items for each branch
5. **All branches verified** to exist locally
6. **Typecheck passes** on the integration branch baseline

**DO NOT proceed if any of these are not complete!**

- you **MUST NOT** include/merge any branches not on this list
- you **MUST** process the branches that I listed in the order that I listed them.
- all of the remotes these branches are on should already be configured, you **MUST NOT** ever add new remotes.

### Pre-Integration Verification

**MANDATORY: Before starting, you MUST count the branches and record the expected total:**

```fish
# Count branches in branch-list.md
grep "^- " ~/.config/opencode/md/branch-list.md | wc -l
# Record this number - you MUST merge exactly this many branches!
```

**CRITICAL**: Before starting, you MUST determine the total branch count dynamically:

```fish
# Count branches dynamically - DO NOT hardcode this number
set BRANCH_COUNT (grep "^- " ~/.config/opencode/md/branch-list.md | wc -l)
echo "Total branches to merge: $BRANCH_COUNT"
```

**MANDATORY**: Record this number. You MUST merge exactly $BRANCH_COUNT branches. If at any point your progress shows a different total, you have a bug in your counting.

Before starting, verify the branch list format to avoid common errors:

```fish
# Check for concatenated branches (multiple branch names on one line)
grep "\- feat/" ~/.config/opencode/md/branch-list.md | grep "feat/.*feat/"
# If this outputs anything, branches are concatenated and must be split into separate lines

# Count total branches dynamically
grep "^\- " ~/.config/opencode/md/branch-list.md | wc -l
# Record this number - you will need it for MERGED-BRANCHES.md
```

**CRITICAL**: Each branch **MUST** be on its own line starting with `- `. If you see multiple branch names on one line (e.g., "feat/branch1- feat/branch2"), split them into separate lines before proceeding.

**IMPORTANT**: You **MUST** create a MERGED-BRANCHES.md document in the project's root directory in which to record which branches were merged to produce the new integration branch. You **MUST** include a Markdown table displaying which branches were merged in this document and a merge log detailing the merges that were performed. Make sure that you include the specific commit hash the merged branch was at when it was merged. Add this file to git.

#### MERGED-BRANCHES.md Required Format

The document **MUST** contain a Markdown table (NOT a bullet list) with checkboxes for tracking. Use this exact format:

```markdown
| Status | # | Branch Name | Remote | Commit Hash | Description |
|--------|---|-------------|--------|-------------|-------------|
| ☐ | 1 | split-config-fixes | upstream | TBD | MUST use only local copy, do NOT pull from upstream |
| ☐ | 2 | feat/base-one-rebrand | origin | TBD | |
| ☐ | 3 | feat/sinister-quotes | origin | TBD | Placeholders MUST be SINISTER_PLACEHOLDERS array |
| ... | ... | ... | ... | ... | ... |
```

**NOT** this format:
```markdown
- split-config-fixes
- feat/branch-name
```

The table **MUST** include:
- **Status column**: Use ☐ for unchecked, ☑ for checked
- Sequential numbering (#)
- Full branch name
- Remote source (upstream, origin, gignit, etc.)
- Commit Hash: Start as "TBD", update after each merge
- Merge Advice: Copy relevant advice from branch-list.md

**CRITICAL**: Create ALL rows with ☐ status before merging any branches. The checklist serves as your source of truth for what remains to be done.

**CRITICAL - MERGED-BRANCHES.md AS CHECKLIST**: MERGED-BRANCHES.md is your PRIMARY tracking mechanism. Before merging ANY branches:

1. Create MERGED-BRANCHES.md with ALL branches listed as ☐ (unchecked)
2. Copy merge advice from branch-list.md into the Description column
3. After each merge, immediately update the file:
   - Change ☐ to ☑ 
   - Replace "TBD" with the actual commit hash
   - Add brief merge notes if there were conflicts
4. Commit and push MERGED-BRANCHES.md after EACH update

This persistent checklist ensures no branches are skipped and survives session restarts.

**CRITICAL - NEVER BATCH MERGE**: You **MUST NOT** try to process branches in batches or in parallel. While it may seem efficient, batch merging causes serious problems:

- **Merge conflicts compound**: When multiple branches modify the same files, batch merging creates overlapping conflicts that are much harder to resolve correctly
- **Lost work**: If a batch merge fails partway through, you may lose track of which branches were actually merged
- **Typecheck failures**: Batch merging multiple branches makes it impossible to determine which branch introduced a type error
- **Incomplete merges**: The `MERGE_HEAD` variable gets overwritten, making it impossible to see what each branch contributed

**YOU MUST**: Process branches individually, one at a time, in the exact order listed in branch-list.md.

**CORRECT WORKFLOW**:
```fish
# For EACH branch:
git merge feat/branch-name --no-ff -m "Merge feat/branch-name"
# If conflicts: resolve them completely
bun turbo typecheck
# Update MERGED-BRANCHES.md with commit hash
git add -A && git commit -m "Update MERGED-BRANCHES.md: checked off branch-name"
git push
# Update todo list to mark branch as completed
# THEN and ONLY THEN move to the next branch
```

**CRITICAL - NEVER REBASE THE INTEGRATION BRANCH**: Do **NOT** use `git pull --rebase` or `git rebase` on the integration branch. Rebasing will clobber merge commits and lose changes from previously merged branches. If there are conflicts with the remote integration branch, use `git pull --no-rebase` or `git push --force` since this is a single-use integration branch for testing purposes.

**CRITICAL BRANCH HYGIENE**: Throughout this entire process, you **MUST** remain on the integration branch. **DO NOT** switch to `dev` or any other branch until the very end when performing the post-integration verification. All merges, commits, and pushes should happen while on the integration branch. The `dev` branch should NEVER receive any of the integration commits.

For each of these branches, you **SHOULD**:

- Merge the branch while staying on the integration branch: `git merge branch-name --no-ff`
- Try your very best to carefully resolve any conflicts that occur. When resolving conflicts, think them through carefully and thoroughly. When multiple branches modify the same file, make sure to incorporate changes from **BOTH** sides - don't just pick one side. Make sure not to let content from dev clobber content from the fix/feature branches that we're trying to merge in! Pay special attention for coments labelled with '**CRITICAL**', these are used to indicate specific changes that **MUST NOT** be clobbered!

### Conflict Resolution Protocol

When you encounter merge conflicts, follow this protocol to ensure you properly incorporate changes from **BOTH** branches:

**Step 1: Understand what each branch contributed**
```fish
# See what the integration branch (base) has:
git show HEAD:path/to/conflicting/file > /tmp/ours.txt

# See what the feature branch has:
git show MERGE_HEAD:path/to/conflicting/file > /tmp/theirs.txt

# Compare them:
diff -u /tmp/ours.txt /tmp/theirs.txt
```

**Step 2: Resolve by combining, not choosing**
```
<<<<<<< HEAD
    code from integration branch
=======
    code from feature branch
>>>>>>> feat/branch-name
```

**DO NOT** just delete one side. Instead, manually edit to include **both** functionalities:
- Keep the integration branch's structure/context
- Add the feature branch's new functionality
- Ensure the result compiles and makes logical sense

**Step 3: Verify your resolution**
```fish
# Check what you resolved:
git diff --cached path/to/conflicting/file

# Verify key functions from BOTH branches are present:
grep -n "functionFromIntegration" path/to/conflicting/file
grep -n "functionFromFeature" path/to/conflicting/file
```

**Step 4: Run typecheck (MANDATORY - NO EXCEPTIONS)**
After EVERY merge (even clean merges without conflicts), you MUST run typecheck:

```fish
bun turbo typecheck
```

**WHY THIS IS CRITICAL**:
- Type errors may not show up in git merge output
- Conflicts can introduce subtle type mismatches even when resolved
- Multiple branches modifying the same types can create incompatible combinations

**IF TYPECHECK FAILS**:
1. Do NOT proceed to the next branch
2. Fix the type errors first
3. Commit the fixes
4. Re-run typecheck until it passes
5. Only then update MERGED-BRANCHES.md and push

**NEVER skip typecheck after a merge, even if the merge was "clean"!**

**Step 5: If unsure, preserve more rather than less**
When in doubt, keep both versions and add a TODO comment. It's easier to clean up later than to recover lost functionality.

- **CRITICAL VERIFICATION**: After each merge, you **MUST** verify that key changes from the branch are actually present in the integration branch. Check for specific functions, types, or code patterns that the branch was supposed to introduce. A merge that reports "Already up to date" may indicate the changes were **NOT** actually merged.

**CONFLICT RESOLUTION VERIFICATION**: If a merge had conflicts, you **MUST** additionally verify:
1. Key functionality from the integration branch is still present (nothing was accidentally lost)
2. Key functionality from the feature branch is now present (the merge succeeded)
3. Run `git log --oneline -5` to confirm the merge commit exists
4. Run the application or tests to ensure the merged code actually works together
5. **Run `bun turbo typecheck` and verify it passes**
- Record the result of handling this branch in MERGED-BRANCHES.md.
- Update MERGED-BRANCHES.md immediately (see update instructions below).

Since you may need to read files to resolve conflicts, you must always use the Read tool on files prior to using the Edit tool on them to avoid errors.

If a git lock file gets in your way, you **SHOULD** just delete it and keep working on merging.

**MANDATORY TYPECHECK AFTER EACH MERGE**: The typecheck command `bun turbo typecheck` **MUST** pass after merging **EACH** branch before you push or proceed to the next branch. This is non-negotiable - even if a merge appears clean, type errors can be introduced.

For each of the feature branches, if you are able to merge in the changes into the integration branch and successfully resolve any conflicts and the tests all pass afterwards, you **MUST** push the changes to the integration branch on origin on git before proceeding on to the next feature branch.

There is a 'test' pre-pushhook in this repository, so it is possible a push may fail with an error message. If this occurs, you **MUST** attempt fix the error and try again until the test passes. If, after trying your best, you still cannot fix the error, you may use `git push --no-verify` to bypass the hook.

### Step 2: Update MERGED-BRANCHES.md After Each Merge

**CRITICAL**: **IMMEDIATELY** after successfully merging a branch (and before moving to the next), you **MUST**:

1. Get the commit hash:
```fish
git log --oneline -1 | cut -d' ' -f1
```

2. Update MERGED-BRANCHES.md:
   - Find the row for the branch you just merged
   - Change ☐ to ☑ in the Status column
   - Replace "TBD" with the actual commit hash
   - Add brief merge notes in Description if there were conflicts

3. **MANDATORY**: Run this verification after each update:
```fish
# Count remaining unchecked branches
grep "^| ☐ |" MERGED-BRANCHES.md | wc -l
echo "Remaining branches to merge: [count]"
```

4. Commit and push the updated MERGED-BRANCHES.md:
```fish
git add MERGED-BRANCHES.md
git commit -m "Update MERGED-BRANCHES.md: checked off [branch-name]"
git push
```

5. **Update the todo list** to mark the branch as completed:
   - Use the todowrite tool to mark the branch's todo item as "completed"
   - This provides real-time visibility for the user

**WHY THIS MATTERS**: If you don't update immediately, you may lose track of which branches are done! The ☐ to ☑ progression is your primary tracking mechanism. The todo list provides immediate visual feedback to the user about progress.

**CRITICAL: PROGRESS TRACKING**: After merging EACH branch, announce your progress using the dynamic count:
```fish
set MERGED_COUNT (grep "^| ☑ |" MERGED-BRANCHES.md | wc -l)
set TOTAL_COUNT (grep "^| ☐ |" MERGED-BRANCHES.md | wc -l)
set TOTAL_COUNT (math $TOTAL_COUNT + $MERGED_COUNT)
echo "Merged $MERGED_COUNT of $TOTAL_COUNT branches: [branch-name]"
```
This helps verify no branches are being skipped!

**CRITICAL**: You must proceed until you have merged ALL branches! Only if you've tried everything you can think of to resolve a conflict without success and are nonetheless unable to resolve the conflict or the tests do not pass afterwards, do not push the changes to git, in this case you **MUST** stop and ask me to step in and help you out instead. Only do this is you've tried everything you can think of! Otherwise, proceed until MERGED-BRANCHES.md shows ZERO unchecked branches (grep "^| ☐ |" returns 0).

Ultrathink! 

## Verifying Merge Results

After each merge, you **MUST** verify the merge was successful by checking that the branch's key changes are present. Do **NOT** rely solely on git's merge output - "Already up to date" or a fast-forward merge may indicate the changes were **NOT** properly integrated.

**Verification procedure for each branch:**
1. Before merging, check the PR diff or branch diff to identify key changes (new functions, types, file modifications)
2. After merging, grep for those specific changes in the integration branch
3. If key changes are missing, re-merge using `git merge branch-name --no-ff`

**Example verification:**
```fish
# Before merge: identify what the branch adds by examining the branch directly
git show branch-name:path/to/file.ts
# or
git diff dev...branch-name --stat

# After merge: verify the changes exist in integration branch
grep -rn "newFunctionName" packages/
grep -rn "NewTypeName" packages/
```

**Common merge pitfalls to watch for:**
- "Already up to date" - The branch may have already been merged, or you may be merging a stale local branch
- Fast-forward merges when there should be changes - Use `--no-ff` flag and verify changes
- **Choosing one side in conflicts** - The most common error! You **MUST NOT NEVER** just pick `--ours` or `--theirs`! **YOU MUST LAWAYS** manually combine both sides' functionality! The `--ours` and `--theirs` switches are **ABOLUTELY FORBIDDEN** and **MUST NOT** ever be used!
- Conflicts that silently drop one side - **ALWAYS** review conflict resolutions carefully, especially for overlapping changes
- Multiple branches modifying the same file - Later merges may need to incorporate changes from **BOTH** the integration branch **AND** the feature branch
- **REBASING THE INTEGRATION BRANCH** - This is the most dangerous pitfall! Never rebase an integration branch as it will destroy merge commits and lose changes

## Finishing touches

Modify the configuration file at `./.opencode/tui.json` to include this property to set my preferred theme:

```
{
  "$schema": "https://opencode.ai/tui.json",
  "theme": "matrix--transparent",
  // remainder of content left as-is.
}
```

**NOTE**: Note that the theme property just at `theme`, **NOT** `tui.theme`!

## Set Integration Branch Version

Update the VERSION constant in `packages/opencode/src/installation/index.ts` to display the integration branch name instead of "local". Find this line:

```typescript
export const VERSION = typeof OPENCODE_VERSION === "string" ? OPENCODE_VERSION : "local"
```

And change `"local"` to the integration branch date:

```typescript
export const VERSION = typeof OPENCODE_VERSION === "string" ? OPENCODE_VERSION : "YYYY-MM-DD-HH-MM"
```

(Replace `YYYY-MM-DD-HH-MM` with the actual branch timestamp, e.g., `2025-12-17-01-17`)

**CRITICAL**: The version **MUST** match the 'YYYY-MM-DD-HH-MM' format, and must not contain the string 'opencode' or a version number like '0.0.0'. The displayed version **MUST NEVER** be computed dynamically! The datetime string **MUST NOT** be treated as a 'fallback value', it **MUST** be strictly hardcoded to this value in all possible situations(whether using 'bun dev' or running a compiled binary). This requires strictly hardcoding the value in  packages/opencode/script/build.ts: the VERSION **MUST NOT** be dynamically constructed! Remember, we will **NEVER** build binaries from this integration branch, the hardcoded version string **MUST ALWAYS** display correctly when it is launched using `bun dev`!

The version **MUST NOT** ever be computed dynamically!

**EXPLICIT EXAMPLES**:

✅ **CORRECT** - Hard-coded literal string:
```typescript
export const VERSION = typeof OPENCODE_VERSION === "string" ? OPENCODE_VERSION : "2026-03-06-01-51"
```

❌ **WRONG** - Dynamic detection (**MUST NOT DO THIS!**):
```typescript
// NEVER do this - dynamic detection is forbidden
export const VERSION = typeof OPENCODE_VERSION === "string" ? OPENCODE_VERSION : detectBranchTimestamp()

// NEVER do this - fallback to "local" is forbidden on integration branches
export const VERSION = typeof OPENCODE_VERSION === "string" ? OPENCODE_VERSION : "local"
```

The string `"2026-03-06-01-51"` (or whatever the actual timestamp is) **MUST** appear as a literal in the source code. It **MUST NOT** be constructed, computed, or detected at runtime.

This ensures the TUI sidebar shows which integration branch is running when using the development wrapper script. Commit this change along with the other finishing touches.

### Fix SemVer Validation for Non-SemVer Version Strings

The `YYYY-MM-DD-HH-MM` format is not a valid SemVer string. The codebase has a function `isOutdated()` in `packages/opencode/src/bun/registry.ts` that uses Bun's `semver.order()` which will throw an error for invalid SemVer strings.

**You **MUST** modify `packages/opencode/src/bun/registry.ts`** to wrap the semver operations in a try-catch block:

```typescript
export async function isOutdated(pkg: string, cachedVersion: string, cwd?: string): Promise<boolean> {
  const latestVersion = await info(pkg, "version", cwd)
  if (!latestVersion) {
    log.warn("Failed to resolve latest version, using cached", { pkg, cachedVersion })
    return false
  }

  try {
    const isRange = /[\s^~*xX<>|=]/.test(cachedVersion)
    if (isRange) return !semver.satisfies(latestVersion, cachedVersion)
    return semver.order(cachedVersion, latestVersion) === -1
  } catch {
    log.warn("Invalid semver format, assuming not outdated", { pkg, cachedVersion })
    return false
  }
}
```

This ensures that custom version strings like `"2026-02-15-13-46"` won't cause the application to crash with an "Invalid SemVer" error.

### Prevent Package.json Pollution

When CHANNEL is **NOT** "local", the `installDependencies()` function in `packages/opencode/src/config/config.ts` writes the VERSION string to `.opencode/package.json`. Since `YYYY-MM-DD-HH-MM` is not a valid npm version, this pollutes the config and breaks both the integration build **AND** the original opencode installation.

You **MUST** also set CHANNEL to "local"** in `packages/opencode/script/build.ts`:

```typescript
// Change these lines:
OPENCODE_CHANNEL: `'${Script.channel}'`,
// To:
OPENCODE_CHANNEL: `'local'`,
```

## Troubleshooting Remote Tracking Issues

**Problem**: `git push` tries to push to upstream instead of origin with error:
```
fatal: The upstream branch of your current branch does not match the name of your current branch
```

**Solution**: Fix branch configuration:
```fish
# Replace BRANCH-NAME with the actual integration branch name
set BRANCH_NAME integration/YYYY-MM-DD-HH-MM
git config branch.$BRANCH_NAME.remote origin
git config --unset branch.$BRANCH_NAME.merge
```

**Prevention**: Always verify branch configuration after creating integration branches:
```fish
git config --list | grep branch.integration
# Ensure .remote=origin and no .merge setting exists
```

## Post-Integration Verification

**CRITICAL**: After completing all integration work, verify that the `dev` branch was never modified (it should still match `origin/dev` exactly):

```fish
# Verify dev still matches origin/dev exactly (without switching to it)
git log origin/dev..dev --oneline
# **MUST** show nothing (empty output)
# If this shows commits, something went wrong - dev should never move during integration
```

**Why this matters**: The integration process should **NEVER** touch `dev`. All work happens on the integration branch. If `dev` has moved forward, it means the integration commits were accidentally committed to the wrong branch.

**What to do if dev has moved forward (this indicates an error occurred)**:
This should not happen if the instructions were followed correctly. If it does happen:
1. **STOP** and inform the user that `dev` was accidentally modified
2. The user will need to decide whether to:
   - Reset `dev` back to `origin/dev` (losing those commits from `dev`)
   - Move those commits to the integration branch if they're missing there

### Verify All Branches Merged

**CRITICAL FINAL CHECK**: Before proceeding to finishing touches, ensure all branches are checked off:

```fish
# Count unchecked branches - MUST be 0
grep "^| ☐ |" MERGED-BRANCHES.md | wc -l
```

If this outputs anything other than 0, you have NOT completed all merges! Go back and merge the remaining branches.

**Also verify:**
```fish
# Count checked branches
grep "^| ☑ |" MERGED-BRANCHES.md | wc -l
# This should equal: grep "^- " ~/.config/opencode/md/branch-list.md | wc -l

# If they don't match, investigate which branches are missing
```

**DO NOT proceed to finishing touches until all branches are ☑ checked!**

### Final Verification Commands

Run these commands to verify the integration is complete and correct:

```fish
# 1. Verify MERGED-BRANCHES.md has proper Markdown table format
grep "^| # |" MERGED-BRANCHES.md
# **MUST** output: | # | Branch Name | Remote | Commit Hash | Description |

# 2. Verify version is hard-coded literal (not dynamic)
grep "export const VERSION" packages/opencode/src/installation/index.ts
# **MUST** show: "YYYY-MM-DD-HH-MM" as a literal string in quotes
# **MUST NOT** show: function calls, variables, or "local"

# 3. Verify all branches were merged (unchecked count should be 0)
grep "^| ☐ |" MERGED-BRANCHES.md | wc -l
# **MUST** output: 0 (all branches checked off)

# Also verify count matches branch-list.md
grep "^| ☑ |" MERGED-BRANCHES.md | wc -l
# This should equal: grep "^- " ~/.config/opencode/md/branch-list.md | wc -l

# 4. Verify no concatenated branches were missed
grep "\- feat/" ~/.config/opencode/md/branch-list.md | grep "feat/.*feat/"
# **MUST** output nothing (empty). If it outputs anything, branches were concatenated.

# 5. Verify dev branch wasn't touched
git log origin/dev..dev --oneline
# **MUST** show nothing (empty output)
```

**Final checklist**:
- [ ] **ALL** branches from branch-list.md have been merged: `grep "^| ☐ |" MERGED-BRANCHES.md | wc -l` **MUST** output 0
- [ ] Integration branch has been pushed to origin
- [ ] MERGED-BRANCHES.md including Markdown table is complete and committed
- [ ] Tests pass on the integration branch
- [ ] Run `bun dev` to confirm that the program will actually start
- [ ] Verified: `dev` branch still matches `origin/dev` (hasn't moved forward)
- [ ] Currently on the integration branch with clean working tree 

**IMPORTANT**: You **MUST** complete the **ENTIRE** plan and merge ALL of the requested branches into the new integration branch!

**CRITICAL**: Don't forget the instructions above about hardcoding the value of VERSION!

Make sure that you commit and push all your changes when you are done. 

**REMEMBER**: You **MUST** base the integration branch off of the local dev branch, and you **MUST** push the integration branch after every single branch that you merge. 

**CRITICAL**:  You **MUST NOT** make use of subagents to perform any of this work, subagents are **ABSOLUTELY FORBIDDEN**!

**DON'T BE LAZY!** You **MUST NOT** stop until you have completed the **ENTIRE** task and merged **ALL** of the requested branches!