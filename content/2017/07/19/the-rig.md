---
title: The Rig
date: "2017-07-19T21:30:00-04:00"
---

Back in May I [wrote](/2017/05/04/metadata-service-two-phase-commit/) about a web service I'm working
on that uses local storage and replication with two-phase commit. I pulled out the core of it
and created a package that I'm calling the Rig.

It's up on GitHub already: https://github.com/Preetam/rig

![The Rig](/img/2017/07/the-rig.png)

The goal of the Rig is to take some web service and add a log and replication on top.

A service is simply something that accepts operations.

```go
type Service interface {
    Validate(Operation) error
    Apply(uint64, Operation) error
    LockResources(Operation) bool
    UnlockResources(Operation)
}
```

Each operation is associated with an entry in a log.

```go
type LogPayload struct {
    Version uint64    `json:"version"`
    Op      Operation `json:"op"`
}
```

And an operation is a method with some data.

```go
type Operation struct {
    Method string          `json:"method"`
    Data   json.RawMessage `json:"data"`
}
```

The service I created has methods like `user_create`, `user_update`, `user_delete`, and so on.

---

The Rig is still *very* rough. It kinda works... meaning it works great during "happy" times.
There's a lot of work to do to handle errors. Synchronous replication is one of those things that
makes a lot of sense in the abstract, but can be all over the place in implementation.

For example, consider synchronous replication with two nodes. Writes succeed as long as both
the primary and the replica acknowledge the write. If the primary fails, usually the replica takes
over without any loss of data. But then you're left with only one node running, and no more replication.
What if instead of the primary failing, only the replica fails? Should the primary keep going? If it does,
didn't we just ignore the synchronous part of the replication?

But that's the expected behavior for most systems. You want the system to be available even if a
node fails. And when a node fails, you don't want to lose any data. It's kind of complicated to
decide whether to just wait for a failing node or consider it failed and move on.

I think it's interesting to see how MySQL/MariaDB does it with semi-sync replication. The master
will wait up to some configurable number of milliseconds for a replica to respond during semi-sync
mode, and when the timeout is exceeded the master will leave the replica behind and continue in
async mode. That way you're not stuck with a failed replica, but are synchronous during "happy" times.

The Rig just crashes the program right now when it can't make progress with the replica =P.
I'm working on it.

---

Other stuff I was looking at:

* https://dev.mysql.com/doc/refman/5.7/en/replication-semisync.html
* DRBD internals: https://docs.linbit.com/docs/users-guide-9.0/p-learn/#ch-internals
  * Some two-phase commit stuff in DRBD: https://git.drbd.org/drbd-9.0.git/blob/HEAD:/drbd/drbd_state.c
* FreeBSD HAST: https://wiki.freebsd.org/HAST
