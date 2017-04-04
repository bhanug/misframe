---
title: lm2log
date: "2017-04-03T23:45:00-04:00"
---

I started [lm2log](https://github.com/Preetam/lm2log) 6 months ago according to GitHub.
I think it was the first useful abstraction I wrote on top of [lm2](https://github.com/Preetam/lm2),
my linked list storage library written in Go.

The API is really simple. You basically just `Prepare`, `Commit`, `Rollback`, and `Get` data.
Because it's using lm2, all of lm2log's operations are fully durable and crash-safe. I don't have
anything else to say about this package besides the fact that it's essential to the two-phase commit
logic of the metadata service I'm building (to be announced!).

The fun thing about this is that it's totally unimpressive. lm2 is unimpressive. It's a slow
storage library implemented as a *linked list*. lm2log is a commit log built on top of lm2. And I'm
building more things on top. Compared to "modern" production database engines, this stuff looks more
like a silly toy. *But*, I'm still getting some great guarantees. I already explained how I can get
[ACID transactions with lm2](/2016/10/04/lm2-transactions/). My goal lately has been to take simple,
unimpressive things and build robust, impressive systems on top. So far so good.
