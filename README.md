goverview
=========
goverview - Get overview about list of URLs

## Installation

```
go get -u github.com/j3ssie/goverview
```

## Example Commands

```shell
goverview - Overview about list of URLs - beta v0.2.2 by @j3ssiejjj

Usage:
  goverview [command]

Available Commands:
  help        Help about any command
  probe       Do Probing on target
  screen      Do Screenshot on target

Flags:
  -B, --burp                Receive input as base64 burp request
  -c, --concurrency int     Set the concurrency level (default 10)
  -C, --content string      Summary File for Content (default 'out/content-summary.txt')
      --debug               Debug output
  -H, --headers strings     Custom headers (e.g: -H 'Referer: {{.BaseURL}}') (Multiple -H flags are accepted)
  -h, --help                help for goverview
  -I, --inputFile string    Custom headers (e.g: -H 'Referer: {{.BaseURL}}') (Multiple -H flags are accepted)
  -i, --inputs strings      Custom headers (e.g: -H 'Referer: {{.BaseURL}}') (Multiple -H flags are accepted)
  -j, --json                Output as JSON
  -l, --level int           Set level to calculate CheckSum
  -N, --no-output           No output
  -o, --output string       Output Directory (default "out")
  -L, --redirect            Allow redirect
      --retry int           Number of retry
  -R, --save-redirect       Save redirect URL to overview file too
  -S, --screenshot string   Summary File for Screenshot (default 'out/screenshot-summary.txt')
      --sortTag             Sort HTML tag before do checksum
  -t, --threads int         Set the threads level for do screenshot (default 5)
      --timeout int         HTTP timeout (default 15)
  -v, --verbose             Verbose output
  -V, --version             Print version
  -W, --wordlist string     Wordlists File build from HTTP Content (default 'out/wordlists.txt')

Use "goverview [command] --help" for more information about a command.


Checksum Content Level:
  0 - Only check for src in <script> tag
  1 - Check for all structure of HTML tag + src in <script> tag
  2 - Check for all structure of HTML tag + src in <script> <img> <a> tag
  5 - Entire HTTP response

Examples:
  # Only get summary
  cat http_lists.txt | goverview probe -N -c 50 | tee only-overview.txt

  # Get summary content and store raw response without screenshot
  cat http_lists.txt | goverview probe -c 20 -M --json

  # Only do screenshot
  cat list_of_urls.txt | goverview --skip-probe

  # Do screnshot
  cat http_lists.txt | goverview screen -c 5 --json

  # Do screnshot based on success HTTP site
  cat overview/target.com-http-overview.txt | jq -r '. | select(.status=="200") | .url' | goverview screen -c 5 -o overview -S overview/target.com-screen.txt
```

## License

`goverview` is made with ♥  by [@j3ssiejjj](https://twitter.com/j3ssiejjj) and it is released under the MIT license.
