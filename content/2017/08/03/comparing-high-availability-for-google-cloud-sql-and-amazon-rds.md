---
title: Comparing high availability for Google Cloud SQL and Amazon RDS
date: "2017-08-03T23:30:00-04:00"
---

[Cloud SQL](https://cloud.google.com/sql/) is Google Cloud's DBaaS platform that supports MySQL
and now PostgreSQL in beta. [RDS](https://aws.amazon.com/rds/) is AWS's DBaaS.
They both have "high availability" options which provide multi-zone redundancy and an automated failover
mechanism. While they sound exactly the same at a high level, they have some important differences!
This post describes some of the bigger differences I've noticed. I'm also focusing on MySQL.

## Implementation

#### Cloud SQL

[Reference](https://cloud.google.com/sql/docs/mysql/high-availability)

Cloud SQL is implemented using **semisynchronous** replication. This isn't surprising since Google contributed
the semisynchronous replication support for MySQL.

You can read more about semisynchronous replication in MySQL in the [reference manual](https://dev.mysql.com/doc/refman/5.7/en/replication-semisync.html).

#### RDS

[Reference](https://aws.amazon.com/rds/details/multi-az/)

The RDS information pages on the AWS web site don't actually specify the implementation besides the
fact that it's **synchronous**. The fact that it's implemented using [DRBD](https://en.wikipedia.org/wiki/Distributed_Replicated_Block_Device) has been mentioned in some
talks and slide decks, like [this](https://www.slideshare.net/AmazonWebServices/aws-reinvent-2016-deep-dive-on-amazon-aurora-dat303/8?src=clipshare)
one from the re:Invent 2016 Deep Dive on Amazon Aurora talk (available on [YouTube](https://www.youtube.com/watch?v=duf5uUsW3TM)).

Because replication happens at the block device level (below MySQL), the standby instance doesn't actually
have MySQL running!

## Read scaling

#### Cloud SQL

[Reference](https://cloud.google.com/sql/docs/mysql/high-availability#about_using_the_failover_replica_as_a_read_replica)

You **can** use the failover replica on Cloud SQL to serve reads.

#### RDS

[Reference](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html)

> The high-availability feature is not a scaling solution for read-only scenarios; you cannot use a standby replica to serve read traffic.

You **cannot** serve read traffic from the standby replica.


## Failover considerations

#### Cloud SQL

[Reference](https://cloud.google.com/sql/docs/mysql/high-availability#how_failover_affects_your_applications_and_your_instances)

> After the failover, the replica becomes the master, and Cloud SQL automatically creates a new failover replica in another zone. If you located your Cloud SQL instance to be near other resources, such as a Compute Engine instance, you can relocate your Cloud SQL instance back to its original zone when the zone becomes available. Otherwise, there is no need to relocate your instance after a failover.

During failover, all existing connections will be dropped. However, after failover, your applications can use the same IP address to connect to the
new primary instance.

[This](https://groups.google.com/d/msg/google-cloud-sql-discuss/WwfY_CwFbVU/IKfo7Rn_BwAJ) Google Groups post by Jay on the Cloud SQL team
also has some great answers.

> 4) **My failover replica has an IP.. do I need to change my clients to use this IP, or will the old primary IP now start pointing at the failover replica? i.e. is this really a floating IP that gets moved?**

> No. There should be zero change required in your clients.  After the failover, your client still connects tot he old primary IP, which now points to the primary instance that is moved to a healthy zone.

> 5) **Can I use my failover replica as a read slave, or must it just sit idle until an event?**

> Yes. A failover replica is perfectly capable of being served as a read replica.  

> 6) **What happens to the old primary in a failover after it comes back. Does it become a failover replica for the new primary, or do I need to do something by hand?**

> The primary stays as primary before and after the failover process. It is just moved to a healthy zone. Therefore there is no such thing as "old primary comes back" as it always there, and there is nothing need to be done by hand.

> 7) **How do I reset my original primary to be the real master after a failover event is complete?**

> Same as questions 6.

It's not entirely clear how long failover time would be. The second post in that thread mentions InnoDB recovery,
which could take a long time. However, the references pages say that the replica becomes the master,
so I'm guessing that failover should be pretty quick.

#### RDS

[Reference](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.MultiAZ.html#Concepts.MultiAZ.Failover)

Multi-AZ failover on RDS uses a DNS change to point to the standby instance. The reference page mentions
60-120 seconds of unavailability during the failover. Because the standby uses the same storage data
as the primary, there will probably be transaction / log recovery (because failover would've looked
like a crash for MySQL starting up on the standby), so failover time might be longer than that.

## Backups

#### Cloud SQL

[Reference](https://cloud.google.com/sql/docs/mysql/high-availability#how_the_failover_replica_is_configured)

Backups must be done on the master instance.

#### RDS

[Reference](https://aws.amazon.com/rds/details/multi-az/#Increased_Availability)

Backups using the Multi-AZ configuration can be done on the standby because its storage data
is the same as the master's.
