title: Thoughts on Go.
date: 2013-06-29 23:50:00
url: thoughts-on-go

At the beginning of the year, I told myself I would learn [Go][]. Stuff
about Go seems to pop up on the front page of Hacker News frequently. In
fact, it’s gotten to the point where I think people upvote articles that
mention Go in their titles. People really like it, or they just like
talking about it. In any case, I think it’s worth trying out.

</p>

I’ve been using Go for a few months now. These days I use it over around
8 hours a day. I feel like sharing my thoughts about it.

</p>

Pros
----

</p>

-   Goroutines. Built-in concurrency is *awesome.* You know how people
    say learning Haskell (or any functional language, for that matter)
    makes you think a little differently? I think Go is like that too.
    You start thinking a little more… concurrently.
-   The syntax isn’t weird. It looks a lot like C. I like C.
-   It’s compiled. It’s fast. It produces a single binary that just
    works.
-   The tools. Oh, man. The tools are wonderful. `go fmt` is great, and
    `go test` is absolutely amazing. I like how the language encourages
    building and testing packages. It just reinforces good application
    structure.
-   Fast compilation makes it feel like an interpreted program. `go run`
    and you’re running a native binary.
-   Garbage collection!
-   Structured, not object-oriented. OOP just seems too bulky for me. Go
    has interfaces which remind me of OOP but they’re a lot simpler.
-   Packages! The standard packages are quite good and it’s incredibly
    easy to `go get` others.
-   Using C libraries is extremely easy. I wrote a couple of blog posts
    showing examples.

Cons
----

</p>

-   It’s too magical. Often times I’m not really sure how things work,
    so I never have guarantees. Garbage collection always seems to be
    happening in some magical box. A really interesting memory issue was
    found with JavaScript [earlier][]. This isn’t really a con, but
    sometimes I wonder how stable the core runtime really is.
-   Writing REST APIs is kind of annoying. I’m coming from years of
    working with Node, and Node with Express is honestly my favorite way
    of building REST APIs. I’d love to use Go, but the routing stuff
    just looks sucky. Maybe this is an idea for a package…?

It’s hard for me to think of any other cons. I know I haven’t been using
it that long, so I guess we’ll just have to wait and see. Working with
Go is really fun. Try it out!

</p>

  [Go]: http://golang.org/
  [earlier]: http://point.davidglasser.net/2013/06/27/surprising-javascript-memory-leak.html

