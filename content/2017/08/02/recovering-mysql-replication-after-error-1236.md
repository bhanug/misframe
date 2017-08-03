---
title: Recovering MySQL replication after error 1236
date: "2017-08-02T23:40:00-04:00"
---

Error 1236 looks like this from `SHOW SLAVE STATUS`:

> Last_IO_Error: Got fatal error 1236 from master when reading data from binary log: 'Client requested master to start replication from position > file size'

In other words, the replica is requesting data at a certain point in the log (its current position),
but the master's log file doesn't reach that point (so there are missing entries). Replication stops
because because this is a logic error: if a replica is caught up to X, then the master *must* have been
at at least X, but it's not! One reason why this may happen is if MySQL hasn't flushed all of the data in the
binlog to disk.

When might MySQL do that? When `sync_binlog = 0`. Read more about that variable in the MySQL
[docs](https://dev.mysql.com/doc/refman/5.5/en/replication-options-binary-log.html#sysvar_sync_binlog).

### Diagnosis

Here's what you'll notice when you get error 1236.

First, take a look at the following fields from `SHOW SLAVE STATUS` on the replica.

```text
# From SHOW SLAVE STATUS
...
Master_Log_File: mysql-bin.001025
Read_Master_Log_Pos: 159997610 👈
...
```

Then take a look at the data directory (or binlog directory) on the master and look for the
binlogs.

```
# On the master
...
-rw-r-----  1 mysql mysql   152760218 Aug  3 00:31 mysql-bin.001025
                                👆
-rw-r-----  1 mysql mysql  1073787553 Aug  3 00:51 mysql-bin.001026
...
```

Notice how the replica wants to be at position **159997610** in the **mysql-bin.001025** binlog file,
but the file is only **152760218** bytes long on the master.

Also notice that there's an additional binlog file in the sequence: mysql-bin.001026.

### Solution

In order to get replication started again, you need to point the replica to read from the
beginning of the new binlog file.

To do that, run `CHANGE MASTER TO` with the new binlog file name and a starting offset of 4.

```
# On the replica
CHANGE MASTER TO MASTER_LOG_FILE='mysql-bin.001026', MASTER_LOG_POS=4;
```

Note that unchanged values stay the same.

### Answers to potential questions

##### Why is it safe to use the next binlog file?

You probably ran into this error after the master crashed. MySQL creates a new binlog file
every time it starts, so the new log file is the next valid starting point.

From the MySQL [docs](https://dev.mysql.com/doc/refman/5.7/en/binary-log.html):

> mysqld appends a numeric extension to the binary log base name to generate binary log file names. The number increases each time the server creates a new log file, thus creating an ordered series of files. The server creates a new file in the series each time it starts or flushes the logs. The server also creates a new binary log file automatically after the current log's size reaches max_binlog_size. A binary log file may become larger than max_binlog_size if you are using large transactions because a transaction is written to the file in one piece, never split between files.

##### Will I lose any data?

(Assuming `sync_binlog = 0`.)

Maybe. If your replica was all caught up before the master crashed, then you probably didn't lose much data (if at all).
If your replica was not caught up and didn't manage to pull the missing binlog records... yes, you
probably lost data. But that's one of the risks of using `sync_binlog = 0`.
