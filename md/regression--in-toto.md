**DIAGNOSTIC:** *In toto first:* "$1"
**DIAGNOSTIC:** *In toto second:* "$2"
**DIAGNOSTIC:** *In toto rest:* "${2..}"

!`git checkout $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) > /dev/null 2>&1 && echo "**SYSTEM:** *Checked out the $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) branch.*"`

!`baseone expand ~/ocs/md/regression--body.md $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) $1 ${2..}