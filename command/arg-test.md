---
ignored: true
---
All arguments: $ARGUMENTS
Second: ${2:default}
Default arg test: ${2..:default string}, ${3:$2:a secondary default}
Slice: ${2..3}
Iffy slice: ${2..3:a value}
