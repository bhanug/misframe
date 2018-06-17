---
title: Current projects
date: "2018-06-17T10:30:00-07:00"
---

Occasionally I like to list or describe what I'm working on
in my free time. Here are the past three:

* [Dec 2017](/2017/12/05/current-projects/)
* [Feb 2017](/2017/02/08/current-projects/)
* [Jun 2015](//2015/06/29/current-projects/)

### Changes since February

Here are things I mentioned last time:

* [**Transverse**](https://transverseapp.com/) is disabled for everyone except me. GDPR is a reason but not the only one. See below.
* **MySQL Explain Analyzer**: I don't use MySQL on a daily basis anymore (see [Hello, ShiftLeft](https://misfra.me/2018/03/22/hello-shiftleft/))
so I'm not really working on this at the moment. I already redesigned it, so that's good. I might extend it
for PostgreSQL explains since I now work with PostgreSQL on a daily basis.
* **Rating tracking app**: Starting to focus more on this now! I'm trying out lots of different places to eat
in the Bay Area and need something more than a spreadsheet to track them all!
* **Guitar app**: I'm not playing guitar as much since I moved but I'll fix that.

### Current focus

This is what I'm focusing on right now:

* **New time series / events storage idea:** I have an idea for a new cloud-native storage engine for events.
The cool part will be automatic indexing using statistical learning techniques. I'm really excited about
this because it'll be a combination of all of my favorite things: statistical learning, time series,
and databases.<br><br>I also want to use this project to learn Rust and Pony. Maybe to stay productive
I'll stick to Go for the main part and use the others to write some tooling.
* **Transverse** needs a refactor to encrypt everything. I didn't clean up the code
when I redesigned it so I need to do that now. I also found a forecasting bug recently.
* **Rating tracking app**: I'm working on the Sketch designs.
* **Guitar app**: I need to start working on the Sketch designs.

### Things I played around with

These are things I started at some point but am not focusing on anymore:

* [SQL-like query execution package in Go](https://github.com/Preetam/query): I pulled out Cistern's
query execution code into a separate package. The idea for this is to allow the user to implement
a few interfaces (e.g. a table interface that lets you iterate through rows) and the package will
execute a SQL-like query against that interface. This is a big project! It's really hard to figure
out what the best interface abstractions are, especially since I eventually want push-down filtering
and aggregations.
* lm2
  * Hole punching: lm2 is append-only and never reclaims space so I wanted to play around with
  file hole punching to keep the append-only design. This can be done but the "garbage collection"
  process would be really complicated since I would have to start moving records in order to punch
  some pages out.
  * lm3: I also wanted to put lm2 on S3. lm3 was going to be lm2 that thinks it was talking to a local
  file but in reality was a bunch of blocks in an S3 bucket. At a certain point this also got complicated
  and wasn't worth working on for now.

### Non-programming stuff

* I started working on a personal finance site about a year ago to write down stuff I learned or
figured out but I didn't get too far. I renewed the domain name recently so I guess I still have the
option to keep it going.
