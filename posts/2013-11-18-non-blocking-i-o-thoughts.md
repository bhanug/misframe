title: Non-blocking I/O thoughts
date: 2013-11-18 21:14:00
url: non-blocking-i-o-thoughts

I started using Node.js a long time ago. It was the first time I saw async execution outside of the browser. It was cool -- still is.

I worked with Node internals for a significant part of a summer. It doesn't take long before you realize that the event loop is one of the most important components of Node. Node uses libuv for its evented I/O, and I sort of got into it, but not really. Conceptually, it sort of made sense.

I've been working a lot using Go and C lately, and also thinking about async stuff. Something you can do in Node is pass a callback to an async file reading function and have that operation not block. Cool. You can also read from a socket asynchronously. Great! How would I do that in C?

I started reading. Hey, you can do non-blocking socket reads using plain ol' vanilla system calls. The kernel is nice enough to provide that for us. I can do async stuff with sockets, but what about files? I started reading some more.

No dice.

Huh... but how is Node doing it? It's non-blocking and single-threaded, right? It can't be using blocking syscalls!

Wrong!

It's kind of embarrassing that I didn't realize this until now. Node is *not single-threaded.*

> The libuv filesystem operations are different from socket operations. Socket operations use the non-blocking operations provided by the operating system. **Filesystem operations use blocking functions internally, but invoke these functions in a thread pool** and notify watchers registered with the event loop when application interaction is required.
>
> -- http://nikhilm.github.io/uvbook/filesystem.html

Doh! The Node event loop is single-threaded, but if you do async file reads, for example, another thread *will be used*.

By now, I'm thinking async file I/O is annoying in C, unless you use a library to manage a thread pool for you.

## libdispatch

At vBSDCon, I heard about [libdispatch](http://en.wikipedia.org/wiki/Grand_Central_Dispatch). It's really interesting! I wouldn't call it an async I/O library (because it's not) -- it's just a different way of thinking about things.

> Dispatch Queues are objects that maintain a queue of tasks, either anonymous code blocks or functions, and execute these tasks in their turn. The library automatically creates several queues with different priority levels that execute several tasks concurrently, selecting the optimal number of tasks to run based on the operating environment. A client to the library may also create any number of serial queues, which execute tasks in the order they are submitted, one at a time.

