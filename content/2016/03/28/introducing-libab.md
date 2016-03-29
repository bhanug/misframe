---
title: Introducing libab, another broadcast library
date: "2016-03-29T01:48:06.405Z"
---

Way back in June 2015 I [wrote](/2015/06/29/current-projects/) about my failure detector project.
I wanted to implement failure detectors for two reasons: to learn C++ and get a better understanding
of an important and fundamental distributed systems topic.

After all these months, all of my work has converged to [libab](https://github.com/Preetam/libab).
It's a small-ish C library that allows you to create a cluster of interconnected nodes and broadcast
a sequence of messages to each of them. It sounds simple enough, but there are two important
properties:

1. Each *committed* message is guaranteed to be present on a majority of the nodes.
2. Messages are *committed* in order.

If you're thinking that this sounds like [atomic broadcast](https://en.wikipedia.org/wiki/Atomic_broadcast),
you're right. It's a lot like atomic broadcast, but not entirely. Atomic broadcast requires that
each participant receives *all* messages; libab only requires a majority.

"So, what's the point?" you may ask. My goal for this project is to implement enough useful
primitives so it's possible to build much more complicated distributed systems primitives on top.
Think about it: given the two properties I listed above, could you write something on top of this
library that implements atomic broadcast? I think you can!

It's available on GitHub for you to take a look at: https://github.com/Preetam/libab

I'm planning on writing a proxy that uses libab in front of Redis. Consistent Redis?! Hopefully I'll
announce something like that in the next *State of the State*!

---

Writing a simulation would've been easy, but I wanted to implement the real thing. That means I
needed a peer-to-peer messaging library that would allow me to broadcast messages between failure
detector nodes.

I wanted the following features:

1. Transparent reconnections
2. TCP
3. Every node should be connected to every other node
4. Easy-to-use broadcasts

After searching for a while, I couldn't find a library that suited my needs. I ended up writing it
myself. It took a long time to get it right (I didn't start with libuv), but it ended up being
really fun to eventually get right.

---

Finally, a list of things I find interesting:

- [Custom message serialization](https://github.com/Preetam/libab/blob/70e2306262201c56b927c47f5bec6b83e82a3be3/src/message/message.hpp)
	- After writing all of those protocol decoders I feel like I could do this in my sleep now :P
- [CMake](https://github.com/Preetam/libab/blob/70e2306262201c56b927c47f5bec6b83e82a3be3/CMakeLists.txt)
	- First time using CMake. It's awesome.
- libuv
- [Automatic reconnections](https://github.com/Preetam/libab/blob/70e2306262201c56b927c47f5bec6b83e82a3be3/src/peer/peer.cc#L86-L141)
	- Nodes have to send each other their reconnection addresses. They can send each other requests too.
	- This is *super* annoying, especially when two nodes reconnect to each other and you have to
      detect and handle the duplicates connections!
- [Separated state-machine implementations](https://github.com/Preetam/libab/blob/70e2306262201c56b927c47f5bec6b83e82a3be3/src/node/role.hpp)
	- These are *almost* ready to be used in independent simulations/tests.
