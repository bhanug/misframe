---
title: Faster MySQL replication with group commit and delay
date: "2017-01-04T22:51:00-05:00"
bestof: true
---

We've been having a problem with MySQL replication at [VividCortex](https://www.vividcortex.com). Replicas periodically
tend to fall behind and we couldn't really figure out how to speed things up.
It wasn't about resources. The replicas have plenty of CPU and I/O available. We're
also using multithreaded replication (a.k.a. MTR) but most of the replication threads
were idle.

One thing that we decided to try out was the new `LOGICAL_CLOCK` parallelization policy
introduced in MySQL 5.7.2. Here's what the [MySQL reference manual](http://dev.mysql.com/doc/refman/5.7/en/replication-options-slave.html#option_mysqld_slave-parallel-type)
says about `slave-parallel-type`:

> When using a multi-threaded slave (`slave_parallel_workers` is greater than 0), this option specifies the policy used to decide which transactions are allowed to execute in parallel on the slave. The possible values are:

> * `DATABASE`: Transactions that update different databases are applied in parallel. This value is only appropriate if data is partitioned into multiple databases which are being updated independently and concurrently on the master. Only recommended if there are no cross-database constraints, as such constraints may be violated on the slave.

> * `LOGICAL_CLOCK`: Transactions that are part of the same binary log group commit on a master are applied in parallel on a slave. There are no cross-database constraints, and data does not need to be partitioned into multiple databases.

We've been using `--slave-parallel-type=DATABASE`, but clearly it has not been offering
the level of parallelism we want. So we tried `LOGICAL_CLOCK`.

Initially, this ended being *slower* than `--slave-parallel-type=DATABASE`. My guess was
we're not grouping enough transactions per binary log commit for this to be a big
improvement. Fortunately, that's something we can tune using `binlog_group_commit_sync_delay`.

[Here](http://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_delay) is the documentation about that:

> Controls how many microseconds the binary log commit waits before synchronizing the binary log file to disk. By default `binlog-group-commit-sync-delay` is set to 0, meaning that there is no delay. Setting `binlog-group-commit-sync-delay` to a microsecond delay enables more transactions to be synchronized together to disk at once, reducing the overall time to commit a group of transactions because the larger groups require fewer time units per group. With the correct tuning, this can increase slave performance without compromising the master's throughput.

Setting `binlog-group-commit-sync-delay` to 3000 (3 ms) and `--slave-parallel-type=LOGICAL_CLOCK`
resulted in a huge improvement in replication delay. Too bad I didn't learn about this sooner!

**UPDATE:** We wrote a follow-up blog post on our company site with some more details (and pictures!):
["Solving MySQL Replication Lag with LOGICAL_CLOCK and Calibrated Delay"](https://www.vividcortex.com/blog/solving-mysql-replication-lag-with-logical_clock-and-calibrated-delay)
