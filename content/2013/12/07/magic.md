---
title: Magic.
date: "2013-12-07"
url: /magic
---


A friend was talking over dinner about when he was asked, during an interview, what happens when you type in a URL into an address bar in a browser. Lots of stuff happens!

I'm just going to start listing stuff. It's practically impossible for me to get everything, so I'll just do a brain dump:

- A DNS request over UDP to get an IP address
- A TCP connection gets established (SYNs and ACKs!)
- A GET request is created by the browser and sent over that connection
- Tons of different things can happen after that, including proxying, database operations, cache reads, whatever.
- The browser might get a response back
- HTML gets parsed and a DOM tree gets generated
- External resources get fetched, scripts run
- Google analytics stuff gets injected so The Man can track you

That's pretty complicated stuff so far, but I haven't even mentioned stuff like...

- TCP retries
- Anycast DNS (which is really cool!)
- IP transit routing
- kernel queues and buffers
- load balancing
- SO MUCH MORE.

When I was younger, the Internet was magical. When I installed my first web server in 4th grade or so, it got less magical. You know what the weird thing is? Now that I know more about what happens, it seems so magical again!

