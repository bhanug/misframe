---
title: List of Time Series Databases
date: "2017-09-03T16:29:53.791Z"
bestof: true
---

**Updated: September 3 2017**

This is not an exhaustive list. If you think I should change something, please leave a comment here
or send me a message on [Twitter](https://twitter.com/PreetamJinka). I'll try to keep it up-to-date
based on feedback and anything new I find. There is a changelog at the end.

## Open source

These are either time series databases or general-purpose databases that work well with time series.
Some are layers on top of existing databases.

- [Aerospike](http://www.aerospike.com/)
	- High performance, in-memory, NoSQL
- [Akumuli](https://github.com/akumuli/Akumuli)
	- Written in C++
	- Query language based on JSON over HTTP
	- Can be used as a server application or an embedded library
- [Apache Apex](https://apex.apache.org/)
	- [DataTorrent](https://www.datatorrent.com/)
- [Apache Cassandra](http://cassandra.apache.org/)
	- Or [Scylla](http://www.scylladb.com/), a much faster C++ implementation of Cassandra
	- Distributed, columnar database
	- Has a query language
- [Apache Kudu](http://getkudu.io/) (Incubating)
	- Columnar, part of the Hadoop stack
	- "fast analytics on fast data"
- [Atlas](https://github.com/Netflix/atlas) by Netflix
	- Written in Scala
	- In memory
	- Stack language for queries
- [Axibase Time Series Database](https://axibase.com/products/axibase-time-series-database/)
	- Visualizations, rules engine, forecasting
- [Beringei](https://github.com/facebookincubator/beringei) by Facebook
	- In memory
	- Open source implementation of ideas presented in their Gorilla paper (link below)
- [Blueflood](http://blueflood.io/)
	- Built on Cassandra
	- Multi-tenant distributed database and metric processing system created by Rackspace
	- Apache 2.0 license
- [Chronix](http://chronix.io/)
	- Built on Apache Lucene, Solr, and Spark
- [CitusDB](https://www.citusdata.com/)
	- Distributed Postgres (through an extension)
- [Cube](http://square.github.io/cube/) by Square
	- Built on MongoDB
- [Dalmatiner](https://dalmatiner.io/)
	- Built on ZFS and Riak Core
- [Druid](http://druid.io/)
	- Column-oriented open-source distributed data store written in Java
- [Elasticsearch](https://www.elastic.co/blog/elasticsearch-as-a-time-series-data-store)
- [EventQL](https://github.com/eventql/eventql)
	- Distributed, columnar database built for large-scale data collection and analytics workloads
	- Supports SQL
- [FiloDB](https://github.com/tuplejump/FiloDB)
	- Distributed, versioned, and columnar analytical database
	- Uses Spark SQL
- [GridGain](http://www.gridgain.com/)
	- In-memory data fabric
- [Hawkular](http://www.hawkular.org/)
	- Open source monitoring solution by Red Hat
	- Metrics storage uses Cassandra
- [HBase](https://hbase.apache.org/)
	- distributed database for very large tables
	- Related: Google Cloud [BigTable](https://cloud.google.com/bigtable/) (hosted)
- [Heroic](https://github.com/spotify/heroic) by Spotify
	- Based on Bigtable, Cassandra, and Elasticsearch
- [InfluxDB](https://influxdata.com/)
	- Written in Go
	- Clustering is a paid feature now
- [KairosDB](https://github.com/kairosdb/kairosdb)
	- Rewrite of OpenTSDB
- [Newts](https://opennms.github.io/newts/)
	- Based on Cassandra
- [OpenTSDB](http://opentsdb.net/)
	- Built on top of HBase
- [Prometheus](http://prometheus.io/)
	- Monitoring system and TSDB
	- Not distributed
	- Polling-based
- [Riak TS](http://basho.com/products/riak-ts/)
	- Query language
	- Apparently 10x faster than Cassandra (I don't have any more details about this)
- [Roshi](https://github.com/soundcloud/roshi) by SoundCloud
	- Time-series event storage
	- Stateless, distributed layer on top of Redis and is implemented in Go
- [SciDB](http://www.paradigm4.com/)
	- Multidimensional arrays
	- ACID
	- By Michael Stonebraker
- [SiriDB](http://siridb.net/)
	- Written in C and focused on performance
	- Query language
- [Timely](https://github.com/NationalSecurityAgency/timely) by the NSA
	- Backed by Accumulo
- [TimescaleDB](http://www.timescale.com/)
	- Built on PostgreSQL (as an extension)
- [Vulcan](https://github.com/digitalocean/vulcan) by DigitalOcean
	- Extends Prometheus adding horizontal scalability and long-term storage
	- Written in Go
- [Warp 10](http://www.warp10.io/)
	- Distributed version uses HBase
	- From Cityzen Data

## Proprietary or internal

These are either proprietary or internal, and not open source.

- [Cityzen Data](http://www.cityzendata.com/)
	- IoT / sensor data platform
- [Infiniflux](http://infiniflux.com/)
	- Time series DBMS with SQL
- [IRONdb](https://www.circonus.com/irondb/)
	- Scalable storage for a Graphite infrastructure. IRONdb is a new product by Circonus,
	who also created “Snowth” a few years ago (see below).
- [kdb+](https://kx.com/products.php) by Kx Systems
	- Very popular in the financial industry
- [Rocana](https://www.rocana.com/)
	- Proprietary columnar TSDB using Apache Lucene, Kafka, and HDFS
- [eXtremeDB](http://financial.mcobject.com/)
	- Made for financial data
	- Columnar, ACID-compliant, SQL support
- Facebook [Scuba](https://research.facebook.com/publications/scuba-diving-into-data-at-facebook/)
	- Fast, scalable, distributed, in-memory database
- [SnappyData](http://www.snappydata.io/)
	- fuses Apache Spark with a highly available, multi-tenanted in-memory database
	- OLTP + OLAP on streaming data
- [TempoIQ](https://www.tempoiq.com/)
	- IoT platform

## Things to look at for ideas


These are not exactly TSDBs, but are interesting resources to take a look at.

- [Apache Arrow](https://github.com/apache/arrow/)
	- Columnar in-memory format and API
- [Apache Drill](https://drill.apache.org/)
	- "Schema-free SQL Query Engine for Hadoop, NoSQL and Cloud Storage"
- [Apache Parquet](https://parquet.apache.org/)
	- Columnar storage format for HDFS
- [BTrDB paper](https://www.usenix.org/system/files/conference/fast16/fast16-papers-andersen.pdf) [PDF]
	- "Optimizing Storage System Design for Timeseries Processing"
- Circonus Snowth
	- [YouTube video](https://www.youtube.com/watch?v=hwHpd20NciE) about the design
- [Facebook Gorilla paper](http://www.vldb.org/pvldb/vol8/p1816-teller.pdf) [PDF]
	- Fast, scalable, in-memory TSDB
- [Pulsar](http://gopulsar.io/)
	- Streaming SQL
- Square [metrics query engine](https://github.com/square/metrics)

---

### Changelog

- **2016-04-09**  
  Initial version
	- Thanks to Csaba Csoma and Damian Gryski ([@dgryski](https://twitter.com/dgryski)) for their
	contributions.
	- Added Apache Drill, Kudu (thanks [Mark Papadakis](https://twitter.com/markpapadakis))
- **2016-04-10**  
  Added Cityzen Data, Hawkular, Infiniflux, TempoIQ, kdb+
	- Thanks to @pganti in the comments
- **2017-01-30**  
  Added SciDB, SiriDB
	- Thanks to @Pranas and @ps22 in the comments
- **2017-04-05**  
  Added Akumuli, Atlas, Beringei, Chronix, Roshi, Timely, TimescaleDB, Vulcan;  
  Ordered by name
	- Thanks to Damian Gryski ([@dgryski](https://twitter.com/dgryski)) and Khalid Lafi
	([@LafiKL](https://twitter.com/LafiKL)) for their contributions.
- **2017-09-03**  
  Added EventQL, eXtremeDB, IRONdb; reorganized sections.
