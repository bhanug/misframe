---
title: Solving Time Series Storage with Brute Force
date: "2015-07-03"
---

Many months ago, I first read "Searching 20 GB/sec: Systems Engineering Before Algorithms" <a class='ref' href='#endnote-ref-1'>[1]</a>, an excellent post by Steve at Scalyr. I re-read it six months ago when I was on winter break between semesters of college. I was traveling, working on Cistern <a class='ref' href='#endnote-ref-2'>[2]</a>, and thinking about time series storage.

By "thinking," I mean "annoyed." I was using Bolt <a class='ref' href='#endnote-ref-3'>[3]</a> (the B+tree-based, transactional storage engine implemented in Go) to store Cistern's time series. The implementation I had was pretty inefficient. There was one key-value pair per time series point without any compression. There was a significant amount of overhead with this approach, especially considering my time series points were 12-bytes at most (8-byte timestamp plus 4-byte float value). I could have implemented compression by packing multiple points into a single value, and then compressing afterwards, but then I'd have to implement batching and all of the compression logic myself.

At some point I just went, "forget B-trees, forget log-structured merge, forget everything!" and decided to try to brute force it. I just need to store a bunch of arrays. How hard could that be, right?

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr">Building a ridiculously simple, &quot;brute force&quot; time series engine <a href="https://t.co/g2cPnvVxwp">pic.twitter.com/g2cPnvVxwp</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/551245706242322433">January 3, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Fast forwarding to the present, you may have seen my blog post introducing Catena <a class='ref' href='#endnote-ref-4'>[4]</a>, heard me talk about it at a meetup or watched a webinar <a class='ref' href='#endnote-ref-5'>[5]</a>, or heard about it from some other source. There are so many topics that get covered when I talk about Catena, and you may think that it took a lot of researching time to come up with its design. That's isn't the case. The truth is that it has mainly been brute forced! Problems and challenges came up as I was implementing it, and they needed to be solved.

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr">As my work on this storage engine progresses... <a href="https://t.co/k0MWIDwJte">pic.twitter.com/k0MWIDwJte</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/553659333104640001">January 9, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

The *really interesting* thing now is that the current design is not unique at all. In fact, it could probably be replicated with RocksDB. How about that? I started from scratch and now it seems like the path leads to RocksDB.

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr"><a href="https://twitter.com/xaprb">@xaprb</a> <a href="https://twitter.com/pauldix">@pauldix</a> <a href="https://twitter.com/PreetamJinka">@PreetamJinka</a> <a href="https://twitter.com/RocksDB">@RocksDB</a> at a high level -- column family per time range, disable auto compaction</p>&mdash; markcallaghan (@markcallaghan) <a href="https://twitter.com/markcallaghan/status/606450119421140992">June 4, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I'd like to remind you that I didn't want to write a time series storage engine. I guess it's what I'm mostly known for now, but that was a complete accident. I never expected this to actually work.

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr">HOLY **** IT WORKS&#10;&#10;A WAL is getting checked twice for some reason though xD <a href="https://t.co/Dk3bYhadcD">pic.twitter.com/Dk3bYhadcD</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/553966524940419072">January 10, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

But it does. Cistern no longer uses Bolt and has been using Catena for a while. I now have time series storage with compression, and it works well enough for me to build a dashboard that I can use to monitor my systems.

<blockquote class="twitter-tweet" lang="en"><p lang="en" dir="ltr">Yay, I can make a dashboard with a custom storage engine! <a href="https://t.co/53D1tSHvHJ">pic.twitter.com/53D1tSHvHJ</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/554170122668355584">January 11, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I guess the moral of the story is that brute force can yield some interesting results. You may end up finding a solution to a problem, like I did when I finally got a better time series storage engine for Cistern. Or, more importantly, you may find that it leads you in a direction that may not have been so obvious when you started off. So, what are the next steps? I've already solved my own time series storage problem, but if I had something that needed to be "production ready," I'd start looking at RocksDB.

<br/>

<a class='endnote' name='endnote-ref-1'>[1]</a> [Searching 20 GB/sec: Systems Engineering Before Algorithms](https://blog.scalyr.com/2014/05/searching-20-gbsec-systems-engineering-before-algorithms/)  
<a class='endnote' name='endnote-ref-2'>[2]</a> [Cistern](https://preetam.github.io/cistern/)  
<a class='endnote' name='endnote-ref-3'>[3]</a> [Bolt](https://github.com/boltdb/bolt)  
<a class='endnote' name='endnote-ref-4'>[4]</a> [State of the State Part III](https://misfra.me/state-of-the-state-part-iii/)  
<a class='endnote' name='endnote-ref-5'>[5]</a> [Catena: A High-Performance Time Series Storage Engine](https://www.slideshare.net/vividcortex/catena-a-highperformance-time-series-data)  
