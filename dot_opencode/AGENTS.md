You are working on configuration for the OpenCode application. 

If you need to examine the source code to OpenCode to help you determine how to perform your edits, you may find its source code through the symbolic link at ~/oc/opencode-source`. Remember, your job is to edit the configuration, not the source code. You **MUST NOT** ever modify the source code to open code itself. 

Two configuration files are in use:

- The primary configuration is in the opencode.json file. 
- A second configuration file that is not tracked in Git is located at opencode.jsonc.

The second configuration file is used to store preferences local to this and is generally NOT the one that you should be modifying. It currently only contains a uniquely selected theme to distinguish this machine from from other machines. 

**NOTE**: There are FOIR separate git repositories involved here, and none are submodules of the others!

- opencode-config (at ~/.local/config/opencode locally, https://github.com/ariane-emory/opencode-config/ on Github)
- opencode-config-supplemental (at ~/.local/config/opencode/supplemental locally, https://github.com/ariane-emory/opencode-config-supplemental/ on Github)
- MUST-have-plugin (at ~/.local/config/opencode/supplemental/plugins/MUST-have-plugins locally, symlinked at ~/.local/config/opencode/plugin/MUST-have-plugins locally, https://github.com/ariane-emory/MUST-have-plugin on Github)
- todo-reminder plugin (at ~/.local/config/opencode/supplemental/plugins/todo-reminder locally, symlinked at ~/.local/config/opencode/plugin/todo-reminder locally, https://github.com/ariane-emory/todo-reminder on Github)
