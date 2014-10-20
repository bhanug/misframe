title: Custom Router
date: 2014-10-19 22:05:00
url: custom-router

So this clever new idea has been floating around in my head
for a couple of days now. If you consider a router, it's basically
connecting directing packets to different subnets. The simple case
is when you have two subnets and a router that's in the middle.

Let's call the router R1, with the two subnets being S1 and S2.
Let's say you're managing a bunch of servers in a data center rack
and they're all on S2. Your router, R1, connects them to your
bandwidth provider who is on S1.

R1 has two interfaces -- one for S1 and one for S2 -- and it
has an IP address on each. Your provider is routing all of the
traffic going to S2 to R1's address on S1. All of your hosts
on S2 are sending the packets that need to go outside the subnet
to R1's address on S2.

This is really simple (conceptually, of course). I think you can
write something that can do this with a couple of raw sockets.
Well, I think I can actually write something that can do
simple routing like this. All you have to do is peek at the
IP headers and rewrap the IP packets with new Ethernet headers.

Argh, I don't think I have the hardware to test this out easily.
Maybe I can play around with some VMs? I haven't tried that out
yet.

For some reason I keep thinking that the kernel doesn't like this.
Perhaps `iptables` will mess with stuff or something. Anyway,
it'd be *really* cool to write a firewall in Go that does some
neat deep packet inspection.

---

Increasing my knowledge of the networking stack :D
