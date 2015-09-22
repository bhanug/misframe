---
title: C is weird.
date: "2013-12-26"
url: /c-is-weird
---


> Falling off the end of a function that is declared to return a value (without explicitly returning a value) leads to undefined consequences.
>
> — http://stackoverflow.com/questions/293499/what-happens-if-you-dont-return-a-value-in-c

Yeah. I forgot to write a `return` line in a function that was supposed to return a pointer. I didn't notice it for days since the program worked fine -- the right thing was being returned!

I tried it out on my BeagleBone Black and it looked like `calloc()` was returning `NULL`. Weird! Then I noticed that I wasn't actually returning anything.

It's interesting that these platforms have different results.

[Here's my oops.](https://github.com/Preetam/vlmap/commit/49ed1e966abc491#diff-e08843ac041a0a54fa44b93b13f7687cL31)

Apparently the program put the result of `calloc()` onto the [register used for the return value on x86](http://stackoverflow.com/questions/7280877/why-and-how-does-gcc-compile-a-function-with-a-missing-return-statement)! Neat!

