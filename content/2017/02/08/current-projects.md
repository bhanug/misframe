---
title: Current projects
date: "2017-02-08T23:45:00-05:00"
---

What am I working on right now? I have a bunch of stuff in progress that's keeping me very busy.

I'm still working on [lm2](https://github.com/Preetam/lm2), my generic key-value storage library
written in Go. There are bug fixes every now and then, but it seems stable for the most part.
That's the nice thing about keeping things simple. I could've spent the next decade working on a
great storage engine, but this slow, simple thing is enough for me!

I'm using lm2 as much as possible. Everything I've been working on lately uses it for storage.
The biggest of those projects is [Epsilon](https://epsilon.infinitynorm.com/), my events storage
service. It took a while to figure out the architecture and design, but now I think I've settled
on something really elegant. There will be more posts on that soon.

If I have an events service, I have to start sending events. The easiest way for me to send real
data to Epsilon was to send web traffic events. Specifically, the traffic to this blog! I blogged
about the [Alpha](https://misfra.me/2017/01/07/new-project-alpha-analytics/) analytics project
recently.

Here's what that looks like right now. All of the data is coming from an Epsilon node.

![Alpha screenshot](/img/2017/02/alpha-screenshot.png)

I've had some long nights trying to figure out responsive D3.js stuff with Mithril.js...

<blockquote class="twitter-tweet" data-conversation="none" data-lang="en"><p lang="en" dir="ltr">This took so freaking long BUT I GOT IT. <a href="https://t.co/Ja0jVSBMG9">pic.twitter.com/Ja0jVSBMG9</a></p>&mdash; Preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/824466275506200576">January 26, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Alpha is neat and useful, but it's not really the thing I really want to build. My goal is to use Epsilon
for personal finance and health info. Kind of like [Mint](https://www.mint.com/) or [Gyroscope](https://gyrosco.pe/),
but waaay more data-driven.

But fetching that information is really tedious. Most personal finance data like credit card transactions,
balances, etc. are available for export, but every company seems to have its own CSV format. Same thing
for personal health data.

Here's an example of a few bank account CSVs:

![CSV formats](/img/2017/02/csv-formats.png)

But that's not even all of the information I want to capture. I have a bunch of other info that isn't
automatically tracked like nutrition stats or exercise stats, which have to be tracked manually
*and regularly.* Real-world, human instrumentation is... complicated :). But I have an entire
lifetime to figure it out.
