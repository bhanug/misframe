---
title: Cistern Design Notes
date: "2017-07-15T20:30:00-04:00"
---

[Cistern](https://cistern.github.io/cistern/) is a project I started 3 years ago, back when I ran
a hosting business and wanted a simple tool to aggregate network flow datagrams from my switches.

![](/img/2017/07/cistern-plots.png)

Cistern hasn't been a priority for me for a while since I stopped running that business and
haven't touched a physical switch in a long time. At a certain point, I wanted it to support
more than just layer 2 and layer 3 network flow information, so I added support (via my
[appflow](https://github.com/Cistern/appflow) package) for generic HTTP application flows.

Development basically stopped at that point. There was a bunch of stuff I didn't like about the
implementation. I wrote a [custom time series storage engine](https://misfra.me/state-of-the-state-part-iii/)
for it, but it's hard to work with *just* metrics for flow data. I wanted raw events to group in
arbitrary ways. The internal architecture of Cistern also moved to a really complicated message
passing system with lots of channels, goroutines, and callbacks.

It's time for the third rewrite.

I don't have a detailed design since I'm just getting started with the rewrite, but here are my
high-level notes:

* Events data model
  * Group by and aggregate events in arbitrary ways
  * Automatic roll-up and dimension reduction
* "State sharing architecture" instead of internal message passing
  * Inspired by Ken Duda's talk about Arista's EOS architecture (on [YouTube](https://www.youtube.com/watch?v=Hfwr6sY27hA))
* Cloud native
  * Support for AWS VPC Flow Logs and CloudWatch Logs
  * Automatic backup and restore from S3
  * Maybe an AMI to deploy into a VPC
* Simple CLI tool to "query" a Cistern node
* (Eventually) Grafana support

So yeah, lots of neat stuff coming soon!

<img src='/img/2017/07/cistern.png' height=200/>

The goal is to keep things simple, developer-friendly, and be a great foundation to build on top of.
