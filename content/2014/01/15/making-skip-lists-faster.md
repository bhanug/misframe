---
title: Making skip lists faster
date: "2014-01-15"
url: /making-skip-lists-faster
---


First off, watch this video: [Bjarne Stroustrup: Why you should avoid Linked Lists](http://www.youtube.com/watch?v=YQs6IC-vgmo). To summarize, Stroustrup states that using vectors is significantly better than using linked lists because linked lists are horrible at cache hits.

[![](http://daviddeley.com/programming/docs/listvsvector.png)](http://daviddeley.com/programming/docs/listvsvector.png)
<br/>(From http://daviddeley.com/programming/docs/page_faults_and_array_addressing.htm)

Notice that reading sequentially from a vector means sequential reads in memory, but sequential reads from a linked list is comparable to random (bad for cache) reads in memory.

How would you make a linked list faster? Put the elements in sequential order in memory. By the way, when I write *memory*, I could also mean a file mapped to memory. Sequential I/O is much faster than random I/O. Practically, it's difficult to rearrange objects in memory. It's incredibly difficult when you have lots of objects.

Skip lists are interesting. They're probabilistic, and you always know which elements will be accessed most frequently (the nodes with the highest levels). I think you could make skip lists faster by putting all of the high-level nodes closer together in memory. Hopefully you'll have a contiguous chunk of memory with all of your frequently accessed nodes which can reside entirely in some level of CPU cache.

If you know how many elements your skip list will have, you can pre-allocate a contiguous chunk of memory with a certain size and only use that for the high-level nodes.

The issue of random reads still remains -- if you do sequential access on the bottom level, you'll still do random reads. More on this later.

