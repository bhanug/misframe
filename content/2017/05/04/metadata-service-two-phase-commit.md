---
title: Metadata service & two-phase commit
date: "2017-05-04T23:30:00-04:00"
---

I have an app in the works (to be announced) and one of its components is a metadata service
that manages users, credentials, and basically any other data a web app would typically have.
I'm doing something different with this app. I'm not using a database server! All of my data is
being stored locally using lm2, my key-value store. The other neat thing it does is use a two-phase
commit protocol with another instance of itself to replicate data.

I could go into a lot of detail about how it works and how it's implemented, but I want to keep
this post short, so I'll just talk about its two-phase commit implementation. Before I do that,
I need to quickly describe the high level design.

The metadata service is an HTTP API that has three components:

1. A set of RESTful, read-only metadata endpoints
2. A single "do" write endpoint
3. A set of endpoints to provide a log abstraction

## Read and write endpoints

The read-only metadata endpoints are things like `GET /users/1`. Typical REST stuff.

The write endpoint, specifically `POST /do`, handles all modifications. It's not very RESTful, but
that's OK because it's designed to handle complicated, transactional operations through a single
endpoint. In my client library, methods like `service.CreateUser(User{})` are all implemented to
call the `/do` endpoint with a special payload. I'll talk about this payload in the next section.

## Log abstraction

This is the most interesting part. The metadata service is designed like a fully versioned database.
Each operation payload sent to the `/do` endpoint has its own version and is stored on a commit log.
Specifically, it's an lm2log, and which I already [wrote about recently](/2017/04/03/lm2log/).

The log abstraction exposes endpoints to prepare, commit, rollback, and retrieve log records.
The `/do` endpoint actually runs a lot of the code used by the log since it's basically a wrapper
around the log operations.

The log abstraction isn't used by clients of the metadata service. It's only used by *other instances
of the metadata service* as part of two-phase commit operations.

### Two-phase commit

The fact that I'm not using a database server means I don't have the convenience of using some
built-in replication mechanism. I still wanted my metadata service to be replicated, so I added some
two-phase commit logic into the `/do` endpoint.

I have two instances of the metadata service. There's one primary and one replica. Only the primary
handles writes, and it runs log operations on the replica.

Here are the two-phase commit stages that are run as part of a `/do` operation:

1. Prepare locally
2. Prepare on the replica
3. Commit locally
4. Commit on the replica
5. Acknowledge write

1 and 2 can be rolled back. From stages 3 and up, there's no turning back. If things can't move
forward for some reason, the service will just crash. On startup, the primary instance always
checks to see if it's in sync with the replica. If it's not in sync and *still* can't get in sync
after attempting to do so, it'll keep crashing. I'm clearly picking consistency over availability.

## Misc details

* Every write operation sent to the `/do` endpoint has to be deterministic since I only have
logical replication. For example, I can't randomly generate IDs during `/do` execution since
the replica could end up with different data.
* There's no automatic failover, so this isn't exactly "highly available."
* I could probably have another instance of the metadata service poll the log of the primary
to serve as an asynchronously replicated copy... :)

## Final thoughts

This post doesn't have a lot of detail, but it's longer than what I wanted already. I can write about
specific failure scenarios in future posts.

I'm thinking about open sourcing the core two-phase commit code as a package. I'd like to reuse this
stuff for other services anyway.
