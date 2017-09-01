---
title: What's coming in Cistern v0.2.0
date: "2017-08-31T21:15:00-04:00"
---

Those who saw my tweets from last night know that I'm really exciting about what I've
been implementing in Cistern for the next release.

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">WHOA THIS IS SO AWESOME <a href="https://t.co/3k89naoaOV">pic.twitter.com/3k89naoaOV</a></p>&mdash; preetam 👨🏾‍💻 (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/903109561619447808">August 31, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

In the v0.1.0 [release post](/2017/08/01/cistern-v0-1-0/) from about a month ago I mentioned that
the next release will focus on generic JSON log messages. That's still really important for me,
but there are a couple of even more exciting things too.

### Query language

Cistern's storage and querying code came from Epsilon which didn't have a query language. It
only had a JSON REST API, and queries had to be described with a JSON object that sort-of resembled
a SQL query. Not user-friendly at all! It was a pain to keep writing that stuff by hand.

Cistern v0.1.0 introduced a pseudo query language in the CLI which uses flags and looks like this:

```text
  -columns 'sum(bytes), count(_id), max(packets), max(bytes)' \
  -group 'source_address, dest_address' \
  -order-by 'sum(bytes)' \
  -limit 3 \
  -descending
```

But that's still not very neat.

Cistern v0.2.0 will come with an *actual query language* with a well-defined [grammar](https://github.com/Cistern/cistern/blob/00ea921a8013dc891a8477d08f3338abed74c0d6/internal/query/grammar.peg)
that looks a lot like SQL. The parser translates a SQL-like string into the JSON query object I described before.

So this

```text
SELECT sum(bytes) GROUP BY dest_address ORDER BY sum(bytes) desc limit 100 POINT SIZE 3h
```

actually works! It's converted into a JSON query object which is now used internally.
I'm planning on writing up another post explaining how that works.

### Dashboard

Having a SQL-like query language makes the CLI just barely better than `curl`. My D3.js and Mithril.js
experiments from another project have been going so well that I decided to borrow that code to
make a UI for Cistern.

It only took me a few hours to come up with this:

[![](/img/2017/08/cistern-ui-1.jpg)](/img/2017/08/cistern-ui-1.jpg)

Here's another screenshot:

[![](/img/2017/08/cistern-ui-2.jpg)](/img/2017/08/cistern-ui-2.jpg)

I copied the layout from [Honeycomb.io](https://honeycomb.io/). It works really well.

### Generic JSON log messages

Of course, generic JSON messages will make it into v0.2.0.

### Release schedule

I hope to release v0.2.0 within a week, and then have a regular release once a month or so.
There probably won't be any huge changes after this. Just a lot of testing and refinement.
