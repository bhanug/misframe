---
title: Time Series Databases Discussion Notes
date: "2015-07-20"
---

A couple of weeks ago at Gophercon, a few of us got together at a table to discuss time series databases. Specifically, we were interested in talking about time series storage. The group included Jason Moiron ([@jmoiron](https://twitter.com/jmoiron)), Paul Dix ([@pauldix](https://twitter.com/pauldix)), Ben Johnson ([@benbjohnson](https://twitter.com/benbjohnson)), Julius Volz ([@juliusvolz](https://twitter.com/juliusvolz)), and others representing companies and projects such as Datadog, InfluxDB, and Prometheus, just to name a few. I was representing VividCortex.

What follow are observations about each group’s time series storage problem and some approaches that they’re taking (or have taken) in order to solve it. Each group is distinct. Datadog is a monitoring company, Prometheus is a monitoring system, and InfluxDB is a time series database, and they all have different requirements in terms of reliability, compression, retention, and so on. Together, they cover a large portion of the time series storage problem space.

## Datadog

Datadog currently uses the popular distributed columnar database Cassandra for their time series. Based on what Jason Moiron has been working on at Datadog recently, it sounds like Cassandra isn’t meeting all of their needs. Their primary concerns about Cassandra seem to be regarding efficiency and reliability.

Many people recommend Cassandra for time series storage, but there are a few that don’t. Paul Dix, the CEO of InfluxDB, mentioned that he’s built two time series databases on top of Cassandra before and that it’s not made for this use case. One problem he mentioned was TTLs or key expiration. Apparently this isn’t very scalable in Cassandra, which makes implementing efficient retention logic quite difficult.

Jason is working on solving this problem using a unique approach. Instead of using an LSM or B-tree approach with some sort of indexing, as many seem to do, he’s working on a system that uses hash tables as buckets to hold time series points. I think he mentioned that it either has compression support now or will sometime in the future.

The interesting thing about Datadog is that they’re not really interested in a distributed time series database. Their system is distributed at a higher layer that involves Kafka, which is a distributed, replicated log. This is something what we do at VividCortex too. Data first get persisted to a Kafka cluster, which serves as a distributed write-ahead log, and eventually make their way into some sort of indexed backend storage system. If that backend storage system goes down and/or loses data, writes can still be accepted writes and replayed from the Kafka stream to recover.

Jason mentioned that their stack is designed such that time series data eventually go read only after a while. Their storage system apparently has multiple tiers involving Kafka, Redis, and some sort of “cold” storage. This means they don’t support backfilling of old data, but as a business whose main goal is to receive and store customers’ monitoring data, this isn’t a problem.

## Prometheus

Prometheus is a monitoring program and not a service like Datadog or VividCortex. It’s something that people install on their own systems and configure so that it polls other systems for metrics and stores them locally. This is quite different from a system that receives time series data from other sources.

As a result, Prometheus has a different set of assumptions that it makes for storing time series data. The most important, in my opinion, is that it does not accept out-of-order writes. There isn’t any way that you can have out-of-order writes because Prometheus polls for new values one-by-one. For this kind of system, it’s not a limitation. Unfortunately, this isn’t acceptable for systems that don’t rely on polling because it’s difficult to guarantee order without data loss.

Prometheus uses LevelDB to store metadata about time series, but they store the actual time series data in their own format using one file per series. This is great for them because all inserts are fast appends to the ends of files (because they’re always in order). Apparently they use XFS to be able to scale to the order of a million series, which means millions of files.

Within each file, they store points using delta encoding instead of using compression. Delta encoding is apparently very efficient for them as a replacement to compression because they can do seeks efficiently within ranges without having to decompress larger blocks.

Prometheus does not use a write-ahead log. They buffer up writes in memory and have an interval-based snapshotting system. In the worst case, they can lose up to 5 minutes of data. It’s interesting to note that their system doesn’t necessarily lose information in this case, but rather precision. For example, Prometheus stores raw counter values instead of computed derivatives, so rates can still be computed after losing data and they’ll just have lower resolution.


## InfluxDB

InfluxDB, being a distributed time series database, is trying to solve all of our time series storage problems. OK, maybe not quite, but they are trying to address a large set of problems. InfluxDB’s storage problem is unlike what the rest of us have to deal with. Being more of a general purpose time series database, it has unique requirements like the ability to store different types of values in time series points. For example, supporting arbitrary byte arrays as time series values is apparently on the roadmap. Furthermore, they’re building a distributed database which brings with it a whole set of challenging problems. For example, they can’t assume that their users will have Kafka in front of their database. They also need to support the ability to insert points at any time in the past to support backfilling. They have a diverse set of users, so they have a lot to consider.

Paul Dix mentioned that InfluxDB has performance problems right now. Their database currently uses Bolt, a transactional B+tree storage engine, for everything (including time series). Bolt is fast, but it’s not fast enough. In order to handle high bursts of writes, they’ve added a write-ahead log in front of Bolt. There was discussion about removing the WAL, but I can’t recall the reasoning.

Something else that the InfluxDB team is thinking about right now is compression. Their storage engine currently doesn’t support compression, but it’s something that any time series storage system desperately needs. They’re thinking about modifying their use of Bolt to support compression, but I have a feeling that it’ll have a big impact on performance and make their storage system a lot more complicated once they start to optimize it.


## Catena

I didn’t really talk about Catena that much, but I just want to summarize some of its features so you can try to see where it fits in. Catena is write-optimized in that it supports concurrent writers and has a write-ahead log. It keeps fixed time ranges, called partitions, in memory and eventually flushes them out to the disk. Each partition is stored as a single file with all of the time series stored in compressed extents. The partitions on disk are made read-only, so it’s not possible to backfill really old data. It supports out-of-order writes, but only to a limited extent. The main issue I see with Catena right now is that it’s not robust in terms of memory usage, but it’s not too bad considering there’s only about two or three weeks worth of work put into it. There is still plenty of potential for improvement.

---

Time series storage is a hard problem. It’s incredibly hard. It’s interesting (and maybe unfortunate) that we all have our own unique set of requirements that seem to prevent us from developing a single solution. In any case, this is such an exciting area that’s in a lot of development, and I feel extremely lucky to be involved and have the opportunity to share my opinion.

<br/>

Thanks to Abi ([@AbiGopal](https://twitter.com/AbiGopal)) for reading a draft of this.
