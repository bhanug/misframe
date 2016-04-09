---
title: List of Time Series Databases
date: "2016-04-09T16:29:53.791Z"
---

This is not an exhaustive list. If you think I should add something, please leave a comment here
or send me a message on [Twitter](https://twitter.com/PreetamJinka). I'll try to keep it up-to-date
based on feedback and anything new I find. There is a changelog at the end.

## Existing Solutions

These are either time series databases or general-purpose databases that work well with time series.
Some are layers on top of existing databases.

- [Apache Cassandra](http://cassandra.apache.org/)
	- Or [Scylla](http://www.scylladb.com/), a much faster C++ implementation of Cassandra
	- Distributed, columnar database
	- Has a query language
- [HBase](https://hbase.apache.org/)
	- distributed database for very large tables
	- Related: Google Cloud [BigTable](https://cloud.google.com/bigtable/) (hosted)
- Apache Apex
	- [DataTorrent](https://www.datatorrent.com/)
- [InfluxDB](https://influxdata.com/)
	- Written in Go
	- Clustering is a paid feature now
- [GridGain](http://www.gridgain.com/)
	- In-memory data fabric
- [CitusDB](https://www.citusdata.com/)
	- Distributed Postgres (through an extension)
- [Aerospike](http://www.aerospike.com/)
	- High performance, in-memory, NoSQL
- [Dalmatiner](https://dalmatiner.io/)
	- Built on ZFS and Riak Core
- [FiloDB](https://github.com/tuplejump/FiloDB)
	- Distributed, versioned, and columnar analytical database
	- Uses Spark SQL
- [Elasticsearch](https://www.elastic.co/blog/elasticsearch-as-a-time-series-data-store)
- [OpenTSDB](http://opentsdb.net/)
	- Built on top of HBase
- [KairosDB](https://github.com/kairosdb/kairosdb)
	- Rewrite of OpenTSDB
- [Blueflood](http://blueflood.io/)
	- Built on Cassandra
	- Multi-tenant distributed database and metric processing system created by Rackspace
	- Apache 2.0 license
- [Druid](http://druid.io/)
	- Column-oriented open-source distributed data store written in Java
- [Prometheus](http://prometheus.io/)
	- Monitoring system and TSDB
	- Not distributed
	- Doesn't support out-of-order writes (since it's based on polling)
- [Newts](https://opennms.github.io/newts/)
	- Based on Cassandra
- [Warp 10](http://www.warp10.io/)
	- Distributed version uses HBase
	- From Cityzen Data
- [Heroic](https://github.com/spotify/heroic) by Spotify
	- Based on Bigtable, Cassandra, and Elasticsearch
- [Cube](http://square.github.io/cube/) by Square
	- Built on MongoDB
- [Riak TS](http://basho.com/products/riak-ts/)
	- Not officially open-source yet
	- Apparently 10x faster than Cassandra (I don't have any more details about this)
- [Axibase Time Series Database](https://axibase.com/products/axibase-time-series-database/)
	- Visualizations, rules engine, forecasting
- [Apache Kudu](http://getkudu.io/) (Incubating)
	- Columnar, part of the Hadoop stack
	- "fast analytics on fast data"

## Things to look at for ideas

These are either proprietary or internal, or not TSDBs.

- [Rocana](https://www.rocana.com/)
	- Proprietary columnar TSDB using Apache Lucene, Kafka, and HDFS
- [SnappyData](http://www.snappydata.io/)
	- fuses Apache Spark with a highly available, multi-tenanted in-memory database
	- OLTP + OLAP on streaming data
- Circonus Snowth
	- [YouTube video](https://www.youtube.com/watch?v=hwHpd20NciE) about the design
- [Pulsar](http://gopulsar.io/)
	- Streaming SQL
- Facebook [Scuba](https://research.facebook.com/publications/scuba-diving-into-data-at-facebook/)
	- Fast, scalable, distributed, in-memory database
- Square [metrics query engine](https://github.com/square/metrics)
- [Facebook Gorilla paper](http://www.vldb.org/pvldb/vol8/p1816-teller.pdf) [PDF]
	- Fast, scalable, in-memory TSDB
- [BTrDB paper](https://www.usenix.org/system/files/conference/fast16/fast16-papers-andersen.pdf) [PDF]
	- "Optimizing Storage System Design for Timeseries Processing"
- [Apache Arrow](https://github.com/apache/arrow/)
	- Columnar in-memory format and API
- [Apache Parquet](https://parquet.apache.org/)
	- Columnar storage format for HDFS
- [Apache Drill](https://drill.apache.org/)
	- "Schema-free SQL Query Engine for Hadoop, NoSQL and Cloud Storage"
---

### Changelog

- 2016-04-09: Initial version
	- Thanks to Csaba Csoma and Damian Gryski ([@dgryski](https://twitter.com/dgryski)) for their
	contributions.
	- Added Apache Drill, Kudu (thanks [Mark Papadakis](https://twitter.com/markpapadakis))
