OUTER all: $ARGUMENTS
OUTER one: $1
OUTER two: $2
OUTER three: $3
OUTER four: $4

First, **MUST** take note of any pre-existing test failures in the $2 branch, you **MUST** remember these for later.

You **MUST** come up with a plan for how to implement a resolution for issue numbered #$2 in the $1 repository on GitHub:

gh --repo $1 issue view $2;

You **SHOULD** take note of any coments on the issue, the discussion may include valuable insights on how to best resolve the issue:

gh --repo $1 issue view $2 --comments;

You **MUST** use the `set_current_session_title` tool to give the session a title following the format `issues|Resolving issue #$2: $ISSUE_TITLE`, substituting in the issue's title as appropriate, for example "issues|Resolving issue #$2: Add a foobar feature".

**IMPORTANT**: You **SHOULD** pay special attention to any comments from your user, ariane-emory, she often has good ideas!

You **MUST** think the changes required to resolve the issue through thoroughly and break the changes down into small steps in order to produce a detailed, step-by-step plan for resolving the issue. 

You **MUST NOT** make any unnecessary changes unrelated to the changes that are required to resolve the isue.

Think carefully about whether the changes necessitate re-generating the TypeScript SDK, and if so be sure to include this step in your plan. 

Group the plan's steps into "phases". After completing each phase, the code **MUST** build correctly and all tests **MUST** pass (except for any preexisting changes from the $2 branch and the enterprise tests, you can ignore those).

Once you've come up with your plan, submit it to me for approval.

Once I have approved the plan, you **MUST** proceed with doing the work in a new branch based off of the $3 branch named $2.

You **SHOULD** use the `set_current_session_title` tool to give the session a name consisting of the string "issues|" followed by the name of the new branch followed by the number and title of the issue.

Once you are done, you **MUST** push the new branch to origin, But you **MUST** NOT create any pull requests.

${4..}