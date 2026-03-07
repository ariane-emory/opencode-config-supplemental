You **MUST** set the title of the current session to "integrations|" followed by the name of the integration branch (for example "integrations|integraton/1970-01-01-12-34").

**CRITICAL**: After creating the integration branch, you **MUST** configure it to track origin (not upstream) to prevent push issues:

```fish
# Replace BRANCH-NAME with the actual integration branch name
set BRANCH_NAME integration/(date +%Y-%m-%d-%H-%M)
git config branch.$BRANCH_NAME.remote origin
git config --unset branch.$BRANCH_NAME.merge
```

This prevents git from trying to push to upstream when integration branches don't exist there.

**VERIFICATION STEP**: After creating and configuring the integration branch, verify the configuration:

```fish
git config --list | grep branch.integration/
# Should show: 
# branch.integration/YYYY-MM-DD-HH-MM.remote=origin
# Should NOT show any .merge configuration
```

**CRITICAL**: You **MUST NOT** use git cherry-pick, that always destroys the very feature/fix we're trying to merge.

**CRITICAL - MERGE STRATEGY**: Stay on the integration branch and merge each feature branch into it:
```fish
# CRITICAL: Verify the branch exists locally before merging
git branch --list feat/branch-name
# If no output, the branch doesn't exist - DO NOT PROCEED without the branch!

# Stay on integration branch - do NOT checkout the feature branch
git merge feat/branch-name --no-ff -m "Merge feat/branch-name"
```

**BRANCH EXISTENCE CHECK**: Before each merge, verify the branch exists:
```fish
# Check if branch exists (run this for each branch before merging)
if not git show-ref --quiet refs/heads/feat/branch-name
    echo "ERROR: Branch feat/branch-name does not exist locally!"
    echo "You MUST have all branches available locally before starting."
    exit 1
end
```

**UNDERSTANDING "OURS" vs "THEIRS" IN MERGES:**
When merging a feature branch INTO the integration branch:
- `ours` = integration branch (the base you're merging INTO)
- `theirs` = feature branch (the one being merged)

**CRITICAL**: You **SHOULD** merge only from local copies of the branches, you **MUST NOT** pull or fetch unless you are explicitly instructed to do so by the user.

Do **NOT** switch away from the integration branch during the merge process.

**CRITICAL**: You **MUST** push the integration branch to origin after merging **EACH** branch into it!

These are the branches we need to merge into the new integration branch:

### Step 1: Create MERGED-BRANCHES.md Checklist

**CRITICAL**: Before merging ANY branches, you MUST create MERGED-BRANCHES.md as a living checklist:

**First, determine the branch count dynamically:**
```fish
set BRANCH_COUNT (grep "^- " ~/.config/opencode/md/branch-list.md | wc -l)
echo "Creating checklist for $BRANCH_COUNT branches"
```

**Create the file with this exact format:**

```markdown
# Integration Branch: integration/YYYY-MM-DD-HH-MM

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
- Copy ALL branches from branch-list.md into the checklist
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

**DO NOT proceed to merging until the checklist is created and committed!**
