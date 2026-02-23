Think carefully about whether the changes necessitate re-generating the TypeScript SDK, and if so be sure to include this step in your plan. 

Group the plan's steps into "phases". After completing each phase, the code **MUST** build correctly and all tests (except the enterprise tests, you can ignore those) **MUST** pass. 

Once you've come up with your plan, submit it to me for approval.

Once I have approved the plan, you **MUST** proceed with doing the work in a new branch based off of the $1 branch named $2.

You **SHOULD** use the `set_current_session_title` tool to give the session a name consisting of the string "Issues:|" followed by the name of the new branch followed by the title of the issue.

Once you are done, you **MUST** push the new branch to origin, But you **MUST** NOT create any pull requests.

${3..}
