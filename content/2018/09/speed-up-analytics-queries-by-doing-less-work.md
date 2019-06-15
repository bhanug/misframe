---
title: Speed up analytics queries by doing less work
date: "2018-09-17T23:45:00-07:00"
twitter_card:
  description: "Here are my top 3 techniques to speed up analytics queries"
  image: "https://misfra.me/img/2018/analytics_speedup_diagram_1.svg"
---

Often times in my work I come across queries like this,

```sql
SELECT SUM(count) FROM data WHERE foo = 'bar'
  AND time >= '2018-08-13 00:00:00+00'
  AND time  < '2018-09-13 00:00:00+00'
```

which is getting a total count for a month with a filter, and they
take a long time to run because there's a significant number
of rows in the `data` table but only a few have `foo = 'bar'`.

<img src='/img/2018/analytics_speedup_diagram_1.svg'/>

We want these queries to execute faster, which is the same as reducing their
latency. Remember that **latency comes from doing work or waiting**. This post
is about modifying the work aspect because analytics queries don't really have
problems with waiting (at least in my experience, but there are exceptions).

<!--more-->

There are two ways to improve how work gets done to speed up your queries:

1. **Do work faster** by allocating more resources (parallelizing execution,
using faster hardware, etc.)
2. **Do less work**

This post is about #2. These approaches aren't mutually exclusive so you can
always try to apply #1 after the techniques I describe below. I'm also putting
this post in a relational database context but these techniques apply generally.
There are many other techniques I can talk about but these are the ones that
stand out the most from what I learned in the past few years.

## Indexing

**Required code changes:** None

The first approach to making this query faster is familiar to those used to
SQL databases. Adding an index on `foo` and using it for your query
lets you do less work by skipping to only the data you're interested in.

<img src='/img/2018/analytics_speedup_diagram_2.svg'/>

If you're using a relational database, you probably won't need any code
changes. You can just add a new index to a table and immediately get benefits.
Keep in mind that indexing has downsides like increased space usage and
more work required for writes because you're making a significant change
to how data is stored.

## Store and use additional metadata

**Required code changes:** Minimal

This approach is less generalized but has been super useful in my work.
You can try to store extra metadata about your data that you can reference
later to reduce the amount of data you need to consider. For example, if you
can store the first time `foo = 'bar'`, you can use that information to truncate
time ranges in your queries to begin at that time without affecting the result.

<img src='/img/2018/analytics_speedup_diagram_3.svg'/>

I say this has minimal code change requirements since I find that we're often
storing this information anyway, and we just have to start looking it up for other
queries.

## Precompute some information

**Required code changes:** Significant

The last approach is precomputation. Looking at my query, notice how my time range
is a month. If the older data isn't being updated often, I can just precompute those
counts in, for example, 1 day aggregates. That way I can do the work ahead of time and
use the results in the future.

I say this has significant code change requirements because it can be tricky to
create rollup processes and implement the logic to handle multiple granularities.
