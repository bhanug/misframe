---
title: Client-side DNS Failover
date: "2017-03-04T16:30:00-05:00"
---

Client-side DNS failover is one of the coolest things I learned about recently.
It's simple. Instead of having a single A record with one IP, you add additional
A records with different IPs, and clients automatically skip over failing servers.

Check out this Webmasters Stack Exchange question:
["Using multiple A-records for my domain - do web browsers ever try more than one?"](http://webmasters.stackexchange.com/questions/10927/using-multiple-a-records-for-my-domain-do-web-browsers-ever-try-more-than-one)

> Pretty much every browser does indeed receive the full list of A records, and does indeed check others if the one it is using fails. You can expect each client to have a 30 second wait when they first try to access a site when a server is down, until it connects to a working address. The browser will then cache which address is working and continue using that one for future requests unless it also fails, then it will have to search through the list again. So 30 second wait on first request, fine thereafter.

For example, let's look at `amazon.com`:

```
$ dig amazon.com
;; QUESTION SECTION:
;amazon.com.			IN	A

;; ANSWER SECTION:
amazon.com.		55	IN	A	54.239.25.192
amazon.com.		55	IN	A	54.239.25.200
amazon.com.		55	IN	A	54.239.17.7
amazon.com.		55	IN	A	54.239.17.6
amazon.com.		55	IN	A	54.239.25.208
amazon.com.		55	IN	A	54.239.26.128
```

If 54.239.25.192 is down, my browser will automatically try 54.239.25.200, and so on.

It's not just web browsers that can do this. Programs like curl and telnet do it too.

```
$ telnet amazon.com
Trying 54.239.26.128...
telnet: connect to address 54.239.26.128: Connection refused
Trying 54.239.25.200...
telnet: connect to address 54.239.25.200: Connection refused
Trying 54.239.25.208...
telnet: connect to address 54.239.25.208: Connection refused
Trying 54.239.17.7...
telnet: connect to address 54.239.17.7: Connection refused
Trying 54.239.17.6...
telnet: connect to address 54.239.17.6: Connection refused
Trying 54.239.25.192...
telnet: connect to address 54.239.25.192: Connection refused
telnet: Unable to connect to remote host
```

If you're writing your own programs, you can think about using this technique to implement
high availability. Go programmers have it easy, as the `net` [package](https://golang.org/pkg/net/#Dialer)
does it automatically *and* it also supports dual-stack fallback. I'm not sure about other
languages<sup>1</sup>.

I can see this being useful for analytics services. Browsers will automatically failover to active
servers, so you don't have to do anything fancy at the back-end. Last year, I [posted](https://misfra.me/2016/04/29/ha-with-libab-and-digital-ocean/) an HA example
with a Digital Ocean Floating IP and leader election with my libab library. That approach makes
failover transparent to the client, but I think this approach is more elegant. Clients are smart!

---

1. I tried to see what Node does and found [this](https://github.com/nodejs/node/issues/708) issue.
Looks like it's not supported and probably won't be supported in core.
