---
title: Appending to a file instead.
date: "2014-05-19"
url: /appending-instead-of-prepending
---


In the previous post, I wrote about inserting records
into a file and keeping them ordered. The trivial solution of
shifting records in a file is horrible in terms of write amplification,
because we rewrite most, if not all, records on every single write
that does not occur at the end. So let's make that better by writing
to the end.

![][1]

In order to keep things in order, we need to know to move from the
newly inserted record back to the top of the file. Let's add that
pointer.

![][2]

If we started a read from the beginning of the file, we would
have no idea that the true first, or *root*, record is in fact somewhere else.
So we need a *root pointer*.

![][3]

We've now just constructed a linked list. More specifically, it's closer to an
[unrolled linked list](https://en.wikipedia.org/wiki/Unrolled_linked_list). There
isn't a lot written about unrolled linked lists online, but I think what I
described above is very similar to LevelDB's (and other LSM implementations')
structure of SSTables. They're just like unrolled linked lists.

[1]: /img/copied/posts/appending-instead-of-prepending/linked_list_file_1.jpg
[2]: /img/copied/posts/appending-instead-of-prepending/linked_list_file_2.jpg
[3]: /img/copied/posts/appending-instead-of-prepending/linked_list_file_3.jpg
