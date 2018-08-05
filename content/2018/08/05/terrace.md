---
title: Terrace Time Series Storage Experiment
date: "2018-08-05T13:05:00-07:00"
twitter_card:
  description: "I am working on a storage format for time series data."
  image: "https://misfra.me/img/2018/terrace.png"
---

I'm working on time series storage again!

My last "serious" time series storage project was Catena that I blogged about
[here](/state-of-the-state-part-iii/).
I wrote Catena to store time series data for a monitoring project I had, and the only
use case I had was plotting charts. Since then my requirements have changed. My time series
have to be more than arrays of points with simple string names. They have to be computed
from _events_ with lots of attributes that I can filter, group, aggregate, and rank.

What are events? They're basically maps with timestamps, like the following log message:

```json
{
  "latency": 0.438107,
  "level": "info",
  "method": "GET",
  "msg": "[Req 9090d339] status code 200, latency 0.44 ms",
  "request_id": "9090d339",
  "status": 200,
  "time": "2018-08-03T00:27:36Z",
  "url": "/goals/ZxbWA79Q"
}
```

And an example (pseudocode) query I could write against events like that is the following,
which should return time series points representing the 95th percentile latency for
HTTP 200 requests in 1 hour intervals.

```txt
SELECT QUANTILE(`latency`, 0.95) WHERE `status` = 200 POINT SIZE 1h;
```

Here are some more requirements:

* "Cloud native" storage. This means the source of truth has to be an object store like S3.
* Support for high cardinality. Events can have arbitrary columns and columns can have arbitrary
  cardinality.
* SQL-like querying. There needs to be grouping, aggregation, ranking, etc.
* Customizable infra -- I should be able to implement replication and backups at a higher
  level. Most time series storage solutions limit you to a specific type of infrastructure
  which isn't the best for all use cases.

## The Terrace Experiment

<img src='/img/2018/terrace.png' width=250/>

This new project is called Terrace and it's an immutable storage format for events. Other
immutable storage formats are Parquet, ORC, and my own Catena's on-disk partition format.
The name comes from how I imagine the storage format being implemented, with lots of layers.
Besides addressing the requirements I listed above, I want to implement automatic indexing.
That's going to be the most interesting part about this project, but I'm not sure if it's going
to be worth it. That's why I'm considering this to be an _experiment_. I'm excited about this
since I don't think anyone else doing something like this, and it's a combination of my
favorite things: databases, time series, and machine learning.

### Automatic Indexing

When you work with a relational database, you think about your data model, define your schema
and indexes, and load your data. At query time, the database engine will consider statistics
about your data and information about available indexes to generate an execution plan. The
execution plan is limited to the indexes you've defined ahead of time. For this new project,
I want to flip that around. Because I'm working with immutable data, I can let an algorithm
analyze the data and determine how to structure and index the data for me.

What I mean by _automatic indexing_ is building up index structures with column ordering
and split points optimizied for certain queries on a data set.

```txt
                  Queries                          Data
 __________________________________________     __________
| SELECT COUNT(*);                         |   |    {}    |     ___________
| SELECT MAX(latency) WHERE user = 2;      |   |    {}    |    | Optimized |
| SELECT user, COUNT(login_time) AS logins | + |    {}    | => | Terrace   |
|   GROUP BY user                          |   | Raw data |    | File      |
|   ORDER BY logins                        |   |    {}    |     ^^^^^^^^^^^
|   DESC LIMIT 10;                         |   |    {}    |
 ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^     ^^^^^^^^^^
```

How will this be done? Machine learning, of course :).

Most of my time with this project will be spent working on automatic indexing and determining
whether or not it's worth it to spend a large amount of compute resources to optimize the
storage and access of immutable event data. Who knows, [maybe brute force is actually more
efficient](https://blog.scalyr.com/2014/05/searching-20-gbsec-systems-engineering-before-algorithms/).

### Storage format only

For the past few years I've been focusing _a lot_ on time series storage infrastructure.
_Infrastructure is not the problem._ We've basically figured that one out. What we haven't
really figured out is how to build a truly cloud-native time series storage system. Cloud-native
means separation of storage and compute, and I think we haven't spent enough time on improving storage
_independent_ of compute. To focus on just that I'm going to work _only_ on the storage format. This
also narrows the scope of the problem to something I can work on myself. Others can figure out infra;
it's where things are really opinionated anyway.

### Next steps

So far I have an idea and have done lots of research. The next steps are to start iterating on the
storage format and work on analyzing data for automatic indexing. I also need test data, so I'll
probably have to write a test data generator. I'll share my progress on Twitter and more posts.
