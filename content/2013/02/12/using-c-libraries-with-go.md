---
title: Using C libraries with Go.
date: "2013-02-12"
url: /using-c-libraries-with-go
---


Here's a simple example of how to use a C library from Go. I made a small library called libfoo.

## libfoo.c
	#include "libfoo.h"
	int foo() {
		return 1;
	}

## libfoo.h
	int foo();

Of course, you have to compile that as a shared object with GCC.

	$ gcc -shared libfoo.c -o libfoo.so

Make sure you copy over the binary to <span class="mono">/usr/lib/</span> and copy the header to <span class="mono">/usr/include/</span>.

## The Go file.
	package main

	/*
	#cgo LDFLAGS: -lfoo
	#include <libfoo/libfoo.h>
	*/
	import "C"

	import "fmt"

	func main() {
		fmt.Println("From libfoo: ", C.foo())
	}

...and that's it. I'm playing around with libguestfs and libvirt using Go. We'll see how that <em>Go</em>es. The puns never end!

