---
title: '“Schema changes” with a key-value store'
date: "2017-05-01T23:55:00-04:00"
---

I've mentioned on Twitter that I'm working on a metadata storage service built on lm2,
my key-value store. One of the things that I started thinking about while I'm implementing
this is schema changes. But they're not the typical kinds of schema changes I'm used to with
relational databases because I'm not working with a relational database :).

With relational databases, I usually have a fixed schema that all rows have to conform to.

| Column A | Column B |
|----------|----------|
| Alice    | 1        |
| Bob      | 2        |
| Bob      | 3        |

At some point, I need to modify that schema with an `ALTER` to add a column, remove a column, etc.

| Column A | Column B | Column C |
|----------|----------|----------|
| Alice    | 1        | X        |
| Bob      | 2        | Y        |
| Bob      | 3        | Z        |

With a key-value store, I basically have a two-column table with a fixed structure.

| Key | Value |
|-----|---------------------------|
| id=1 | username=alice, email=alice@example.com, name=Alice |
| id=2 | username=bob, email=bob@example.com, name=Bob |
| id=3 | username=bob2, email=bob2@example.com, name=Bob |

With key-value stores, instead of having lots of columns, you have to encode columns into
the keys and values. With ordered key-value stores, you need to use prefixes to preserve locality
and make range reads easier.

What if your design changes and you need some new encoding or mapping? You have to rewrite all
of the affected key-value pairs!

Something that I wanted to avoid with my service is taking it down to rewrite all of my data
during schema changes. This is something that comes up with relational databases too. ALTERs
can lock and rewrite entire tables, which can prevent queries from using them and take a long time.
On the bright side, these days we're starting to see more systems capable of doing *online* ALTERs
so you can keep using those tables while the operation happens in the background.

So I decided to use a technique similar to what those fancy relational databases are doing.
It's pretty simple. I can just add a version to each key-value pair and have my application
rewrite it on demand.

Here's an example showing a `deleted` "column" being added:

| Key | Value |
|-----|---------------------------|
| id=1 | **version=1**, username=alice, email=alice@example.com, name=Alice |
| id=2 | **version=2**, **deleted=0**, username=bob, email=bob@example.com, name=Bob |
| id=3 | **version=1**, username=bob2, email=bob2@example.com, name=Bob |

You can use this technique to implement lots of other "online schema changes" with key-value stores.
