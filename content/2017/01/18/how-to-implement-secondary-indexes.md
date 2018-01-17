---
title: How to implement secondary indexes
date: "2017-01-18T00:40:00-05:00"
bestof: true
---

This post is about implementing secondary indexes on top of an ordered key-value store.
This topic is interesting for at least two reasons. First, you may actually need to do
this if you're implementing secondary indexes on top of something like LevelDB, RocksDB, Bolt,
or some other key-value storage library. Second, seeing how this is done from an
implementation perspective can help you understand how databases like MySQL and PostgreSQL
handle secondary indexes.

Suppose we have a table implemented on top of a key-value store. Specifically, assume
that the key-value store has unique, ordered keys.

Here's the row definition:

```
Row := (id, username, email, deleted timestamp, name)
```

There are five columns. We'll let `id` be the primary key. There's a
`deleted` column to allow usernames to be reused. More on that later.

Here's what the table would look like when it's mapped to keys and values:

**Table (Primary key is ID)**

| Key | Value |
|-----|---------------------------|
| id=1 | username=alice, email=alice@example.com, deleted=0, name=Alice |
| id=2 | username=bob, email=bob@example.com, deleted=0, name=Bob |
| id=3 | username=bob2, email=bob2@example.com, deleted=0, name=Bob |

The primary key is ID, and it's part of the "key" part of the key-value
store, so it's unique and ordered.

I mentioned that usernames are allowed to be reused. The `deleted` column
allows us to logically mark rows as deleted without actually deleting
data. One thing we need to do is make sure active usernames are unique.
We can do that by adding a *unique index*.

### Unique index

A unique index can be implemented on top of the same kind of ordered
key-value store with unique keys. This time, instead of the ID being the
key, we'll use the (username, deleted) tuple instead. As for the value,
we'll use that to point the secondary index entry back to the original
row. We can do that with a row identifier, i.e. some value that uniquely
identifies a row. Sounds like a primary key, doesn't it? That'll do the
trick.

Here's what the unique secondary index looks like:

**Unique index on (username, deleted)**

| Key | Value |
|----------|-------|
| username=alice, deleted=0 | id=1 |
| username=bob, deleted=0 | id=2 |
| username=bob2, deleted=0 | id=3 |

Each key in this index corresponds to exactly one row in the table
because the ID is a unique identifier.

### Non-unique index

What if we wanted to add a non-unique index? We're still going to use the
same ordered key-value store, but we need to modify it slightly since
it requires that keys are unique. We need to take our non-unique keys and
make them unique. Fortunately, it's quite easy. All you need to do is
introduce a unique element to a key to make the whole thing unique.
We already have something like that: the row identifier. This time, instead
of keeping that as part of the value, we'll move it into the key.

Here's what the non-unique index looks like:

**Index on (name)**

| Key | Value |
|----------|-------|
| name=Alice, id=1 |  |
| name=Bob, id=2 |  |
| name=Bob, id=3 |  |

The trick is to simply *ignore* the row identifier part of the key whenever
we access an element. Just like before, we know where to look to find the
rest of the row if we need unindexed columns.

### How MySQL implements secondary indexes

What I've described so far is essentially how secondary indexes work in MySQL<sup>1</sup>. Here
is the relevant documentation about it:

> All indexes other than the clustered index are known as secondary indexes. In InnoDB, each record in a secondary index contains the primary key columns for the row, as well as the columns specified for the secondary index. InnoDB uses this primary key value to search for the row in the clustered index.

> -- https://dev.mysql.com/doc/refman/5.7/en/innodb-index-types.html

I think it's useful to think about this when you're designing your indexes. No need to include
primary key columns in your secondary indexes if you don't need to!

### How PostgreSQL implements secondary indexes

Just for comparison, I think it's also interesting to think about how PostgreSQL implements
secondary indexes. Again, it's essentially the same, but instead of using the primary key
columns to uniquely identify rows, Postgres uses *tuple identifiers (TIDs)*. Not only do
TIDs uniquely identify a row (and therefore a primary key), but they identify a row at *a
particular version*. Unlike with MySQL, where primary key columns are implicitly part of
the secondary index entries, getting the primary key columns in Postgres requires a separate
lookup. Also, updating a row always updates secondary indexes even if the primary key stays
the same, because of TID changes.

Here's the relevant documentation about that:

> An index is effectively a mapping from some data key values to tuple identifiers, or TIDs, of row versions (tuples) in the index's parent table. A TID consists of a block number and an item number within that block (see Section 65.6). This is sufficient information to fetch a particular row version from the table. Indexes are not directly aware that under MVCC, there might be multiple extant versions of the same logical row; to an index, each tuple is an independent object that needs its own index entry. Thus, an update of a row always creates all-new index entries for the row, even if the key values did not change. (HOT tuples are an exception to this statement; but indexes do not deal with those, either.) Index entries for dead tuples are reclaimed (by vacuuming) when the dead tuples themselves are reclaimed.

> -- https://www.postgresql.org/docs/current/static/indexam.html

If you want to learn more about MySQL vs PostgreSQL indexes, check out
["Why Uber Engineering Switched from Postgres to MySQL"](https://eng.uber.com/mysql-migration/).

---

I hope this was useful. This is how I think about secondary indexes, and I think it's much more
valuable to think about things like this in a more abstract way. Only looking at specific
implementations won't always help you build up your intuition, and (IMO) intuition is extremely
valuable when you're working with complex systems.

**Footnotes**

1. If I'm wrong, please leave a comment!
