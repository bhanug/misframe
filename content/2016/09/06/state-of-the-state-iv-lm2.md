---
title: "State of the State Part IV: lm2, A Linked List Storage Library"
date: "2016-09-06T22:57:59-04:00"
---

State of the State is my series covering my storage and time-series experiments
and adventures. Here are parts I-III, in case you're curious to see how I ended
up here!

- [Part I](https://www.misfra.me/state-of-the-state/)
- [Part II](https://www.misfra.me/state-of-the-state-part-ii/)
- [Part III](https://www.misfra.me/state-of-the-state-part-iii/)

---

During part II, I briefly mentioned my [metricstore](https://github.com/Preetam/metricstore) storage
library, which handles time-series metrics storage. metricstore is mainly just the wrapper; the
actual data storage and retrieval happens in a separate package called [listmap]
(https://github.com/Preetam/listmap).

listmap has a *very* simple design. It's a key-value store where records are appended to the end
of a memory-mapped file. Each record has a header that stores information like the lengths of the
key and value. There are also two offsets pointing to the adjacent records, in key order. There is
an additional field to indicate whether the record has been removed or not.

In code, it looks like this:

```
struct {
	prev    uint64
	next    uint64
	keylen  uint16
	vallen  uint16
	removed bool
}
```
Again, quite simple. It's almost like how you would write a doubly-linked list in memory, except
I was doing it against a memory-mapped file, with memory operations. Bad!

Eventually I moved on to write [catena](https://github.com/Cistern/catena) (described in Part III),
which became a "log-structured" and much safer time-series storage engine.

## Introducing lm2

[lm2](https://github.com/Preetam/lm2), short for listmap2, is the second version of my listmap idea.
I always liked the simplicity of having an append-only linked list on disk. This time, it's much
safer and has a more interesting design.

The highlights:

* Ordered key-value data model
* Append-only modifications
* Fully durable, atomic writes
* Cursors with snapshot reads

lm2, just like listmap, is an ordered key-value storage library. Each instance of a map is called
a *collection*. You set and delete key-value pairs, and you are given a cursor to iterate through
records in a collection.

New records are appended to the end of the data file. Deletes happen logically, so you'll never
actually remove records. To reclaim space, you'll have to rewrite an entire collection.

### Performance

I'd rather not spend too much time talking about specific performance characteristics since this
project is still super young, performance numbers are hard to interpret, and it's not really the
main goal of this project.

That said, I *am* interested in making this linked list go as fast as possible!

First, unlike listmap, there is no memory mapping. Records are loaded into memory and flushed out
to disk during metadata updates. This allows for more control over which records are already in
memory, so there is a record cache which contains a small subset of records.

The record cache information is regularly serialized to disk, so you could experience a crash or
shutdown and recover cached records without having to wait for things to warm up naturally. This
makes a *huge* difference on recovery, since this is where you really see the low O(n) searches make
a difference.

Because of the append-only nature, it was easy to support concurrent readers that don't block
writes. There is, however, only a single writer.

### Durability

lm2 is durable. Durability was my top priority, and something I wanted to get done before moving on
to things like performance. While new records are appended to the data file, lm2 still needs to do
in-place updates to modify offset pointers. This is not copy-on-write. Therefore, lm2 depends on a
write-ahead log (WAL). Writes also happen in user-defined batches.

Note that new key-value record data *does not* get written to the WAL. I didn't want to append
writes to both the WAL and the data file, so they only get appended to the data file. This costs an
additional fsync, which is an OK trade-off.

Writes happen in 3 stages:

1. Append new data to the data file
2. Append data file updates to the WAL
3. Apply in-place data file updates, including data file header updates

Each stage can crash at any moment and, depending on where in the sequence the crash happened, the
write will either entirely happen or it won't at all. Writes are atomic.

At the moment, the only thing you should do if a write fails is bail out and re-open a collection.
There's no logic to recover from a failed write. This is where an *undo log* would come in handy!

## ACID

I believe it's possible to implement ACID transactions on top of anything that provides the
following two properties:

1. Atomic & durable writes
2. Snapshot reads

lm2 provides both. So, can you implement serializable, ACID transactions on top of lm2? Yeah!

In fact, one really easy way is to just use a mutex. It's that simple! I'll leave it as an exercise
for you to figure out how to do it.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Sorry, I&#39;m not impressed with serializable isolation via a single writer mutex.</p>&mdash; Preetam ᕕ( ᐛ )ᕗ (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/537313622410952704">November 25, 2014</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Of course, it's not as impressive as multiversion concurrency control (MVCC), optimistic
transactions, and multiple writers, but it's a *lot* simpler to implement.

---

That's it for this part of *State of the State*. There's a lot more on the way! If you want more
frequent updates, storage engine shower thoughts, and some other rambling on a daily basis, [follow
me on Twitter](https://twitter.com/PreetamJinka).

You can find [lm2 on GitHub](https://github.com/Preetam/lm2).
