package docs

func init() {

	Definition["get"] = "# _murex_ Language Guide\n\n## Command Reference: `get`\n\n> Makes a standard HTTP request and returns the result as a JSON object\n\n### Description\n\nFetches a page from a URL via HTTP/S GET request\n\n### Usage\n\n    get url -> <stdout>\n    \n    <stdin> -> get url -> <stdout>\n\n### Examples\n\n    » get google.com -> [ Status ]\n    {\n        \"Code\": 200,\n        \"Message\": \"OK\"\n    }\n\n### Detail\n\n#### JSON return\n\n`get` returns a JSON object with the following fields:\n\n    {\n        \"Status\": {\n            \"Code\": integer,\n            \"Message\": string,\n        },\n        \"Headers\": {\n            string [\n                string...\n            ]\n        },\n        \"Body\": string\n    }\n    \n    The concept behind this is it provides and easier path for scripting eg pulling\n    specific fields via the index, `[`, function.\n    \n    #### `get` as a method\n    \n    Running `get` as a method will transmit the contents of STDIN as part of the\n    body of the HTTP GET request. When run as a method you have to include a second\n    parameter specifying the Content-Type MIME.\n    \n    #### Configurable options\n    \n    `get` has a number of behavioral options which can be configured via _murex_'s\n    standard `config` tool:\n    \n    config: -> [ http ]\n    \n    To change a default, for example the user agent string:\n    \n    config: set http user-agent \"bob\"\nget: google.com\n    \n    This enables sane, repeatable and readable defaults. Read the documents on\n    `config` for more details about it's usage and the rational behind the command.\n\n### See Also\n\n* [`getfile`](../docs/commands/getfile.md):\n  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.\n* [`post`](../docs/commands/post.md):\n  HTTP POST request with a JSON-parsable return\n* [config](../docs/commands/commands/config.md):\n  \n* [square-brace-open](../docs/commands/commands/square-brace-open.md):\n  "

}