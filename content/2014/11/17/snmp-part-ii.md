---
title: SNMP Part II
date: "2014-11-17"
url: /snmp-part-ii
---


This post is more about the details of how SNMP is currently implemented in Cistern.
The code at the latest commit as I write this is [here](https://github.com/PreetamJinka/cistern/tree/4f57ab68c9a18266908a7221823b24085bd39d1c/net/snmp).

SNMP communication itself is pretty simple. There are requests and responses.

![](http://static.misfra.me/images/posts/snmp-part-ii/get-request-response.jpg)

*PDU* stands for "protocol data unit." I think of them as structs. If you look
in the RFCs you'll see them defined like this, in the ASN.1 language:

```
GetResponse-PDU ::=
  [2]
      IMPLICIT SEQUENCE {
          request-id
              RequestID,

          error-status
              ErrorStatus,

          error-index
              ErrorIndex,

          variable-bindings
              VarBindList
      }
```

That's a pretty dense way of describing a `GetResponse`. You can just think
of the `GetRequest` and `GetResponse` PDUs as structs with certain fields.

![](http://static.misfra.me/images/posts/snmp-part-ii/get-request-response-pdus.jpg)

You'll notice that they have the same structure. The only thing that changes between the two
is the PDU type identifier which goes in the header and the `VarBindList`. The `RequestID`
stays the same. This is how you know which request you got a response to, and it's very important.

---

If you run Wireshark (or tcpdump) as you run `snmpget` (or Cistern :P) over SNMPv3 with encryption,
you'll see something like this:

![](http://static.misfra.me/images/posts/snmp-part-ii/wireshark-screenshot.png)

There are four packets of communication. The last two are encrypted. You see the first two show up
because they are required to fetch the necessary parameters to do encryption.

![](http://static.misfra.me/images/posts/snmp-part-ii/discovery.jpg)

The `GetRequest` is blank. The response is a `Report`, a PDU similar to a `GetResponse`. It's
returning a COUNTER of how many invalid SNMP packets were dropped. The important part of this
response is that we get the following parameters necessary for encryption: `EngineID`, `EngineTime`,
and `EngineBoots`. Once we have those parameters, we can start encrypting the rest of our packets.

---

This is what SNMP communication may potentially look like:

![](http://static.misfra.me/images/posts/snmp-part-ii/communication-diagram.jpg)

Things are out of order, lots of things are going on at the same time, and it's all happening
over a single port. I suppose you *could* open a new socket for each request, but that's not
scalable at all! So this is our situation: a single socket with request-response sequences
happening concurrently, and it has to be really fast and efficient.

Concurrency for the win
---

It turns out that it's quite simple to structure a program using goroutines and channels
to handle this scenario. We basically have to send requests over the socket, receive responses,
and send the corresponding response to whatever made that request. Here's how Cistern does it
at the moment:

1. Create a map from `RequestID`s to channels.
2. Start a goroutine that listens on the socket, decodes the `RequestID` from response PDUs, and sends
to a corresponding channel in the map.
3. If a channel is found, send the response data to it.
4. Delete the channel from the map.

That goroutine keeps running for the length of the session. Nothing else is reading from that socket.

Then, when a request is to be made:

1. Create a request packet, and remember the `RequestID`.
2. Create a channel. Let's call it C.
3. Send the packet over the socket, and set C in the map.
4. Start another goroutine to do a "timeout." Basically, if the channel still exists in the map
after a certain period of time, *close the channel* and delete the entry in the map
5. This is the cool part. Attempt to read from the channel like so: `decoded, ok = <-C`. If C is
closed, `ok` will be false. Otherwise, we should see something for `decoded`.
6. If the request timed out, try again (only a finite number of times).

So it turns out this approach works really well. You can start up requests concurrently and
stuff doesn't blow up. Concurrency (and Go) for the win.

Single socket?
---

The current implementation uses a separate socket for each device that Cistern connects to.
If you had 1,000 devices, you'll have to open 1,000 sockets. That doesn't seem efficient.
I think everything can be done using a single socket. Specifically, [`UDPConn.ReadFromUDP`](http://golang.org/pkg/net/#UDPConn.ReadFromUDP)
needs to be used. Basically, it allows you to read from a socket and know where it came from.

The "session" logic in Cistern will get more complicated. I'm honestly not sure what the performance
would look like with one socket vs many. I don't even have many SNMP devices to test against.

Conclusion
---

The single socket approach will probably be coming by the end of the year or sometime
next year.

I really want this to be the world's most concurrent + efficient SNMPv3 implementation. It'll take
a lot of work to get there, though. I really would like an SNMP simulator. It is definitely possible.
I just need the time to do it...

 I think this will crush Net-SNMP in terms of thread-safety :P.

---

How the heck do I add this kind of stuff to my résumé / LinkedIn? Listing "SNMP" doesn't really capture
all of it!
