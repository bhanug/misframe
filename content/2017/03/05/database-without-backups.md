---
title: A database without backups?
date: "2017-03-05T21:40:00-05:00"
---

Here's an interesting thought experiment.

Suppose you had an immutable database, maybe something like Datomic or git. That means
all modifications to the database only add new data. Existing data will not change. That also
means there could be some sort of versioning associated with the database that you can reference
later. Also suppose that this database was replicated across data centers and/or geographic regions.

Do you need backups?

Well, what are backups used for? I can think of two things.

1. Recovery after data deletion or unavailability.
2. Recovery after data corruption.

I think both are addressed by an immutable, replicated database. If your primary region is lost,
you can still access your data from another replica. If you've corrupted your data somehow in the
most recent version, you can roll back your changes to some previous version.
Maybe your immutable database cleans up old versions after a while, and you'd want backups in case
that system had a bug and deleted active data... but you could also have bugs in your backup
retention system that causes the same issue.

I don't think a system like this needs backups. Neat!
