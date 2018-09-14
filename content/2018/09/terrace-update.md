---
title: Terrace Storage Experiment Update
date: "2018-09-05T13:05:00-07:00"
twitter_card:
  description: "I am working on a storage format for time series data."
  image: "https://misfra.me/img/2018/terrace.png"
---

### MISC

* Need to emphasize: this is not a problem that can be solved by infra

### What is Terrace?

Terrace is my new storage format for events and time series.

### Structure of Terrace files

A Terrace file is made up of nested levels. Each level describes a subset
of events that have similar attributes, like an attribute with a specific value.


![](/img/2018/terrace_level_visualization.png)

An event in the top right highlighted sublevel would look like this:

```json
{
  "timestamp": 1536218943,
  "region": "us-west-1",
  "server": "db-1",
  "level": "INFO",
  "query": "SELECT * FROM users",
  "latency": 0.48
}
```

<img width=200 src='/img/2018/terrace_generation_steps.png'>


## Time Series Benchmarking Suite (TSBS)

The folks at TimescaleDB created a set of tools that form a framework
to test out time series storage solutions. It's adapted from an [earlier
work](https://github.com/influxdata/influxdb-comparisons) published by InfluxData.
You can read more about it [here](https://blog.timescale.com/time-series-database-benchmarks-timescaledb-influxdb-cassandra-mongodb-bc702b72927e).

I made a [fork](https://github.com/Preetam/tsbs/tree/json) of this project to
print generated time series events as JSON. I used these JSON events to work on the Terrace generator.

---

## Cost calculation
