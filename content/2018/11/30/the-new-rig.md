---
title: The New Rig
date: "2018-11-30T20:05:00-08:00"
---

When I first wrote the Rig, I wanted a framework that I could use to keep
the data of my stateful services safe, and that safety came from replication
and two-phase commit. I blogged about its implementation
[last year](https://misfra.me/2017/07/19/the-rig/). It was _OK_. I've used it
for Transverse since the beginning, and while nothing went terribly wrong and
it seemed to work, it was too complicated and difficult to test. I didn't trust
it completely to keep my data safe, so I had a pending task to add backups
to Transverse.

There were several problems with the old Rig's design. I'll just list them so
I can focus on the new version in this post.

* There had to be another server running, which was wasteful for a service like mine with almost no traffic.
* Failover was manual and untested.
* The Rig needed its own endpoint so it was weird to hook it into an existing application.

The new Rig is almost completely different. It doesn't use synchronous replication or two-phase commit,
and instead uses S3 as a durable log and snapshot store. It also assumes that there's only one instance
of your service running too. The goal is the same though: protecting application data during instance failure.

* **Before:** replication
  * Dependencies: Another instance running the Rig
  * Failover: Manual
  * Backup and restore: Not supported
* **Now:** log
  * Dependencies: S3
  * Failover: Automatic
  * Backup and restore: Same as failover

The only thing that's mostly the same is the `Service` interface that your
application code needs to satisfy. There are two additions to the interface: `Snapshot` and `Restore`.
These methods make snapshot backups and restores first-class.

```go
type Service interface {
    Version() (uint64, error)
    Validate(Operation) error
    Apply(uint64, Operation) error
    Snapshot() (io.ReadSeeker, int64, error)
    Restore(uint64, io.Reader) error
}
```

This made it easy to rewrite all of the Rig but keep most of Transverse the same. I just had to
implement the two new methods.

## S3

The new Rig stores log and snapshot data in S3. Here is what the top level layout looks like:

```
LATEST
LOG/
SNAPSHOT/
```

The `LATEST` object contains the version of the latest snapshot, e.g.

```
48e
```

The `LOG/` prefix contains log record batches. Each object has one or more log records. Each
log record is some sort of mutation operation.

```
LOG/0000000000000462
LOG/0000000000000463
LOG/0000000000000468
...
```

The `SNAPSHOT/` prefix contains objects with full snapshots of the application data. Note that
it's up to the application to interpret what a snapshot means. You don't have to actually store
all of your application data in a single S3 object. I do it with Transverse since it's tiny, but
I can imagine other applications just storing references to data located elsewhere.

```
SNAPSHOT/0000000000000430
SNAPSHOT/000000000000044d
SNAPSHOT/0000000000000462
```

### Consistency

S3's consistency model is something to keep in mind.

> Amazon S3 provides read-after-write consistency for PUTS of new objects in your S3 bucket in all regions with one caveat. The caveat is that if you make a HEAD or GET request to the key name (to find if the object exists) before creating the object, Amazon S3 provides eventual consistency for read-after-write.
> https://docs.aws.amazon.com/AmazonS3/latest/dev/Introduction.html

* Must not try to access the next log object unless you know it exists
* 
