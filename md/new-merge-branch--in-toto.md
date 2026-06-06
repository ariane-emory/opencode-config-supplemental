This is the new merge command in toto. 

In toto first: "$1"
In toto second: "$2"
In toto rest: "${2..}"

!`git reset --hard HEAD > /dev/null 2>&1 && echo "SYSTEM: Reset repo state."; `
!`git clean -fd > /dev/null 2>&1 && echo "SYSTEM: Cleaned untracked files.";`
!`baseone expand ~/ocs/md/new-merge-branch--body.md $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) $1 ${2..}`