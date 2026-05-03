outer body all: "${..}"
outer body 1: "$1"
outer body 2: "$2"
outer body 3: "$3"
outer body 4: "$4"
outer body 5: "$5"
outer body 6: "$6"

!`baseone expand ~/.config/opencode/md/resolve-gh-issue--body.md ${1..3} ${4:dev} "${5..}"`
