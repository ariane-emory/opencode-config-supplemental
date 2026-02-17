Set the title of the current session by prefixing the name of the branch at its beginning.

So, for example, if the session is currently named "foo bar baz" and the current branch is named "feat/quux", its new title **SHOULD** be "feat/quux foo bar baz".

In addition, bookmark the current session.

**IMPORTANT**: Use the `get_current_session_title`, `set_current_session_title` and `bookmark_current_session` tools to do this. These are **agent tools** that you can invoke directly - do NOT search for their implementation in the working codebase, just call them. They are available to you regardless of which branch you are working in.
