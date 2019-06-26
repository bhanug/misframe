---
title: S3 is not just for backups
date: "2017-02-10T00:30:00-05:00"
---

For a long time, I treated S3 as a backup location for databases. After spending the last few months
thinking about cloud architecture and designing Epsilon, I slowly learned that services like S3 can
be a core storage component in a cloud architecture.

Here's the gist of my previous thinking:

On a single instance of a database like MySQL, data are in two places: cached in the buffer pool in
memory or on disk.

![Buffer pool and disk](/img/2017/02/buffer-pool-disk.png)

The disk is considered to be durable storage. For backups, depending on the strategy you use, you
either end up with a single snapshot of the data or incremental pieces. Everything ends up going to
S3 (or some equivalent) and stays untouched until you actually need to use the backups (maybe for
disaster recovery).

Now, I believe that S3 can play a way more active role. What you can also do is add S3 as another
storage tier.

![Buffer pool, disk, and S3](/img/2017/02/buffer-pool-disk-s3.png)

Instead of treating S3 as some continuous backup destination, you can actually treat it like the
system of record, i.e. what your storage system actually is. At that point, the disk starts to look
like a cache, just like the buffer pool. But unlike the buffer pool, it's a durable cache. *But
does it need to be?*

At [VividCortex](https://www.vividcortex.com/), all of our "big" data are stored in Kafka before
reaching MySQL. If we had continuous backups to S3, we actually wouldn't require MySQL to be fully
durable. There would be data that are in MySQL that haven't reached S3 yet, but those are already
durable in Kafka.

In MySQL terms, a MySQL instance just holds dirty data before checkpointing to S3. If we have a crash
or instance failure, we recover using the Kafka log. Just like how pages can be read into the buffer
pool on-demand from the disk, you can do the same using S3.

Redshift apparently does that already:

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Redshift pages data in from S3.<a href="https://t.co/C8z2aXyPoz">https://t.co/C8z2aXyPoz</a> <a href="https://t.co/jFWl3wRg3j">pic.twitter.com/jFWl3wRg3j</a></p>&mdash; Preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/829185425755471872">February 8, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

With this approach, you can scale compute independently of storage, which I think is a *huge* win.

---

If you want to learn more, I suggest watching Netflix's talk from re:Invent 2016:

**"Using Amazon S3 as the fabric of our big data ecosystem."**

YouTube: https://www.youtube.com/watch?v=o52vMQ4Ey9I

Slides:

<iframe src="//www.slideshare.net/slideshow/embed_code/key/go7q2pCfjYhPx" width="595" height="485" frameborder="0" marginwidth="0" marginheight="0" scrolling="no" style="border:1px solid #CCC; border-width:1px; margin-bottom:5px; max-width: 100%;" allowfullscreen> </iframe> <div style="margin-bottom:5px"> <strong> <a href="//www.slideshare.net/AmazonWebServices/aws-reinvent-2016-netflix-using-amazon-s3-as-the-fabric-of-our-big-data-ecosystem-bdm306" title="AWS re:Invent 2016: Netflix: Using Amazon S3 as the fabric of our big data ecosystem (BDM306)" target="_blank">AWS re:Invent 2016: Netflix: Using Amazon S3 as the fabric of our big data ecosystem (BDM306)</a> </strong> from <strong><a target="_blank" href="//www.slideshare.net/AmazonWebServices">Amazon Web Services</a></strong> </div>
