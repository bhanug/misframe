---
title: Docker Compose
date: "2017-06-19T22:30:00-04:00"
---

I think it's rare that I get excited about software these days. Things don't seem as magical
as they used to. But I tried out [Docker Compose](https://docs.docker.com/compose/) recently and
*whoa*.

I cloned [github.com/singram/mongo-docker-compose](https://github.com/singram/mongo-docker-compose)
and ran `docker-compose up` and *boom*. I had a full set of MongoDB containers running on my laptop
and talking to each other.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Docker Compose is neat! A whole MongoDB cluster (13 containers!) on my laptop.<a href="https://t.co/IUz2Oaum8Y">https://t.co/IUz2Oaum8Y</a> <a href="https://t.co/9bh7U0TqAi">pic.twitter.com/9bh7U0TqAi</a></p>&mdash; preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/874620583300014080">June 13, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

These are *Linux* containers running MongoDB on my *macOS* laptop. They're automatically able to connect
to each other. This is magical!

A coworker asked me if I was messing around for fun (probably because I got so excited), but this was
actually useful for some MongoDB features that I was adding to VividCortex agents! I think the last time
I installed MongoDB was like 5 years ago when I was still in high school, but I was able to get a
sharded cluster up and running with Docker Compose with a single shell command. I saved a *lot* of
time.

---

Want to know what I worked on? Go ask my coworkers at Velocity or MongoDB World this week!

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Can&#39;t wait for <a href="https://twitter.com/hashtag/velocityconf?src=hash">#velocityconf</a> in San Jose this week! <a href="https://twitter.com/VividCortex">@Vividcortex</a> will be at Booth #724!</p>&mdash; VividCortex (@VividCortex) <a href="https://twitter.com/VividCortex/status/876502165258424322">June 18, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Can&#39;t wait for <a href="https://twitter.com/hashtag/MDW17?src=hash">#MDW17</a> in Chicago! Come say hi to the <a href="https://twitter.com/VividCortex">@Vividcortex</a> team at Booth #17!</p>&mdash; VividCortex (@VividCortex) <a href="https://twitter.com/VividCortex/status/876862412552626182">June 19, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>
