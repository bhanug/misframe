---
title: Cistern v0.1.0
date: "2017-08-01T20:50:00-04:00"
---

It's out! You can go download a binary on the GitHub [release page](https://github.com/Cistern/cistern/releases/tag/v0.1.0).

<img src='/img/2017/08/cistern-v0.1.0.png' width=400/>

It's a 100% rewrite (again), but this time I have a much better idea of where to take things.
As I mentioned in the [design notes](/2017/07/15/cistern-design-notes/) post, there's a bunch
of work to do, and this release is a first big step in the right direction.

## Brief overview

Here's how you get started with v0.1.0 and what you can do with it. It only supports VPC Flow Logs.

### Config file

Cistern uses a JSON config file. I have a single CloudWatch Logs group for my VPC Flow Logs
named "flowlogs", so this is what my config file looks like:

```
{
  "cloudwatch_logs": [
    {
      "name": "flowlogs",
      "flowlog": true
    }
  ],
  "retention": 3
}
```

### Starting up Cistern

On startup, Cistern immediately pulls events from the CloudWatch Logs group to catch up
to the most recent data, and then polls for new messages.

```text
$ AWS_REGION=us-east-1 ./cistern
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 795 events
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 805 events
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 877 events
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 971 events
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 869 events
2017/08/01 19:32:07 Flow Logs group flowlogs: aggregated 555 events
```

### Example event

This is what the events look like. Note that not all of the fields from the Flow Log record
are in this event. That's because there's some initial grouping done to turn several Flow Log
records into a single event.

```json
{
  "_id": "1501551000000000|flowlog",
  "_tag": "flowlog",
  "_ts": "2017-08-01T01:30:00Z",
  "bytes": 3895,
  "dest_address": "10.21.155.20",
  "dest_port": 36439,
  "packets": 10,
  "protocol": 6,
  "source_address": "172.31.31.192",
  "source_port": 443
}
```

Now that you know the structure of the events, let's take a look at some queries.

### Queries

Cistern now has a CLI program with uses the HTTP API. You can now write "queries" that resemble
SQL. Cistern previously only supported time series, but v0.1.0 is built on an events architecture
so you can do more fancy things like filtering, grouping and aggregation, ranking, and still
generate time series!

Here's a simple query that calculates the sum of all of the `bytes` fields from the Unix timestamp
1501025801 to now.

```text
$ cistern-cli -collection flowlogs -start 1501025801 -columns 'sum(bytes)'
{
  "summary": [
    {
      "sum(bytes)": 11439463
    }
  ],
  "query": {
    "columns": [
      {
        "aggregate": "sum",
        "name": "bytes"
      }
    ],
    "descending": false,
    "time_range": {
      "end": "2017-08-01T19:37:03-04:00",
      "start": "2017-07-25T19:36:41-04:00"
    }
  }
}
```

(The `query` section of the output is the JSON representation of the query, and it's only
used for debugging. I'll omit that section for the rest of my examples.)

I can aggregate multiple columns at the same time too.
`count(_id)` just counts the number of events, like `count(*)`.

```text
$ cistern-cli -collection flowlogs -start 1501025801 \
  -columns 'sum(bytes), count(_id), max(packets), max(bytes)'
{
  "summary": [
    {
      "count(_id)": 893,
      "max(bytes)": 219376,
      "max(packets)": 180,
      "sum(bytes)": 11439463
    }
  ]
}
```

I can also group these aggregates by another field, like `protocol`.
Now I can see the breakdown of those aggregates by TCP, UDP, and ICMP events.

```text
$ cistern-cli -collection flowlogs -start 1501025801 \
  -columns 'sum(bytes), count(_id), max(packets), max(bytes)' \
  -group 'protocol'
{
  "summary": [
    {
      "count(_id)": 823,
      "max(bytes)": 219376,
      "max(packets)": 180,
      "protocol": 6,
      "sum(bytes)": 11429888
    },
    {
      "count(_id)": 57,
      "max(bytes)": 1280,
      "max(packets)": 7,
      "protocol": 17,
      "sum(bytes)": 8557
    },
    {
      "count(_id)": 13,
      "max(bytes)": 107,
      "max(packets)": 2,
      "protocol": 1,
      "sum(bytes)": 1018
    }
  ]
}
```

I'm also not limited to a single field to group by. Maybe I'm interested in the aggregates
grouped by source address *and* destination address. But that could generate a lot of groups,
so I can trim down my results by passing in an `order-by` parameter, as well as a limit.

Now I'm getting the top 3 (source, destination) groups by total bytes.

```text
$ cistern-cli -collection flowlogs -start 1501025801 \
  -columns 'sum(bytes), count(_id), max(packets), max(bytes)' \
  -group 'source_address, dest_address' \
  -order-by 'sum(bytes)' \
  -limit 3 \
  -descending
{
  "summary": [
    {
      "count(_id)": 111,
      "dest_address": "52.54.154.173",
      "max(bytes)": 217333,
      "max(packets)": 180,
      "source_address": "172.31.31.192",
      "sum(bytes)": 6000569
    },
    {
      "count(_id)": 49,
      "dest_address": "52.54.236.132",
      "max(bytes)": 219376,
      "max(packets)": 177,
      "source_address": "172.31.31.192",
      "sum(bytes)": 3110402
    },
    {
      "count(_id)": 128,
      "dest_address": "172.31.31.192",
      "max(bytes)": 21338,
      "max(packets)": 73,
      "source_address": "52.54.154.173",
      "sum(bytes)": 756098
    }
  ]
}
```

Finally, we can't forget about time series! Here's an example where I'm grouping by
`protocol` and aggregating the sum of `packets` and `bytes` in `24h` (24 hour) ranges.
I also added a filter which only selects events where the `source_address` is `"172.31.31.192"`.

```text
$ cistern-cli -collection flowlogs -start 1501025801 \
  -columns 'sum(packets), sum(bytes)' \
  -group 'protocol' \
  -point-size 24h \
  -filters 'source_address eq "172.31.31.192"'
{
  "summary": [
    {
      "protocol": 17,
      "sum(bytes)": 532,
      "sum(packets)": 7
    },
    {
      "protocol": 1,
      "sum(bytes)": 136,
      "sum(packets)": 4
    },
    {
      "protocol": 6,
      "sum(bytes)": 9898528,
      "sum(packets)": 12995
    }
  ],
  "series": [
    {
      "_ts": "2017-07-25T00:00:00Z",
      "protocol": 6,
      "sum(bytes)": 2497,
      "sum(packets)": 18
    },
    {
      "_ts": "2017-07-26T00:00:00Z",
      "protocol": 17,
      "sum(bytes)": 76,
      "sum(packets)": 1
    },
    {
      "_ts": "2017-07-26T00:00:00Z",
      "protocol": 6,
      "sum(bytes)": 1090358,
      "sum(packets)": 1786
    },
    {
      "_ts": "2017-07-27T00:00:00Z",
      "protocol": 6,
      "sum(bytes)": 2357328,
      "sum(packets)": 2706
    },
    {
      "_ts": "2017-07-28T00:00:00Z",
      "protocol": 1,
      "sum(bytes)": 68,
      "sum(packets)": 2
    },
	// ...
  ]
}
```

## Up next

The next release will still focus on AWS CloudWatch Logs, but will support generic JSON log messages,
so you can run the same kinds of queries I just demonstrated on any JSON objects. That's the kind of
feature I'd really like to use on a daily basis.
