---
title: Replication with Epsilon... or Not
date: "2016-11-20T13:00:00-05:00"
---

I'm working on an events database called [Epsilon](/2016/11/05/epsilon/). I have a bunch of things
that I'd like to build on top of Epsilon, and I really want to be a solid data platform. That means
one of the key requirements is replication. Why? I think replication has two benefits:

1. **Availability:** If one instance goes down, another can handle requests.
2. **Safety:** If one instance goes down without recovery, another has a copy of the data.

Over the past few days, I've been working on implementing basic replication into Epsilon.
You can kind of tell by the recent releases in [libab](https://github.com/Preetam/libab/releases)
=). I started with the usual database replication route. All writes first go to a "master" and then
get replicated to the other servers in the cluster. But this is hard work. And it can get *really*
complicated.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Replication sucks.</p>&mdash; Preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/793296523731824641">November 1, 2016</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Now I'm taking things in a different direction. I still want to get those two benefits I mentioned
above, but without complicated replication logic / coordination at the storage layer. This new
approach pushes replication logic out to clients, and each Epsilon instance will be completely
standalone. It's not really "replication" in the traditional database sense.

Here's how I'm thinking about implementing them.

### Availability

This one starts off simple. I'll just run multiple Epsilon instances =). Without adding anything
else in the back-end, that means clients have to push data to multiple places. This creates other
problems like inconsistency. What if one instance fails and another is OK? Again, it's up to the
client to take care of that stuff, but I think it's easier to handle those situations at a higher
level. I'm working on spec'ing out some neat features into Epsilon to make that possible and
(relatively) simple for applications.

### Safety

The entire Epsilon infrastructure will be built on DigitalOcean, which has point-in-time snapshots.
Assuming snapshots are durable, I only have to worry about new data that has not been included in a
snapshot. The nice thing about building an events database is that these events are (at least in my
use cases) usually somewhere else to begin with. If I have to revert back to a snapshot, I can just
resend the missing events.

---

I haven't gone into detail about this since I'm still at the planning stage. So far, it seems like
it can be an awesome infrastructure, especially for one person.
