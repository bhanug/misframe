---
title: ARP 101.
date: "2014-06-10"
url: /arp-101
---


ARP
---
ARP stands for Address Resolution Protocol. It's a protocol to translate network layer
(layer 3) IP addresses to link layer (layer 2) MAC addresses. Why would you need this?
Let's look at a simple example.

![](/img/copied/posts/arp-101/hosts.png)

Here we have a couple of hosts. We need them to communicate with each other and they
both have IP addresses on the same subnet. They're connected to a layer 2 switch. We
want to have communication between these two servers over TCP, which is a layer above
IP. Remember, layer 2 devices don't know what IP addresses are! We can't just say,
"hey switch, send these packets to 10.0.0.5." The switch doesn't know IP addresses.
The switch *does* speak Ethernet. Our servers are connected to the switch via an
Ethernet connection, so they'll communicate using Ethernet packets.

![Ethernet packet structure](/img/copied/posts/arp-101/ethernet_packet_format.png)

Look, there are "source" and "destination" fields. We're going to use those. Whenever
a host sends an Ethernet packet, it fills those fields in. When the switch receives
the packet, it reads the destination field and sends the packet to whichever physical
port it needs to go to. It also looks at the source field to remember that the port
that it got the packet from has a host with that MAC address.

Every switch has a MAC table. This table shows which provides a mapping from MAC
addresses to physical ports.

![](/img/copied/posts/arp-101/mac_table.png)

The switch adds a MAC entry to its internal table when it recognizes a new address
coming from one of its ports. But what if you don't know the destination's MAC address?
There is a special MAC address, FF:FF:FF:FF:FF:FF, which is a *broadcast* address. It's just
like a broadcast IP address in a subnet. Any Ethernet packets sent to the broadcast
address will be transmitted to every port on the switch. When a host receives a broadcast
packet, it'll know that it's a broadcast packet because of the destination address.
It will also know the MAC address that it came from, because of the source address.

![](/img/copied/posts/arp-101/broadcast.png)

In the case of our two hosts, they don't know each other's MAC addresses. But that
does not mean that they are not able to communicate -- they can send broadcast
packets! However, this isn't ideal because they're not just sending packets to each
other, but rather everyone on the network. It's like trying to have a conversation
in the middle of a large crowd. It doesn't work so well. This is where ARP comes in.

Using ARP, we can send a question to FF:FF:FF:FF:FF:FF and it'll be received by all
the hosts on the network: "Who has IP address 10.0.0.5? Tell me, 10.0.0.2." Now, if
someone has IP address 10.0.0.5, they will respond by sending a packet to the MAC
addresse listed as the source of the original request. In addition, they now know
the MAC-IP pair of the requester! This means they won't have to send another ARP
"who has" request later on.

At this point, the two hosts know the other's MAC address and IP addresses. They can
communicate with each other directly!

Now, can you think of ways exploit this? Check out broadcast storms and ARP spoofing.
