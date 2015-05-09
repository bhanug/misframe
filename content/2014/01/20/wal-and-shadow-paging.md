---
title: Write-ahead logs and shadow paging
date: "2014-01-20"
url: /wal-and-shadow-paging
---


There are at least two ways to provide atomicity and durability in databases.
The first, and (I think) more common approach, is
[write-ahead logging](http://en.wikipedia.org/wiki/Write-ahead_logging). The
other is [shadow paging](http://en.wikipedia.org/wiki/Shadow_paging). The only
system that I'm aware of that uses shadow paging is [lmdb](http://symas.com/mdb/).

Write-ahead log (WAL)
---
This is pretty easy to understand. Every transaction is first appended to a log
as a set of mutations to the state. Eventually, these mutations will make their
way into the primary data structure (usually some kind of B-tree). In the event
of a crash, mutations can be rolled back by undoing mutations specified in the
log.

The cool thing about write-ahead logs is that writes happen sequentially. If
you're inserting lots of random data, you'll still be doing fast, sequential
writes. Eventually though, this data will be flushed into the primary data
structure and you can't avoid random writes. Still, this is great for short
bursts of writes.

Shadow paging
---
Shadow paging has a lot in common with persistent data structures. With shadow
paging, you *never* modify existing data. With the WAL approach, existing pages
will be modified when the log is flushed. In other words, updates are in-place.
With shadow paging, updates are append only. If you're using a tree, any
modification to the nodes will result in a new root, and essentially a new tree.
Atomicity in this case isn't difficult to grasp. If a transaction is committed,
the existing root node gets replaced with the new root. Otherwise, the new root
is discarded.

Another neat feature if atomic hot backups. This is a consequence of the
append-only, copy-on-write nature of the system.

Thoughts
---
So... what's better? It depends, of course!
Personally, I think it's *much* easier to understand shadow
paging. If you're using a memory-mapped tree, for example, things get a lot simpler.
You don't have to worry about undoing writes, keeping track of a file, making sure
data gets fetched from the log as needed, etc.

In terms of speed, I think it can go both ways. For short bursts of writes, I think
the WAL approach might be faster. Reading might be faster with shadow paging, since
you'll never have to dig through a log for recent changes.
As you start to look at the long-term performance, I think it could go either way.

