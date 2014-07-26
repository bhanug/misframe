title: State of the... state!
date: 2013-12-17 13:24:15
url: state-of-the-state

Three months ago, I started working on [fickle](https://github.com/PreetamJinka/fickle), a key-value store written in Go. The point of this post is to just jot down some thoughts on where the project is right now and where I'd like to take it.

Adversaria
---
This all goes back to [Adversaria](http://misfra.me/adversaria). It's that small program I wrote in Java to help me log traffic. I use it in combination with a little [Flot](http://www.flotcharts.org/) code to make pretty graphs:

![](https://31.media.tumblr.com/81741791c5f33df8624d6fc250f88d7e/tumblr_inline_mxyoeqMxeh1rs73cz.png)

It's really simple, which is great, but it's also limited. It's be great to have something that looked as simple but with more functionality.

Fickle
---
This whole thing sort of came together as I went along. Here's a summary of what Fickle actually is:

* Ordered key-value store (I love these!)
* In-memory
* It has networking
* Clients use a simple binary protocol
* Fast reads since values are stored in a hash map

At the core
---
Maps in Go are hash maps. I needed an ordered map (that's basically what a key-value store is). I had to write one: [Lexicon](https://github.com/PreetamJinka/lexicon). Lexicon uses another tiny [package](https://github.com/PreetamJinka/orderedlist) I wrote, which introduces ordering to Go's `container/list` package.

This *entire* thing is built on top of Lexicon. Fickle is essentially a wrapper that exposes the Lexicon data structure over a network. *That's all.*

Tests!
---
This is one of the major reasons why I chose Go to write this stuff: tests are *easy*.  Here's what tests look like for Fickle:

	func Test1(t *testing.T) {
		i := NewInstance(":12345")
		go i.Start()
		time.Sleep(time.Millisecond * 100) // Wait for it to start up
		conn, err := net.Dial("tcp", ":12345")
		if err != nil {
			t.Error(err)
		}
		write(conn, "foo", "bar")
		if !verify(conn) {
			t.Error("Bad write!")
		}
		if r := read(conn, "foo"); r != "bar" {
			t.Errorf("Bad read! Got %v", r)
		}
	}

Start up an instance, open up a TCP connection, write and read from the connection, and print any errors that come up. It's not just the main program that has tests -- its dependencies are also well-tested.

Not only that, all of these tests are run automatically at [Drone](https://drone.io/) after every push to GitHub:
![](https://31.media.tumblr.com/0ae411257310108a4ba2272e2407676a/tumblr_inline_mxypyb0fxy1rs73cz.png)

Did I mention Drone is free for open-source projects?

Keeping things safe
----
I don't want to lose my data, but as I mentioned, Fickle only stores data in memory. Everything is lost when the instance stops. Ouch!

There's two ways I can keep data: save it on the disk or have a replica. These aren't mutually exclusive. I actually had replication working earlier, but it was removed due to design changes.

The way I implemented replication was by broadcasting commands to replicas. Of course, this is incredibly trivial and doesn't guarantee consistency and many things can go wrong.

Saving to disk can also be trivial. I could just store a snapshot of the state and reload it, and I could also append a log of the operations sent to the database. The log would be "replayed" to restore data.

Transactions
----
I'd love to add transactions to this. It's probably going to be incredibly difficult, but I shall try!

Looking ahead
---
Have you seen this?
![](https://pbs.twimg.com/media/BbksqhsCIAASoUl.png:large)

It's a [CuBox-i](http://cubox-i.com/). It's a tiny ARM system that can have up to 4 cores and 2 GB of RAM. And it's not that expensive.

Here's an aside: I have a bunch of cPanel customers at [Bitcable](https://bitcable.com/), and I wanted to look at the distribution of MySQL database sizes. Here's what it looks like:
![](https://31.media.tumblr.com/63b2996bb8861072c304d4fcac2a10ef/tumblr_inline_mxyqoxzVFd1rs73cz.png)

Most are less than 99 MB! You could easily fit those into RAM, after accounting for overhead. It would be very realistic for me to, for example, put Misframe on a CuBox. I actually want to do that.

That's basically my goal for 2014: to get this key-value store to a point where I can depend on it to host my blog. It would also be cool to have a framework to test out interesting distributed computing things, like coordination protocols (eg. Paxos or Raft).

