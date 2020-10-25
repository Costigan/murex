package docs

func init() {

	Definition["murex-parser"] = "# _murex_ Shell Docs\n\n## Command Reference: `murex-parser` \n\n> Runs the _murex_ parser against a block of code \n\n## Description\n\n`summary` define help text for a command. This is effectively like a tooltip\nmessage that appears, by default, in blue in the interactive shell.\n\nNormally this text is populated from the `man` pages or `murex-docs`, however\nif neither exist or if you wish to override their text, then you can use\n`summary` to define that text.\n\n## Usage\n\n    <stdin> -> murex-parser -> <stdout>\n    \n    murex-parser { code-block } -> <stdout>\n\n## Detail\n\nPlease note this command is still very much in beta and is likely to change in incompatible ways in the future. If you do happen to like this command and/or have any suggestions on how to improve it, then please leave your feedback on the GitHub repository, https://github.com/lmorg/murex\n\n## See Also\n\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [commands/`murex-docs`](../commands/murex-docs.md):\n  Displays the man pages for _murex_ builtins\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_"

}
