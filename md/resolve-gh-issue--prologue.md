You **MUST** come up with a plan for how to implement a resolution for this Github issue in the $1 repository: .

gh --repo $1 issue view $2;

You **SHOULD** take note of any coments on the issue, the discussion may include valuable insights on how to best resolve the issue:

gh --repo $1 issue view $2 --comments;

**IMPORTANT**: You **SHOULD** pay special attention to any comments from your user, ariane-emory, she often has good ideas!

You **MUST** think the changes required to resolve the issue through thoroughly and break the changes down into small steps in order to produce a detailed, step-by-step plan for resolving the issue. 

You **MUST** use the `set_current_session_title` tool to give the session a title following the format `issues|Resolving issue #$ISSUE_NUMBER: $ISSUE_TITLE`, substituting in the issue's number and title as appropriates, for example "issues|Resolving issue #1234: Add a foobar feature".

You **MUST NOT** make any unnecessary changes unrelated to the changes that are required to resolve the isue.

