in-toto all: $ARGUMENTS
in-toto 1: $1
in-toto 2: $2
in-toto 3: $3

!`baseone expand ~/.config/opencode/md/resolve-gh-issue--prologue.md $1 $2`
!`baseone expand ~/.config/opencode/md/resolve-gh-issue--epilogue.md ${3..}`
