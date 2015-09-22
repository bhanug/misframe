---
title: Custom Router Part II
date: "2014-10-22"
url: /custom-router-part-ii
---


Welcome back! In my [previous post](http://misfra.me/custom-router) I
described this interesting idea of writing a router.
I had no idea whether or not it would work. I knew it was possible, of course. I
run two OpenBSD routers in a failover setup with CARP (this blog is routed through
them, FYI).

Setup
---
Turns out my BeagleBone Black is great for testing this out! When you plug in
a factory default BeagleBone Black into a computer, it sets up a network over
USB.

![](http://static.misfra.me/images/posts/custom-router-part-ii/bbb-network.png)

	eth1      Link encap:Ethernet  HWaddr 78:a5:04:c8:8c:a3  
	          inet addr:192.168.7.1  Bcast:192.168.7.3  Mask:255.255.255.252

It shows up on my laptop as eth1. My laptop is assigned the address 192.168.7.1 and
the board has 192.168.7.2. sshd is running on the board, so I can easily SSH in via
192.168.7.2:22.

The board has no other connections. There's no WiFi, and I don't have it connected
over Ethernet. It cannot send packets out to the Internet. What it *can* do is
send packets to my laptop, and my laptop *is* connected to the Internet. The Internet
is just a series of tubes right?

I'll just summarize the steps I took to getting this board to reach the rest
of the Internet.

Default gateway
---
First, we need to set my laptop as the board's default gateway. Otherwise,
it won't know where to send packets outside the subnet!

	root@beaglebone:~# ping 199.58.162.130
	connect: Network is unreachable

It's a simple one-liner:

	root@beaglebone:~# ip route add default via 192.168.7.1
	root@beaglebone:~# route -n
	Kernel IP routing table
	Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
	0.0.0.0         192.168.7.1     0.0.0.0         UG    0      0        0 usb0
	192.168.7.0     0.0.0.0         255.255.255.252 U     0      0        0 usb0

So, does it work? Let me run tcpdump on my laptop and run ping again on the board...

	  ∂ [21:06:42] [~]: sudo tcpdump -i eth1 icmp
	tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
	listening on eth1, link-type EN10MB (Ethernet), capture size 65535 bytes
	21:06:49.397470 IP 192.168.7.2 > misfra.me: ICMP echo request, id 1451, seq 1, length 64
	21:06:50.406267 IP 192.168.7.2 > misfra.me: ICMP echo request, id 1451, seq 2, length 64
	21:06:51.405815 IP 192.168.7.2 > misfra.me: ICMP echo request, id 1451, seq 3, length 64

Great! Packets are reaching my laptop... but they're not going anywhere after that. They're
simply dropped.

Packet sniffing
---
I gave a talk at beCamp 2014 on [packet sniffing](https://github.com/Preetam/packet-sniffing).
There are a few examples that I reuse over and over simply because they're great templates for me.

tcpdump is obviously seeing the packets we want, so we can too. Raw socket it up.

```go
fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, htons(syscall.ETH_P_ALL))
if err != nil {
	log.Fatal(err)
}

log.Println("Listening on a raw socket...")

. . .

n, _, err := syscall.Recvfrom(fd, buf, 0)
if err != nil {
	log.Fatal(err)
}

. . .
```

Decoding packets
---
What you get from reading from AF_PACKET + SOCK_RAW packets are Ethernet frames.
You need to decode these. I use my [proto](https://github.com/Preetam/proto/blob/master/ethernet.go)
package for that.

Routing itself and packet injection
---
You just got a packet, and you have to route it somewhere else. Where (and how) do you send it?
Well...  the simple answer would be to send it to your default gateway. There's obviously more to it,
but you can figure that out on your own (I did :P).

How do you actually send that packet to the default gateway? We know that the default gateway is on the
same subnet you are. Therefore, we're only working at layer 2 of the [OSI model](http://en.wikipedia.org/wiki/OSI_model).
There's a really simple answer for this one: just modify the MAC addresses in the Ethernet packet header
and write the packet back into the socket. That's it! I think this is called [packet injection](http://en.wikipedia.org/wiki/Packet_injection).
The Wikipedia page makes it sound evil...

Getting packets back
---
One thing you have to be careful about is getting packets back from the Internet.
I set up a static route on my WiFi router to route the 192.168.7.0/30 subnet
to my laptop.

![](http://static.misfra.me/images/posts/custom-router-part-ii/static-route.png)

Yes, the subnet mask is incorrect but it doesn't make a difference in this situation.

End result
---

Ta-daaaah!

![](http://static.misfra.me/images/posts/custom-router-part-ii/demo.gif)

On the left is the log output of a Go program that's reading and writing from/to
a raw socket, and printing out the Ethernet frames it's receiving and sending.
On the right is an SSH session on the board while I run `apt-get update`.

My Go program does the first routing to the Internet (and last in the inbound
direction).

	2014/10/22 21:39:30 Listening on a raw socket...
	2014/10/22 21:39:30 <nil>
	2014/10/22 21:39:33 this one needs to go to the gateway
	2014/10/22 21:39:33 {78:a5:04:c8:8c:a3 26:06:05:5f:40:f4 0 2048 <snip>
	2014/10/22 21:39:33 {e8:de:27:bb:6b:aa 9c:4e:36:59:b2:54 0 2048 <snip>
	2014/10/22 21:39:33 <nil>
	2014/10/22 21:39:33 this one needs to go to the BeagleBone Black
	2014/10/22 21:39:33 <nil>
	2014/10/22 21:39:34 this one needs to go to the gateway

Traceroute
---
Check this out... it's cool.

	root@beaglebone:~# traceroute misfra.me
	traceroute to misfra.me (199.58.162.130), 30 hops max, 60 byte packets
	 1  192.168.0.1 (192.168.0.1)  4.363 ms  6.072 ms  13.493 ms
	 2  10.2.33.1 (10.2.33.1)  16.266 ms  16.125 ms  16.266 ms
	 3  10.1.10.1 (10.1.10.1)  15.894 ms  15.743 ms  15.598 ms

How come my laptop (192.168.7.1) isn't showing up? Let's think about
how traceroute(1) actually works. In short, it sends out multiple packets
with different TTLs (time to live). Quoting [Wikipedia](http://en.wikipedia.org/wiki/Time_to_live)...

> The TTL field is set by the sender of the datagram, and reduced by every router on the route to its destination. If the TTL field reaches zero before the datagram arrives at its destination, then the datagram is discarded and an ICMP error datagram (11 - Time Exceeded) is sent back to the sender.

The reason why we don't see 192.168.7.1 is because my Go program does not
decrease the TTL (I'll explain why later). My Go program also does not send ICMP
datagrams. Think about this for a second. Isn't my custom router invisible? Well,
it's not since it's the default gateway, but what if it was a few hops down? Isn't
that scary? Uhm...

Checksums and TTLs
---
If you open up a diagram of an Ethernet frame (look [here](http://en.wikipedia.org/wiki/Ethernet_frame#Structure)),
you'll notice that there's a field called "Frame check sequence". This is a 32-bit CRC -- a checksum.
Turns out that NICs take care of calculating this checksum for you, so you don't have to worry about it when
constructing Ethernet frames.

IPv4 packets are different. You have to make sure the checksum stays consistent. The reason why
I'm not decreasing the TTL is because I'd have to recalculate the IPv4 packet checksum, and
I simply didn't have time for that (I had to eat dinner :D).

Conclusion and next steps
---
This was a pretty cool thing to write before dinner. It's not that long, either. All of the code
is available on GitHub and is MIT licensed ('cause I like your freedom):

https://github.com/Preetam/gateway-experiment

Now I'm thinking about inter-VLAN routing, stateful firewalls, routing tables, etc.
It would be neat to try to implement some of this stuff in Go. I already wrote a [post](http://misfra.me/router-on-a-stick) on inter-VLAN routing. That was over two years ago? I keep redoing stuff, but every time I do it I go a level lower :P.

Well, I hope that was informative. Ask me questions on Twitter: [@PreetamJinka](https://twitter.com/PreetamJinka)
