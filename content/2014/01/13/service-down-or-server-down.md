---
title: Service down or server down?
date: "2014-01-13"
url: /service-down-or-server-down
---


How do you check if a server (an entire physical server or VM) is down? If it's running some sort of web server, I'll usually open up a browser and check if a page loads. This isn't very useful, since a web server process could be down but the server is still up. In that case, I try pinging it. If that doesn't work, I'll try SSHing into it. If none of those work, I am pretty sure that the entire server is down.

What if you wanted to do these checks from a Go program? As far as I know, ICMP echos (what ping uses) are not supported by the `net` package in Go. You could probably figure out a way to do it, but are there other methods?

![](https://31.media.tumblr.com/19cc34ca920a237601b6883d0593bbb5/tumblr_inline_mzdf7oJtBp1rs73cz.jpg)

(From http://www.cisco.com/web/about/ac123/ac147/archived_issues/ipj_9-4/syn_flooding_attacks.html)

I think I've figured out a neat trick. TCP uses a 3-way handshake. The "initiator," or client, starts off by sending a SYN packet to a destination with a specific port. If there's something listening on that port, the listener, or server, sends back a SYN-ACK. If there's nothing listening on that port, the listener responds with RST.

![](https://31.media.tumblr.com/df763c48a63863a92a4fea453dfc7295/tumblr_inline_mzdfmiygD51rs73cz.png)

Now, what if the server is down? You won't get anything back!

I tried out a neat way to use this. This is a really simple example in Go.

    conn, err := net.Dial("tcp", address)

If `err` is `nil`, everything's fine. If it's not `nil`, you could check for two cases. If `err` has `"connection refused"`, that means the server responded but nothing is listening on that port (there may be other reasons as well). If `err` has `"no route to host"`, then I'm *pretty* sure that server is down.

Of course, if the client / initiator has been blocked at a lower level of the stack, none of this is useful.

Useful? Did I get this completely wrong? Let me know on [Twitter](https://twitter.com/preetamjinka).

