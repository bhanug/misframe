title: Nontrivial pipes
date: 2014-09-16 00:10:00
url: nontrivial-pipes

I'm going to take a simple concept, UNIX pipes, and basically frame (or misframe?) a nontrivial scenario.

    $ cat /proc/cpuinfo | grep CPU | wc -l
    4

We all know pipes, of course. You pipe stuff in and pipe stuff out. Easy enough. I think the piping in the example above is *context-free*. Each pipe has a stream of continuous bytes for its input and output. This is the simplest case.

Now let's say we're writing a program to simulate a user running a CLI application, like R. In our simulator, we'll execute `R` as a child process. Then, since R is a CLI program, we'll have to work with its stdin and stdout file descriptors (which are pipes).

It's easy enough to simulate a user typing in commands -- simply write data to R's stdin. To get the output, read from its stdout. We know this, and the entire scenario is simple enough. Let's also call it "context-free."

But what if you wanted to capture the output for each command separately? Now we're moving away from the context-free, infinite streams of bytes. It's more like chunked data. If you send R

    c(1:20)

you'll get back

     [1]  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20

Obviously the response, the printed vector, is part of the same "chunk" but the "context-free pipe" has no awareness of that.

This is where things get nontrivial. Our simulator needs to send `c(1:20)`, or some other form of input, and read the response from R. How long is the response? How many bytes do I need to read from the pipe? How will I know when it has finished outputting data?

Conclusion
---
Please, don't actually build a simulator like this for R. There are better ways to do this, like telling R to [output to a file](http://www.statmethods.net/interface/io.html).

However, if you "zoom out" and think about UNIX pipes in general, these concepts are important. This is why most protocols over TCP include a payload length. You can think of TCP sockets as pipes, and think about how these questions are formed in that scenario.

Anyway, I hope that made sense. Monday night thoughts :P.
