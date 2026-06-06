**DIAGNOSTIC:** *In toto first:* "$1"
**DIAGNOSTIC:** *In toto second:* "$2"
**DIAGNOSTIC:** *In toto rest:* "${2..}"

!`git reset --hard HEAD > /dev/null 2>&1 && echo "**SYSTEM:** *Reset repo state.*"; `
!`git clean -fd > /dev/null 2>&1 && echo "**SYSTEM:** *Cleaned untracked files.*";`
!`git fetch > /dev/null 2>&1 && echo "**SYSTEM:** *Ran git fetch.*"; `
!`git checkout $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) > /dev/null 2>&1 && echo "**SYSTEM:** *Checked out the $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) branch.*"`
!`git checkout dev > /dev/null 2>&1 && echo "**SYSTEM:** *Checked out the dev branch.*"`


!`baseone expand ~/ocs/md/merge-dev-into-branch-by-pr-resolve-conflicts-and-push--body.md $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) $1 ${2..}`