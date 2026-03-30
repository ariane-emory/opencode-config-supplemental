---
description: Plan to resolve a GH issue with options.
agent: plan

---

Let's try to come up with a plan for how to implement a resolution for this Github issue in the sst/opencode repository: 

gh --repo anomalyco/opencode issue view $1;

Take note of any coments on the issue, the discussion may include valuable insights on how to best resolve the issue:

gh --repo anomalyco/opencode issue view $1 --comments;

IMPORTANT: Pay special attention to any comments from your user, ariane-emory, she often has good ideas!

Think the changes required to resolve the issue through thoroughly and break the changes down into small steps in order to produce THREE detailed, step-by-step plans for resolving the issue. 

Each plan SHOULD make the minimal changes required to solve the issue. Do not make any unnecessary changes unrelated to the changes that are required to resolve the isue.

Once you have come up with three options, evaluate their pros and cons and present the options to me to give me an opportunity to choose between them, along withyour comments on the relative merits of each option. Present these options as a numbered list.

You MUST come up with three different plans for how this issue could be resolved and you MUST present these options to me to choose from.

Once you've come up with your plan, submit it to me for approval.

Once I have approved the plan, proceed with doing the work in a new branch based off of the $1 branch named $2.

Once you are done, you MUST push the new branch origin, But you MUST NOT create any pull requests. 

${3..}
