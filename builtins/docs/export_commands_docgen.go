package docs

func init() {

	Definition["export"] = "# _murex_ Language Guide\n\n## Command Reference: `export`\n\n> Define a local variable and set it's value\n\n### Description\n\nDefines, updates or deallocates an environmental variable.\n\n### Usage\n\n    <stdin> -> export var_name\n    \n    export var_name=data\n\n### Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> export hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » export hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n### Detail\n\n#### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !export var_name\n    \nFor compatibility with other shells, `unset` is also supported but it's really\nnot an idiomatic method of deallocation since it's name is misleading and\nsuggests it is a deallocator for local _murex_ variables defined via `set`.\n\n#### Scoping\n\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n#### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out ($foo)\n    bar\n\n### Synonyms\n\n* `export`\n* `!export`\n* `unset`\n\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [equ](../commands/equ.md):\n  \n* [let](../commands/let.md):\n  "

}