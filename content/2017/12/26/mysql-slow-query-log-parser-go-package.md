---
title: MySQL slow query log parser Go package
date: "2017-12-26T23:05:00-05:00"
---

tl;dr: Simple, MIT-licensed, available on GitHub: https://github.com/Preetam/mysqllog

---

I couldn't find a simple slow query log parser in Go so I decided to write one. The two I found
are:

* Honeycomb's [in Honeytail](https://github.com/honeycombio/honeytail/blob/master/parsers/mysql/mysql.go)
which is Apache licensed. But it does query text normalization which requires SQL parsing which I don't need.
* Percona's in their [go-mysql](https://github.com/percona/go-mysql/blob/master/log/slow/parser.go)
repo. It's AGPL licensed.

I don't like how those packages use channels. There's just too much plumbing required when you have
an input channel and an output channel, or a stop channel. With my package you just need to create
a parser and feed it slow query log data line-by-line. (I think channels are overused in some
Go programs but that's a separate discussion.)

### Example usage

I included a [small program](https://github.com/Preetam/mysqllog/blob/a5a229f69f224e733f20759937f5860c92af85a0/cmd/stdin-parser/main.go)
in the mysqllog package to read a MySQL slow query log from stdin and print the events as JSON objects to
stdout.

This is what it looks like right now:

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Preetam/mysqllog"
)

func main() {
	p := &mysqllog.Parser{}

	reader := bufio.NewReader(os.Stdin)
	for line, err := reader.ReadString('\n'); err == nil; line, err = reader.ReadString('\n') {
		event := p.ConsumeLine(line)
		if event != nil {
			b, _ := json.Marshal(event)
			fmt.Printf("%s\n", b)
		}
	}
}
```

To use that, I first spun up an RDS instance and used Honeycomb's [rdslogs](https://github.com/honeycombio/rdslogs) tool
to get the slow query log data from my instance to stdout. Then I just piped that to my program.
Works great, and only took a couple of hours or so!

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Current status. RDS logs. <a href="https://t.co/pFB2lEsdSZ">pic.twitter.com/pFB2lEsdSZ</a></p>&mdash; P R E E T A M (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/944752661101993984?ref_src=twsrc%5Etfw">December 24, 2017</a></blockquote>
<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>

I only tested it with MySQL 5.7 on RDS. It doesn't support older formats of the slow query log.
I will accept patches if you have them!
