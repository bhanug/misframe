title: Gambling with cheap RAM
date: 2013-12-30 17:04:15
url: gambling-with-cheap-ram

I got an email from [123Systems](http://123systems.net/) about a week ago:

![](https://31.media.tumblr.com/678055852db80bb4020e687d71b2cd1a/tumblr_inline_myn1ycfmg41rs73cz.png)

First off, I got that email because I used to be a customer of 123Systems. I would *not* recommend them if you're looking for a reliable VPS host. I've had issues with them and their support isn't the best.

They are, however, really cheap. Read that promotion carefully -- a 2 GB RAM VM for $25 per *year*. That's a little more than $2 a month. 3 TB of bandwidth is also a generous amount. They're clearly overselling, but who doesn't?

With that offer, 123Systems is an order of magnitude cheaper than Digital Ocean.

I'm a mathematics major, so I (should) like math problems. Here's an interesting one...

## Background

123Systems states that they have a "99.9% uptime guarantee." Here's what that basically means to most people, if you're not sure: your service will be up 99.9% of the time (0.1% of one month is 43.8 minutes), or we'll give you a service credit.

So, it's not really *guaranteed* to be up 99.9% of the time. Let's be pessimistic and say that a server will be down for a maximum of 15 hours total.

## The problem(s)

Suppose you're a data gambler. You're willing to lose all of your data, but you'd rather not. You want to store all of your data in 123Systems VMs, and only in memory. You'll have replica VMs, so you'll only lose your data if *all* replicas are down. At any given hour, a VM is either up or down for that entire hour. All VMs are down for 15 hours total each month (720 hours), and their downtime is distributed randomly. When a VM is back up, it can rereplicate data back.

If you had N replicas, what's the probability that your data will be safe after a month (720 hours)? In other words, what's the probability that at least one VM was up for each hour in a month?

How many replicas would you need to ensure that the probability of not losing your data is 99.99%? How much would that cost (assuming $2 a VM)?

-----

If you have an answer, or if there's not enough information to solve this, let me know on [Twitter](https://twitter.com/preetamjinka).

