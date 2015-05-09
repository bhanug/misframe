---
title: Fun to obligation.
date: "2014-07-27"
url: /fun-to-obligation
---


A few months ago I was sitting in the back seat of a minivan on the way to New Jersey to visit family.
I had music playing through my headphones and I was staring at a bunch of [sFlow specification
documents](http://www.sflow.org/developers/specifications.php) that I had downloaded onto my laptop.
sFlow is a packet and counter sampling specification, and I wanted to write a decoder in Go.

It's pretty neat to have the entire history on GitHub. I had my first [interesting commit](https://github.com/PreetamJinka/sflow-go/commit/64a8afb40eaa322f655105ae2f569550335219fc) pushed at
9 PM on a Sunday. You can browse through the commits to see how I went from a prototype that
barely did anything, to a much higher quality Go package. I worked on it in my free time (because
I have so much :P).

I sometimes tell people about sFlow and my decoder, but I find that no one knows what it is.
I guess it makes sense since sFlow usage is usually limited to like... network engineers and the like.
Even the network engineers I know aren't really programmers. So yeah, no one is really using my package
or talking about it other than me. It's just a personal project, and I like to work on it whenever
I have time. It's largely incomplete simply because I don't need the rest of the features,
but I still use it and make fun [little tools](https://github.com/PreetamJinka/flowtools) for demonstration.

The other day I got a notification<sup>1</sup> about a pull request: https://github.com/PreetamJinka/sflow-go/pull/10.
It was a bug fix, and the person seems to be from Turkey and works for an ISP. I was actually shocked, and for two
reasons. First, I didn't know that people would actually use this personal project of mine. Cool! And
I didn't know it was buggy! I didn't notice anything wrong, but I haven't tested it much (it is a personal project,
after all).

I saw the email in the morning so I had to wait until after work to take a look. I think in general, if
I see something wrong with one of my projects I just make myself a reminder and fix it later. It was different
this time because I knew someone was expecting a fix :P. I went to sleep at like 3 AM after finding a fix.

I spent a lot of my free time working pro bono to fix this. Turns out it's still not fixed! Again, I need
to spend time trying to fix it for nothing in return. Feels a lot like an obligation at this point.

Something I started for fun to pass the time is now an obligation eating up my time. Interesting. I also think
it says something about refinement. Figuring something out and creating something new always seems to be fun,
but refinement seems tedious. But hey, I think I'm becoming a huge fan of nitpicking.

By the way, I have high expectations for this package. The only other implementations of the sFlow protocol
that I've seen are written in C and Perl. I can't imagine anyone writing a high-performance, "cloud"-ready
flow analyzer using C or Perl. If you see one, let me know! Go, on the other hand, seems much more suitable.
I want this package to be the most elegant (and performant) sFlow decoder (and encoder, hopefully :P) implementation
out there.

---

1. By the way, GitHub's filtering is either broken or too complicated. I have no idea why I see organization
events on my personal newsfeed.
