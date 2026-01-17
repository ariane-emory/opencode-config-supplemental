---
description: Fix badly merged placeholder.
agent: build
---
Something seems to have gone wrong when merging the feat/sinister-quotes branch, the critical formatting change in the place holders that had been on line 887 of packages/opencode/src/cli/cmd/tui/component/prompt/index.tsx appears to have been lost, it's still showing the "Ask anything" part and the quotes around the placeholder.
