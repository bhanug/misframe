---
title: Thoughts on garbage collection
date: "2013-11-23"
url: /thoughts-on-garbage-collection
---


A few days ago, I was in the shower thinking about garbage collectors. I think some of my best ideas come from the shower, and I think it's because I don't have anything to distract me.

I was thinking about easy it was to use the [Hans Boehm garbage collector](http://www.hpl.hp.com/personal/Hans_Boehm/gc/):

> Empirically, this collector works with most unmodified C programs, simply by replacing malloc with GC_malloc calls, replacing realloc with GC_realloc calls, and removing free calls.

What a great abstraction! I was thinking about how it knew a block of memory should be freed. I thought, well... an easy way would be to just scan through the entire address space of a program and count up references. Of course, this isn't efficient *at all,* but it's a good start.

Cool. I had an idea. How would I actually implement it? The first step was to figure out how to scan through a program's address space. I won't go into detail, but let's just say I spent a good 3 or 4 hours following links and reading up on stack allocations and registers. Not fun.

I don't remember the details since most of this stuff happened really late at night, but I do remember having a really hard time trying to figure out how to search through address spaces. I tried a bunch of silly things, and got far too many memory access violations. Lots and lots of trial and failure.

At some point, I stumbled onto `/proc/self/maps`. I probably yelled out, "ARE YOU SERIOUS?" or something because I was really excited. The kernel tells us what memory mappings a program has access to! It's really silly, now that I think about it. I spent far too much time trying to figure this out. I just didn't know what I didn't know, so I didn't know what to search for on Google!

Hans Boehm garbage collector looks through `/proc/self/maps` too. I didn't look at its source code first to give myself a researching challenge ;).

Here's what I have so far: [GitHub Gist](https://gist.github.com/PreetamJinka/7611115).

Here's how it works:

1. Call `malloc` and remember the address. We need to pass that address to `free` later on.

2. Read through `/proc/self/maps` and get all the memory ranges

3. Look for possible pointers in that range, ignoring the pointer we use to keep track of the `malloc`'d memory.

4. If there aren't any references, free that block.*

*Here's the issue: there's always one reference on the stack. I'm not sure where that is. If I knew what it was, I'd ignore it :). Since I don't know where it is, I free a block if it has 1 reference. If you know how to fix it, let me know!

## Other thoughts

You'll never know whether some address in memory is a pointer or a value. The best thing to do is assume it's a pointer. This means it's conservative.

I have a VERY large amount of respect for the people who write garbage collectors. These things are complicated. They're hard to test. It's non-deterministic. They're confusing. Heck, even I got confused at times when I was writing stuff like this:

    if( (intptr_t)(*(void**)i) == (intptr_t)(address) &&
				(intptr_t)(i) != (uintptr_t)(ignore)

I'm casting a `void*` to a `void**` and then dereferencing it? Huh?! The more I think about it, the less it makes sense. It made sense to me late at night.

