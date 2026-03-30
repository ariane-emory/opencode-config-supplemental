in-toto  all: "${..}"
in-toto 1: "$1"
in-toto 2: "$2"
in-toto 3: "$3"
in-toto 4: "$4"
in-toto 5: "$5"
in-toto 6: "$6"

!`baseone expand ~/.config/opencode/md/resolve-gh-issue--body.md ${1..3} ${4:dev} "${5..}"`
