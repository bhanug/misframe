---
title: Observium Annoys Me
date: "2014-12-30"
url: /observium-annoys-me
summary: I used to use Observium to monitor my switches. It was great, but now it is just buggy and bad.
---

I first started using Observium in 2011 or 2012. I was a senior in high school. I wasn't that good at programming. I mean, I could write code in a few languages, knew the basic data structures, Big-O, etc. but I was not familiar with many higher level concepts like monitoring. I knew about SNMP, but I didn't know anything at all about the implementation. As [Bitcable's](https://bitcable.com/) infrastructure grew to include network switches and more hardware, I needed a monitoring tool. I saw Nagios, Cacti, and others but they all intimidated me. I didn't have time to learn those tools. Things like college applications and math homework took up most of my time.

Observium was different. I was able to install it relatively quickly and had everything running without any issues. My daily routine when I got to school at 8 AM was to open up Observium in the syslab and get an overview of everything. It was enough.

I eventually had to reinstall Observium on another VM. This time, installation wasn't so smooth. I was able to add one device, but couldn't add another. Huh?

```
$ ./add_device.php <snip>.com ap v3 adminusr <snip> <snip> sha aes 161 udp
Try to add <snip>.com:
Trying v3 parameters observium/noAuthNoPriv ... 
No reply on credentials observium/noAuthNoPriv using v3.
Trying v3 parameters adminusr/authPriv ... 
Devices skipped: 1.
```

What went wrong? What's the error? Why did it get skipped?

I also found this goodie in the SNMP include file:

```
$ cat /opt/observium/includes/snmp.inc.php 
<?php

/**
 * Observium
 *
 *   This file is part of Observium.
 *
 * @package    observium
 * @subpackage snmp
 * @author     Adam Armstrong <adama@memetic.org>
 * @copyright  (C) 2006-2014 Adam Armstrong
 *
 */

## If anybody has again the idea to implement the PHP internal library calls,
## be aware that it was tried and banned by lead dev Adam
##
## TRUE STORY. THAT SHIT IS WHACK. -- adama.
```

It looks like Observium uses the Net-SNMP library with PHP. And Net-SNMP isn't exactly the greatest library out there...

<blockquote class="twitter-tweet" lang="en"><p><a href="https://twitter.com/PreetamJinka">@PreetamJinka</a> no… most languages have their own implementations b/c net-snmp is such a pile. (much of net-snmp is thread safe).</p>&mdash; Theo Schlossnagle (@postwait) <a href="https://twitter.com/postwait/status/545755326608580608">December 19, 2014</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I want to like Observium. I love the idea. This is what they write on their [home page](http://observium.org/):

> Network monitoring for all.
> 
> Observium is an autodiscovering network monitoring platform supporting a wide range of hardware platforms and operating systems including Cisco, Windows, Linux, HP, Juniper, Dell, FreeBSD, Brocade, Netscaler, NetApp and many more. Observium seeks to provide a powerful yet simple and intuitive interface to the health and status of your network.

That sounds great, but I think they messed up along the way. Observium is buggy. What I dislike even more is the fact that they charge for a "Professional" edition and release an open-source edition with limited features.

> The Open Source edition only receives critical security updates between 6-monthly release cycles and is best for small non-critical deployments, home use, evaluation or lab environments.

They don't say anything about new features. I'm interpreting this as, "the open source version will stay the same unless there are security issues, and anything new is something you have to pay for." That annoys me. A lot.

<blockquote class="twitter-tweet" lang="en"><p>This is a GigE interface. I shouldn&#39;t be seeing this &gt;.&lt;. <a href="http://t.co/K7fG3o9cBG">pic.twitter.com/K7fG3o9cBG</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/407931417134260224">December 3, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I don't like Observium anymore. It's clunky (have you seen the [dependencies](http://www.observium.org/wiki/Installation)?!), buggy, and I don't like what the authors are doing.

And that's why I'm working on [Cistern](http://preetamjinka.github.io/cistern/). I want to show you how *I* think network monitoring should be done. I think I'm on the right track because [many](https://cloudhelix.com/) [companies](http://www.arbornetworks.com/products/peakflow) [are](http://www.solarwinds.com/solutions/network-flow-analyzer.aspx) [doing](http://www.metaforsoftware.com/blog/netflow-traffic-analyzer-beyond-nbad) [similar](https://www.sevone.com/supported-technologies/network-performance-management) [things](http://www.ca.com/us/opscenter/ca-network-flow-analysis.aspx).
