---
title: Adversaria.
date: "2013-01-03"
url: /adversaria
---

Everyone (or almost everyone) uses [RRDtool](http://oss.oetiker.ch/rrdtool/) to store traffic data. I wanted something a little different.

Couchbase woes!
---

I started using Couchbase to log traffic. I had a document for each data point and was able to run map-reduce queries with Couchbase views. It worked great for a couple of months and then for some reason... weird stuff happened. Over 5,000 documents stopped getting mapped. That's not good.

I used a 4-node Couchbase cluster, because why not? :P
I removed a node, started the rebalancing process, and then added the node back. After rebalancing the documents again, everything seemed to work. Awesome. Woo, fault tolerance!

Then stuff got weird again. For some reason after I hit 40,000 documents, it would never finish indexing the bucket to enable querying. Now the problem was much worse. I couldn’t query ANYTHING. I just left it for a day and came back to find an entire node out of disk space. I found a 32 GB index file. A *32 GB index file.* My entire bucket was < 50 MB. Whoops. Not sure how that happened...

Well, I need a better way to monitor traffic.

Adversaria
---
Adversaria means “a miscellaneous collection of notes, remarks, or selections; a commonplace book; also, commentaries or notes.” It’s just a little storage engine for traffic data.

Features:

- RRD / circular buffer (fixed size on disk, just like RRDtool)
- Data stored in binary
- Range reads
- Antichronological inserts (probably will be removed)
- JSON output

<pre><code>$ adversaria export /tmp/foo.db 1357245300 1357245999
{
	"1357245306":[0.36529, 2.94072],
	"1357245606":[0.48615, 1.18695],
	"1357245906":[0.21887, 0.9282]
}
</code></pre>

Internally there are multiple "circular buffers". I keep 3 months worth of samples with 5-minute resolution as the primary buffer and the secondary buffer stores 9 months worth of hourly samples. It’s really simple but really useful.

Source code?
----
[GitHub!](https://github.com/PreetamJinka/adversaria-java)
