---
title: Removing duplicates in Bash history
date: "2016-12-27T21:15:00-05:00"
---

I have a habit of spamming commands when I get impatient. Like this:

```
 374  ssh -v preetam@10.10.13.139
 375  ssh -v preetam@10.10.13.139
 376  ssh -v preetam@10.10.13.139
 377  ssh -v preetam@10.10.13.139
 378  ssh -v preetam@10.10.13.139
...
 402  ssh -v preetam@10.10.13.139
 403  ssh -v preetam@10.10.13.139
 404  ssh -v preetam@10.10.13.139
 405  ssh -v preetam@10.10.13.139
```

By default, Bash on macOS doesn't remove duplicates, so I have to skip
through them when I use the `↑` key to go back through my history.

Fortunately, I can add one line to my `~/.profile` to remove those duplicates.

```bash
export HISTCONTROL=ignoredups
```
