---
title: Databases as a log
date: "2016-10-15T12:40:00-04:00"
---

I tweeted this a few days ago and I've been thinking more about it.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Databases can just be a log. Everything else is a read optimization.</p>&mdash; Thought retweeter (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/786566068026302464">October 13, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

A database is a set of information. You can add, modify, or remove information from that set.
All of these types of changes can be expressed as instructions. What happens when you put these
instructions together? You get a log.

Starting with an append-only, immutable log as a database makes you think about all kinds of stuff
differently. When I started working on lm2, I had a fancy (but still relatively simple) [design]
(https://github.com/Preetam/lm2/blob/d96eaa16f44cd015762d9db8c58fec9152c3b41b/DESIGN.md#design) with
block allocation, MVCC, free lists and garbage collection, and the possibility of adding
compression.

It was complicated. It was hard to think about. It looked like a mountain of work and I didn't want
to keep hacking away on something so complicated. At one point I just threw out everything, got rid
of my notes, and started with something super simple. Now it's just a log with a small linked list
read optimization and some caching.

The fact that it's so minimal means it doesn't restrict much you in terms of what you can do.
Records are never deleted. Ever. Of course, that means you'll use more disk space, but that can be
solved at a higher level. The upside is that you can have snapshot reads *forever*. That's a pretty
cool feature, in my opinion.

I've been thinking lately that it's probably best to push read optimizations up the stack as much
as possible. That's the approach [RDS Aurora](https://aws.amazon.com/rds/aurora/) took with their
storage engine, and they can do some amazing things. I'm pushing for this kind of approach at work
too. Our database is a log. MySQL is a read optimization.
