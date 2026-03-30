body all: $ARGUMENTS
body 1: $1
body 2: $2
body 3: $3
body 4: $4

First, you **MUST** run the tests in the ${4:dev} branch to observe any preexisting failures. You **SHOULD** remember these for later in this procedure. 

You **MUST** come up with a plan for how to implement a resolution for this Github issue in the $1 repository:

gh --repo $1 issue view $2;

You **SHOULD** take note of any coments on the issue, the discussion may include valuable insights on how to best resolve the issue:

gh --repo $1 issue view $2 --comments;

**IMPORTANT**: You **SHOULD** pay special attention to any comments from your user, ariane-emory, she often has good ideas!

You **MUST** think the changes required to resolve the issue through thoroughly and break the changes down into small steps in order to produce a detailed, step-by-step plan for resolving the issue. 

You **MUST** use the `set_current_session_title` tool to give the session a title following the format `issues|Resolving issue #$2: $ISSUE_TITLE`, substituting in the issue's title as appropriates, for example "issues|Resolving issue #$2: Add a foobar feature".

You **MUST NOT** make any unnecessary changes unrelated to the changes that are required to resolve the 

Think carefully about whether the changes necessitate re-generating the TypeScript SDK, and if so be sure to include this step in your plan. 

Group the plan's steps into "phases". After completing each phase, the code **MUST** build correctly and all tests (except for any pre-evisting errors that already existed on the the ${4:dev} branch and any enterprise tests, you can ignore those) **MUST** pass. 

Once you've come up with your plan, submit it to me for approval.

Once I have approved the plan, you **MUST** proceed with doing the work in a new branch based off of the ${4:dev} branch named $3.

Once you are done, you **MUST** push the new branch to origin, But you **MUST NOT** create any pull requests.