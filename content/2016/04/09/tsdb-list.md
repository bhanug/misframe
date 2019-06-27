---
title: List of Time Series Databases
date: "2016-04-09T23:15:00-07:00"
bestof: true
twitter_card:
  description: "List of all time series databases as I hear about them"
---

**Updated: May 13 2019**

This is not an exhaustive list. If you think I should change something, please leave a comment here
or send me a message on [Twitter](https://twitter.com/PreetamJinka). I'll try to keep it up-to-date
based on feedback and anything new I find. There is a changelog at the end.

**You can also submit a PR to update this page:** https://github.com/Preetam/misframe/blob/master/content/2016/04/09/tsdb-list.md

## Open source

These are either time series databases or general-purpose databases that work well with time series.
Some are layers on top of existing databases.

- [Aerospike](https://www.aerospike.com/)
	- High performance, in-memory, NoSQL
- [Akumuli](https://github.com/akumuli/Akumuli)
	- Written in C++
	- Query language based on JSON over HTTP
	- Can be used as a server application or an embedded library
- [Apache Apex](https://apex.apache.org/)
	- [DataTorrent](https://www.datatorrent.com/)
- [Apache Cassandra](https://cassandra.apache.org/)
	- Or [Scylla](https://www.scylladb.com/), a much faster C++ implementation of Cassandra
	- Distributed, columnar database
	- Has a query language
- [Apache Kudu](https://getkudu.io/) (Incubating)
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
- [Blueflood](https://blueflood.io/)
	- Built on Cassandra
	- Multi-tenant distributed database and metric processing system created by Rackspace
	- Apache 2.0 license
- [Chronix](https://chronix.io/)
	- Built on Apache Lucene, Solr, and Spark
- [CitusDB](https://www.citusdata.com/)
	- Distributed Postgres (through an extension)
- [ClickHouse](https://clickhouse.yandex/)
	- Distributed columnar database
	- Powers [Yandex.Metrica](https://clickhouse.yandex/docs/en/introduction/ya_metrika_task/) (basically Russia's Google Analytics)
- [Cortex](https://github.com/weaveworks/cortex)
	-  Multitenant, horizontally scalable Prometheus as a Service
- [CrateDB](https://crate.io/)
	- Distributed SQL database
	- Fully searchable document oriented data store
	- Uses Presto for SQL, Elasticsearch and Lucene for storage
- [Cube](https://square.github.io/cube/) by Square
	- Built on MongoDB
- [Cyanite](https://cyanite.io/)
	- Compatible with the Graphite ecosystem
	- Stores data in Cassandra
- [Dalmatiner](https://dalmatiner.io/)
	- Built on ZFS and Riak Core
- [Druid](https://druid.io/)
	- Column-oriented open-source distributed data store written in Java
- [Elasticsearch](https://www.elastic.co/blog/elasticsearch-as-a-time-series-data-store)
	- Java & Lucene
	- Support for live time series resampling
	- Distributed data storage
- [EventQL](https://github.com/eventql/eventql)
	- Distributed, columnar database built for large-scale data collection and analytics workloads
	- Supports SQL
- [FiloDB](https://github.com/tuplejump/FiloDB)
	- Distributed, versioned, and columnar analytical database
	- Uses Spark SQL
- [GridDB](https://griddb.net/en/)
	- Horizontally scalable NoSQL DB
- [GridGain](https://www.gridgain.com/)
	- In-memory data fabric
- [Hawkular](https://www.hawkular.org/)
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
- [IoTDB](https://github.com/apache/incubator-iotdb.git)
	- Written in Java
- [KairosDB](https://github.com/kairosdb/kairosdb)
	- Rewrite of OpenTSDB
- [M3DB](https://m3db.io/)
	- Distributed time series database using M3TSZ float64 compression
	- Developed at Uber
- [Newts](https://opennms.github.io/newts/)
	- Based on Cassandra
- [OpenTSDB](https://opentsdb.net/)
	- Built on top of HBase, BigTable, or Cassandra
- [Pinot](https://github.com/linkedin/pinot)
	- Realtime distributed OLAP datastore
	- Horizontally scalable
	- Used at LinkedIn
- [Prometheus](https://prometheus.io/)
	- Monitoring system and TSDB
	- Not distributed
	- Polling-based
- [Riak TS](https://basho.com/products/riak-ts/)
	- Query language
	- Apparently 10x faster than Cassandra (I don't have any more details about this)
- [Roshi](https://github.com/soundcloud/roshi) by SoundCloud
	- Time-series event storage
	- Stateless, distributed layer on top of Redis and is implemented in Go
- [SciDB](https://www.paradigm4.com/)
	- Multidimensional arrays
	- ACID
	- By Michael Stonebraker
- [SiriDB](https://siridb.net/)
	- Written in C and focused on performance
	- Query language
- [sonnerie](https://github.com/njaard/sonnerie)
	- Written in Rust
	- Append only (no updates yet)
	- Simple TCP API
- [Timely](https://github.com/NationalSecurityAgency/timely) by the NSA
	- Backed by Accumulo
- [TimescaleDB](https://www.timescale.com/)
	- Built on PostgreSQL (as an extension)
- [Vulcan](https://github.com/digitalocean/vulcan) by DigitalOcean
	- Extends Prometheus adding horizontal scalability and long-term storage
	- Written in Go
- [Warp 10](https://www.warp10.io/)
	- Geo Time Series database
	- Storage engine: standalone version uses LevelDB, distributed version uses HBase
	- Analytics environment: [WarpScript](https://www.warp10.io/content/03_Documentation/04_WarpScript/01_Concepts) (900+ functions)
	- From [SenX](https://senx.io/)
- [Yuvi](https://github.com/pinterest/yuvi)
	- In-memory storage engine for recent time series metrics data
	- Implemented in Java
	- Supports OpenTSDB metric ingestion and OpenTSDB queries

## Proprietary or internal

These are either proprietary or internal, and not open source.

- [Google BigQuery](https://cloud.google.com/bigquery/)
	- Managed data warehouse for analytics hosted on Google Cloud
	- BigQuery can do lots of things in addition to time series (also see RedShift)
- [Infiniflux](https://infiniflux.com/)
	- Time series DBMS with SQL
- [IRONdb](https://irondb.io/)
	- Scalable storage for a Graphite infrastructure. IRONdb is a new product by Circonus,
	who also created “Snowth” a few years ago (see below).
- [kdb+](https://kx.com/products.php) by Kx Systems
	- Very popular in the financial industry
- [AWS Redshift](https://aws.amazon.com/redshift/)
	- Managed data warehouse for analytics hosted on AWS
	- Redshift can do lots of things in addition to time series (also see BigQuery)
- [Rocana](https://www.rocana.com/) (acquired by Splunk)
	- Proprietary columnar TSDB using Apache Lucene, Kafka, and HDFS
- [eXtremeDB](https://financial.mcobject.com/)
	- Made for financial data
	- Columnar, ACID-compliant, SQL support
- Facebook [Scuba](https://research.facebook.com/publications/scuba-diving-into-data-at-facebook/)
	- Fast, scalable, distributed, in-memory database
- [quasardb](https://www.quasardb.net/)
	- Distributed transactional key-value store with distributed secondary indexes and native time series support
	- Written in C++14
- [SnappyData](https://www.snappydata.io/)
	- fuses Apache Spark with a highly available, multi-tenanted in-memory database
	- OLTP + OLAP on streaming data
- [TempoIQ](https://www.tempoiq.com/)
	- IoT platform
- [VictoriaMetrics](https://github.com/VictoriaMetrics/VictoriaMetrics/wiki/Single-server-VictoriaMetrics)
	- Long-term remote storage for Prometheus

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
- [Facebook Gorilla paper](https://www.vldb.org/pvldb/vol8/p1816-teller.pdf) [PDF]
	- Fast, scalable, in-memory TSDB
- Honeycomb's columnar data store "Retriever"
	- [YouTube video](https://www.youtube.com/watch?v=tr2KcekX2kk) about the design
- [Pulsar](https://gopulsar.io/)
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
- **2018-03-23**  
  Added ClickHouse, M3DB, Honeycomb's data store, quasardb, Cortex, Cyanite, BigQuery, RedShift,
  CrateDB, Pinot;
  misc cleanup and expanded notes.
	- Thanks to Damian Gryski ([@dgryski](https://twitter.com/dgryski)), gilles t ([@akiragt](https://twitter.com/akiragt))
	Misha Brukman ([@MishaBrukman](https://twitter.com/MishaBrukman/status/904536969996374016)),
	Andrew Montalenti ([@amontalenti](https://twitter.com/amontalenti)), and KurtB, Andy Ellicott in the comments.
- **2018-09-13**  
  Added Yuvi, sonnerie. Updated M3DB link to official website.
- **2019-05-13**  
  Added GridDB and VictoriaMetrics.
