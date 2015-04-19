title: Cistern: The Vision
date: 2015-04-19 02:00:00
url: cistern-the-vision
summary: Reinventing network monitoring

Background
---
As a hosting provider, I've had my fair share of DDoS attacks. My company doesn't do any peering with transit providers. We just have a single upstream provider at our Ashburn data center. My provider has an automated DDoS detection system, which is made from scratch, that detects anomalous flows and either automatically blocks traffic or sends email alerts. I sometimes get alerts forwarded to me in case it's an outbound anomaly originating from one of my clients' VMs. I asked my provider for details and he was generous enough to share them, and I thought it was extremely fascinating. It isn't very complicated and seems to work well. I'd rather not share details here, but I can get you in touch for more details.

---

A little over a year ago, one of my customers submitted a support ticket because his site was on the front page of Hacker News and lots of visitors were seeing CloudFlare errors because the origin server (hosted by me) wasn't responding. I SSH'd into the cPanel server to see what was going on. Everything looked fine. I probably tweaked settings for 30 minutes or so until I gave up. I still didn't know what was wrong.

I thought it may have been a network issue. I tried to SSH into another server. Connection timeout. Huh? I tried again. This time, it worked. I disconnected and tried repeatedly and it seemed like four in five attempts to connect would fail. I tried different servers. Same issue. It *has* to be a switch issue, I thought. These are different physical machines, and the only thing they have in common is the switch. I've had switch problems before, and they're extremely annoying to diagnose. This could be bad, I thought, especially since I didn't have a spare switch. It's worth mentioning that I was getting Panopta alerts for all servers during this time. *Preetam was probably in panic mode.*

I asked my provider if they knew why my TCP connection establishment was so poor. We couldn't figure it out. I was told that there was a big fiber cut in the DC area, and was sent the following message:

> Welcome to the Cogent Communications status page. Customers located in the Washington DC area may be experiencing latency and/or packet loss. This is being caused by a fiber cut. Our fiber vendor is aware of the issue and working to repair the damage. There is currently no ETR. The master case is HD5615392.

The problem wasn't solved even after the fiber cut issue was resolved. After some more communication with my provider, I was eventually told that there was a 32k/sec SYN rate limit placed on our network. Hitting that limit would explain why my connections were so horrible! But 32k/sec SYNs is *a lot* of SYN traffic. Something weird is going on.

I wish I remembered how I solved the problem. Turns out that one of my clients' VM was sending an outbound attack with a high SYN rate, so that's why we were hitting that rate. I spent *hours* trying to figure out what was going on. I just wanted to figure out why my TCP traffic was doing so poorly. I should have asked a better question: "what is my network doing?"

---

Given the right tools, those problems could have been solved in seconds or minutes. Really. I am trying to build one, and it's called Cistern.

Cistern
---
[Cistern](http://preetamjinka.github.io/cistern/) is a flow collector. I wrote about it previously [here](http://misfra.me/state-of-the-state-part-ii). Its main purpose is to serve as the destination for flow data, in the form of packet samples and counters. It will aggregate these data, analyze them, and serve as a platform to build richer systems. I was certainly surprised to learn how much information can be extracted from packet samples. Those things are quite dense, and they provide a level of insight that you can't get from any other method.

Flows are an efficient, scalable method of collecting information. Cistern currently decodes flows using the [sFlow](http://sflow.org/) protocol. sFlow does have its limitations. It isn't very useful on its own, and in fact, it wasn't designed to be. SNMP polling is used to fetch metadata, like interface names, when necessary. sFlow and SNMP together maximize monitoring capability.

This is a short summary of what I want to see in Cistern in the near future:

* State of the art software engineering and analytics
    * Device autodiscovery
    * Automated threat detection
    * Efficient statistical analysis

* High quality, language-agnostic APIs
	* RESTful JSON

* Open-source core with plenty of features to start with
	* Anomalous flow detection
	* Flood detection
	* IP spoof detection

This whole project is like a big puzzle. I have a blurry vision of what it'll end up like, but I'm basically starting from scratch. I need to figure out how to break it down into individual pieces, implement those pieces, and then figure out how to put everything together again.

Here are some of the pieces that I have so far:

* Cistern
	* https://github.com/PreetamJinka/cistern
* Fully async, thread-safe SNMP v3 implementation. It uses a single socket for all SNMP traffic, so there is plenty of scalability
	* https://github.com/PreetamJinka/snmp
* Compressed time series storage engine written from scratch
	* https://github.com/PreetamJinka/catena
* The only open-source (as far as I know) sFlow implementation in Go
	* https://github.com/PreetamJinka/sflow
* OSI layer 2, 3, 4 protocol decoding
	* https://github.com/PreetamJinka/proto
* AngularJS powered web UI
	* https://github.com/PreetamJinka/cistern/tree/gh-pages/ui

I think it's a bit inaccurate to describe this as a monitoring system. I like the following comment someone made on Hacker News about Observium (something I [don't like](http://misfra.me/observium-annoys-me)):

> (In response to "Observium: An auto-discovering network monitoring platform")

> Another system focused on the wrong thing in monitoring: on alerts and charts. Those are merely methods of consuming data, not the only ones and not even the most important ones a decent monitoring system should do.

> Sending e-mail or displaying a set of charts or a status table is simple. Allowing to collect, collate and aggregate the data (metrics and events) in arbitrary way, also as an afterthought, is what monitoring system should do. With virtually everything on the market, when a need for any processing not anticipated by monitoring system author arises, one needs to write much stuff outside said system.

> We need less systems resembling invoicing systems and more systems resembling general purpose databases.

> This is why monitoring *still* sucks.

> -- https://news.ycombinator.com/item?id=9248672

I think it's very important for Cistern to behave like a database of network data. We shouldn't just want charts and thresholds! Imagine having access to top flows, top talkers, aggregations and time series of protocol metrics, hardware diagnostics, and more, all from a web UI, JSON API, or query language. Cistern should be a tool and not just a resource.

Business model
---
This is my startup idea. I think there is business potential here, but that is not my primary goal. I think it's important for everyone to have access to tools that provide more insight into networks. That's why Cistern will always have a free and open core. I want it to be accessible so even 16-year-olds can use it, dig around the source code, make changes, and learn. I think I would have wanted something like that when I was younger. More realistically though, I think the small-scale, hosting providers would need this more than anyone. They probably don't make enough to spend thousands on monitoring software, and they probably don't have the resources to administer their network as much as it needs.

I think there is a lot of opportunity for revenue with support and custom integrations. I think this makes a lot of sense for an open-source project. Custom additions to fit into a specific environment will require development regardless of what tools you choose, so maybe it would be better if the original developers do it? I can already imagine people building plugins to inject rules or configuration updates into OpenBSD firewalls, Cisco, Brocade, and Juniper routers and switches, and so on.

Requests
---
I don't think Cistern is ready for others to use yet, but I'd like to get some alpha testers at some point. I'm mainly building it for myself at this point, but it would be useful to hear what others want to see.

I'm also looking for donations! I'd like to add support for Cisco NetFlow but I don't have any Cisco gear. If you are, or someone you know is, willing to donate Cisco hardware that supports NetFlow (or just SNMP), let me know.

Why don't you just use ____?
---
Building from scratch is a great way to learn.
