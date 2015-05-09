---
title: Callback Magic with Go?
date: "2013-06-04"
url: /callback-magic-with-go
---


Go sounds great. You can write concurrent programs using goroutines and
channels without ever touching that callback nonsense, right? Yeah!

</p>

I love FoundationDB. On the surface it’s extremely easy to use — it’s
just a key-value store. All of the complexities are so well abstracted
that you don’t have to worry about them. I wanted to try to use
FoundationDB with Go. However, there’s no Go driver for it yet. I
decided I should try to fix that.

</p>

FoundationDB’s [C-API][] works with [futures][]. There’s no way around
them. Now, you can use futures in two ways. You can block and wait for a
future to be ready, or you can set a callback (ew!).

</p>

Blocking doesn’t sound so bad with Go, right? If you have something that
takes a long time, you’ll run it in a separate goroutine and keep going
in your program. Goroutines are cheaper than threads so you don’t have
to worry about a lot of overhead. I figured I could just launch a new
goroutine and block for the future inside. Note that this doesn’t block
the main program.

</p>

As it turns out, if you have a blocking system call in a goroutine, it
will not yield to the scheduler. Keep in mind that a single OS thread
will have multiple goroutines running on it. Go creates a maximum of N
OS threads where N is the number of CPU cores you have available. If the
Go runtime sees that a goroutine will not yield, it will spin up a new
*OS thread*! These new threads don’t count towards the maximum.

</p>

So here’s the deal. If I end up creating 100 futures and block each of
them in a separate goroutine, I’m essentially creating 100 OS threads.
Suddenly, the idea of cheap goroutines is gone and I’m left with a
poorly designed program. There’s only one option left: callbacks.

</p>

Here’s how to get futures working with Go. FoundationDB’s
`fdb_future_set_callback` takes a function pointer and a callback
argument that gets passed into the function. First you have to create a
Go function that takes a pointer to a channel and sends something on the
channel. This will be the callback. You’ll create a channel in your main
program, and when you need the future to be ready, you’ll read something
from the channel. If the future’s not ready, it’ll block. Otherwise,
you’ll get a result immediately.

</p>

I know this is hard to follow. It takes a while to wrap your head around
it. In the end, an end-user doesn’t really care. All of this callback
magic is abstracted away and you end up with a really easy-to-use
database driver.

</p>

[Take a look at the entire prototype.][]

</p>

  [C-API]: http://foundationdb.com/documentation/beta1/api-c.html
  [futures]: http://en.wikipedia.org/wiki/Futures_and_promises
  [Take a look at the entire prototype.]: https://github.com/PreetamJinka/fdbgo

