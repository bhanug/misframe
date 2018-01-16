---
title: Managing memory
date: "2014-01-13"
url: /managing-memory
---


I found [`libgc`](http://www.hpl.hp.com/personal/Hans_Boehm/gc/), a garbage collector library written in C, a few months ago and wanted to write something like it. The usage is really simple: replace all `malloc`s with `gc_malloc` and delete all occurrences of `free` in your code.

Now, a few months later, I'm thinking that it's not going to happen :). I'm not giving up, but rather changing my focus. I learned some more about other garbage collectors, especially those with persistent data structures. Actually, I could say that I've already written some form of garbage collector. My `vlmap` library, which is a versioned skip list, has a function to delete older nodes. The implementation resembles a mark-and-sweep collector --  nodes are marked as removed and sweeping occurs when you call a certain function.

The next step for me is to write a memory allocator. I've already been reading up on this and it sounds like a neat challenge. The idea is to partition a chunk of memory (it could be allocated memory, a memory-mapped file, a block device, etc) into chunks of different sizes and minimize fragmentation. There's lots of interesting techniques like bucketizing and keeping a free list, which is a list of chunks that you can use.

As usual, I'll probably do it wrong multiple times before getting it somewhat right.

