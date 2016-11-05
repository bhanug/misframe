---
title: "Epsilon: An Events Database"
date: "2016-11-05T12:30:00-04:00"
---

I'm working on an events database. I work for a monitoring company, so I know about the value
monitoring can provide. We monitor as much as we can about databases.

<div class='bigquote'>What gets measured gets managed.<br>&mdash;Peter Drucker</div>

There's a bunch of stuff in my personal life that's tracked / monitored already.

* Health (weight, step count, sleep information)
* Ratings (movies, TV shows, food)
* Financial (transactions, balances, credit scores)

All of these things can be treated as events. Unfortunately, all of these events are stored in
totally separate places. Health data is in Apple's HealthKit or FitBit, my ratings are basically
only on Twitter or Facebook as status updates, and financial information is at my banks. I want
everything in one place. I'm building a database for it, and (for now) it's called *Epsilon*.

Epsilon isn't anything impressive. I want it to be enough to make a few dashboards, but I also want
it to be a playground for experimentation.

Eventually, I'd like to have a system with

* Leader election and broadcasts using [libab](https://github.com/Preetam/libab)
* Storage backed by [lm2](https://github.com/Preetam/lm2)

The storage part is already working (I've had tons of practice 😛). I'll dig into replication in the
next few weeks.
