**DIAGNOSTIC:** *In toto first:* "$1"
**DIAGNOSTIC:** *In toto second:* "$2"
**DIAGNOSTIC:** *In toto rest:* "${3..}"

!`baseone expand ~/ocs/md/regression--body.md $(gh --repo ariane-emory/opencode pr view $1 --json headRefName --jq .headRefName) $1 ${2..}`