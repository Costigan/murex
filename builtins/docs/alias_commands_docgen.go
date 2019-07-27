package docs

func init() {

	Definition["alias"] = "# _murex_ Language Guide\n\n## Command Reference: `alias`\n\n> Create an alias for a command\n\n### Description\n\n`alias` defines an alias for global usage\n\n### Usage\n\n    alias: alias=command parameter parameter\n    \n    !alias: command\n\n### Examples\n\nBecause aliases are parsed into an array of parameters, you cannot put the\nentire alias within quotes. For example:\n\n    # bad :(\n    » alias hw=\"out Hello, World!\"\n    » hw\n    exec: \"out\\\\ Hello,\\\\ World!\": executable file not found in $PATH\n    \n    # good :)\n    » alias hw=out \"Hello, World!\"\n    » hw\n    Hello, World!\n    \nNotice how only the command `out \"Hello, World!\"` is quoted in `alias` the\nsame way you would have done if you'd run that command \"naked\" in the command\nline? This is how `alias` expects it's parameters and where `alias` on _murex_\ndiffers from `alias` in POSIX shells.\n\nIn some ways this makes `alias` a little less flexible than it might\notherwise be. However the design of this is to keep `alias` focused on it's\ncore objective. For any more advanced requirements you can use a `function`\ninstead.\n\n### Detail\n\n#### Allowed characters\n\nAlias names can only include alpha-numeric characters, hyphen and underscore.\nThe following regex is used to validate the `alias`'s parameters:\n`^([-_a-zA-Z0-9]+)=(.*?)$`\n\n#### Undefining an alias\n\nLike all other definable states in _murex_, you can delete an alias with the\nbang prefix:\n\n    » alias hw=out \"Hello, World!\"\n    » hw\n    Hello, World!\n    \n    » !alias hw\n    » hw\n    exec: \"hw\": executable file not found in $PATH\n    \n#### Order of preference\n\nThere is an order of preference for which commands are looked up:\n1. Aliases - defined via `alias`. All aliases are global\n2. _murex_ functions - defined via `function`. All functions are global\n3. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function from the interactive command line\n4. variables (dollar prefixed) - declared via `set` or `let`\n5. auto-globbing prefix: `@g`\n6. murex builtins\n7. external executable files \n\n### Synonyms\n\n* `alias`\n* `!alias`\n\n\n### See Also\n\n* [`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [`function`](../commands/function.md):\n  Define a function block\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg *.txt)\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`private`](../commands/private.md):\n  Define a private function block\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [let](../commands/let.md):\n  \n* [source](../commands/source.md):\n  "

}
