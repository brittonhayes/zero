# ZERO

> Find zero-days related to your code, deployments, and projects.

<div>
<img height="320px" src="./images/Zero.png" alt="zero"/>
</div>

## Installation

### Go

```sh
go get github.com/brittonhayes/zero
```

## Usage

```sh
❯ zero --help
Find zero-days related to your code, deployments, and projects.

Usage:
  zero [command]

Available Commands:
  help        Help about any command
  run         Fetch all matching zero day feed results
  serve       Start zero http server

Flags:
      --config string   config file (default is ./.zero.yaml)
  -d, --debug           enable debug mode
  -h, --help            help for zero

Use "zero [command] --help" for more information about a command.
```

## Modes

Zero comes with a few modes baked in. You can run it as command line tool, CI check, or view the results in a web
browser.

### CLI

> Running `zero` in the command line will look like this. A list of findings from
> your RSS feeds and patterns. This output can be customized with go text/templates
> if you have a preferred output format.

```shell
❯ zero run

TITLE:      ZDI-CAN-13000: Microsoft
SOURCE:     zeroday initiative - upcoming
LINK:       http://www.zerodayinitiative.com/advisories/upcoming/
PATTERN:    cvss\sscore\s[1-9]\.\d
FINDING:    "CVSS score 3.1"
---

TITLE:      ZDI-CAN-12977: Delta Industrial Automation
SOURCE:     zeroday initiative - upcoming
LINK:       http://www.zerodayinitiative.com/advisories/upcoming/
PATTERN:    cvss\sscore\s[1-9]\.\d
FINDING:    "CVSS score 7.8"
---
```

### Web

> Running `zero` in as an HTTP server wil look like this. It starts up a simple
> go http endpoint to serve your results in a clean dashboard.


<details>
<summary>Click to view preview</summary>
<pre>
❯ zero serve
INFO[0000] HTTP Server Started
INFO[0003] HOST="localhost:8091" METHOD=GET PATH=/
</pre>    
<img src="https://cdn.discordapp.com/attachments/145424435304726528/818764333319192626/unknown.png">
</details>

## Author

👤 **Britton Hayes**

* Website: https://brittonhayes.dev
* Github: [@brittonhayes](https://github.com/brittonhayes)

## Show your support

Give a ⭐️ if this project helped you!
