---
title: Atomic hot backups with lm2
date: "2017-03-14T18:30:00-04:00"
---

Backups are an important topic for any storage system. Here's how I make hot,
atomic backups with lm2. It's both interesting and boring at the same time. It's an interesting
topic because atomic hot backups are a really useful feature to have, but it's boring because you
basically get it for free with lm2. :)

First, let me summarize what I mean by "hot" and "atomic."

An **atomic backup** is a backup copy that's essentially a snapshot of the data at a
single moment in time.  
A **hot backup** means you can create a backup copy while a system is still running.

lm2 provides cursors with snapshot views. When you make a cursor, you get to access data as if
the collection is frozen in time. There may be an active writer that's changing records, but
the cursor doesn't see any of that happening.

To create an atomic hot backup using that cursor, you simply have to read every key and write it
somewhere else (like S3) in whatever format you want (like JSON or CSV).

Here's the super simple code skeleton:

```go
var c *lm2.Collection

// Create a cursor, which provides a snapshot view of the collection
cur, err := c.NewCursor()
if err != nil {
	// Handle err
}
// Scan through each record
for cur.Next() {
	// Save the key-value pair somewhere
	save(cur.Key(), cur.Value())
}
```

The other nice thing about lm2 is that you can keep cursors open as long as you want.
lm2 never deletes data. At the moment, you're limited to creating cursors at the latest snapshot,
but it's actually possible to create cursors at *any point in the past*.

When would that be useful? Well, you could do neat things like diff two snapshots and determine
exactly what changed. Maybe you have two copies of an lm2 collection at different snapshots
(maybe one's restored from an older backup) and you need to sync the two without transferring an
entire backup's worth of data. It's just an idea for now.
