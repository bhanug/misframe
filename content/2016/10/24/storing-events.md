---
title: Thoughts on storing events in a key-value store
date: "2016-10-24T23:00:00-04:00"
---

I'm implementing an event storage system on top of lm2, an ordered key-value store. The events I'd
I'd like to store are just flat maps (i.e. a map of JSON strings, numbers, and bools). The important
question right now is, what should the key-value layout be?

With an RDBMS, one of the first things I think about is the primary key. That determines the primary
order of data and establishes a constraint for uniqueness.

That's easy enough, in my case. I'm going to start with this:

[<img src='/img/2016/10/event_primary_key.jpg' width='50%'/>](/img/2016/10/event_primary_key.jpg)

Events are ordered by time. I threw in an additional component to be able to have more than one
event per Unix timestamp. It could also be a sub-second component of the timestamp.

At the moment, there are two ways I can think of to store values.

[<img src='/img/2016/10/event_key_value_layout.jpg' width='75%'/>](/img/2016/10/event_key_value_layout.jpg)

The first has one key-value pair for each event. It's simple. However, updating an event requires
the entire value to be changed. And because lm2 is purely append-only, updating events also leads to
dead keys.

The second approach is better for updating events, since each field could be updated individually,
but it has a bunch of overhead. The number of total key-value pairs will be much higher, and it
requires the primary key to be stored several times for each event.

I'm probably going to go with the first format for now. It's easier to implement :).
