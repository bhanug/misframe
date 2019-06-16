---
title: SNMP Part I
date: "2014-11-09"
url: /snmp-part-i
---


[SNMP](https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol) stands for Simple Network Management Protocol.
In case you haven't heard, it's not simple. SNMP is older than me, and it's used everywhere in networking.

I need SNMP support for [Cistern](https://github.com/Preetam/cistern). sFlow is great for statistics
but it's not useful to get general information. You can easily get interface statistics from sFlow datagrams,
but you may want to know what the interface description strings are. You may want to know which VLANs those interfaces
are assigned to. You may even want to get the description strings of VLANs:

![](/img/copied/posts/snmp-part-i/observium-vlans.png)

The idea of using flows to get this information is silly. The best way to do this is via polling, and doing so when you need to.
That's where SNMP comes in.

So, why write a decoder in Go? There's [Net-SNMP](https://net-snmp.sourceforge.net/), which is basically the defacto
library for this stuff. It has a C library that you can easily use from Go. The issue is that it's not thread safe.
That's a little annoying for Go programs. The other issue is that I dislike using cgo. A pure-Go implementation
is much cleaner, in my opinion.

Protocol summary
---
SNMP is not like sFlow at all. sFlow is unidirectional, i.e. packets go one way. SNMP has requests and responses.
To make things even more complicated, it's all over UDP. That means it's your responsibility to handle that state.
SNMP is also used for things like switched PDUs (power distribution units), so you can use SNMP to do remote reboots.
This is very sensitive stuff, so you really need your datagrams to be encryped. SNMPv3 supports encryption, which is
great but complicated!

SNMP uses a binary encoding format called [ASN.1](https://en.wikipedia.org/wiki/Abstract_Syntax_Notation_One). There's a
standard Go package called [encoding/asn1](https://golang.org/pkg/encoding/asn1/), but it's quite awkward to use. There's
some funk with struct tags and reflection. It seems that most people write their own ASN.1 encoding and decoding functions.

I spent most of my time reading specifications, diagrams, and RFCs to figure out just how to send valid datagrams.
Wireshark was incredibly useful in this case. Its interface is excellent. You can click on specific bytes in the hex and it'll
tell you which field they correspond to.

<blockquote class="twitter-tweet" lang="en"><p>Look, I&#39;m getting a response! <a href="https://t.co/yenOjrn2R1">pic.twitter.com/yenOjrn2R1</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/528057163806437376">October 31, 2014</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I grabbed a Brocade FWS 624 switch off Ebay. I'm using that for my SNMP tests at home. I think it's too risky to mess around
with production equipment that customers depend on (duh!). This is a layer 2 switch with SNMP and sFlow capability, so it's really helpful.
I really don't need anything with 24 ports, but it's hard to find cheap managed switches that support sFlow. I got really lucky with this find
on Ebay, especially since it's a Brocade device and those are the only ones I have experience with.

![](/img/copied/posts/snmp-part-i/fws-624.jpg)

<blockquote class="twitter-tweet" lang="en"><p>Hand-crafting SNMP packets right now :&#39;( <a href="https://t.co/faJNB5iH1Q">pic.twitter.com/faJNB5iH1Q</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/528811871567347713">November 2, 2014</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

A lot of this is quite tedious. Once I got used to hand-crafting packets and the encoding, I whipped up some simple
Go types that encoded themselves. Then I was able to write stuff like this:

```go
conn.WriteTo(Sequence{
	Int(3), // this is an INTEGER
	Sequence{
		Int(rand.Intn(1000000000)),
		Int(65507),
		String("\x04"), // this is an OCTET STRING
		Int(3),
	},
	String(Sequence{
		String(""),
		Int(0),
		Int(0),
		String(""),
		String(""),
		String(""),
	}.Encode()),
	Sequence{
		String(""),
		String(""),
		GetRequest{
			Int(rand.Intn(1000000000)),
			Int(0),
			Int(0),
			Sequence{},
		},
	},
}.Encode(), &net.UDPAddr{
	IP:   net.IPv4(10, 2, 33, 100),
	Port: 161,
})
```

The last major piece was encryption, which is a major PITA. You can't use Wireshark for this (duh!).
I used [github.com/tiebingzhang/WapSNMP](https://github.com/tiebingzhang/WapSNMP/) as a reference for the
encryption stuff. I thought about forking it but it doesn't have a license. It's also a fork, so I'm not
really sure how to deal with that. I also didn't think it was good, idiomatic Go code. :-/

After hours of work over a few weeks...

<blockquote class="twitter-tweet" lang="en"><p>IT WORKS! <a href="https://t.co/OXhIZX6uqA">pic.twitter.com/OXhIZX6uqA</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/531541736062214144">November 9, 2014</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

Not simple at all. At least now I have something that works and can iterate really quickly. And with that,
SNMP gets added to my LinkedIn profile :P.
