### This text is from the new merge command in toto. 

**SYSTEM:** *In toto first:* "$1"
**SYSTEM:** *In toto second:* "$2"
**SYSTEM:** *In toto rest:* "${2..}"

!`git reset --hard HEAD > /dev/null 2>&1 && echo "**SYSTEM:** *Reset repo state.*"; `
!`git clean -fd > /dev/null 2>&1 && echo "**SYSTEM:** *Cleaned untracked files.*";`
!`git checkout $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) > /dev/null 2>&1 && echo "**SYSTEM:** *Checked out the $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) branch.*"`

!`baseone expand ~/ocs/md/merge-dev-into-branch-by-pr--body.md $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) $1 ${2..}`