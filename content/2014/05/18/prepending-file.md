---
title: Prepending to a file.
date: "2014-05-18"
url: /prepending-file
---


Prepending to a file is relatively expensive, but to what extent?
I decided to check. I wrote 2^23 integers to a file, in ASCII, separated
by a newline. On my ThinkPad with spinning rust, it takes a little longer
than 12 seconds. The file size is around 31 MB.

In order to "prepend" an integer to the beginning, we essentially have
to shift all of the other elements forward and then write what we need
at the beginning. So we write everything again, and then some. I wrote
a small Go program that does this, and here are the results:

```
  ∂ [-]: go run middle_insert.go 
Time to insert 4194304 integers: 12.202411374s
Time to prepend an int to the file: 579.472752ms
```

Prepending is 21 *times* faster! Dat cache.

A few notes worth mentioning:

* You have to be careful about overwriting, so you need to shift by
the byte-length of whatever you're prepending. Even if you're working in terms
of lines, files aren't.
* You can shift byte-by-byte, but that's not how I/O works.
* There's lots of stuff going on under-the-hood here. <span class='mono'>mmap</span> is a
complicated beast, in my opinion, and there are lots of gotchas.
* I think the size of the prepended value is irrelevant.
* All of the fancy stuff happens with <span class='mono'>syscall.Mmap()</span> and <span class='mono'>copy()</span>.
I thought this was an interesting use of slices!
* Even though prepending is faster, you really are rewriting a bunch. Write amplification
should not be ignored here, and it's pretty significant (~16,000,000)!

```go
package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

const N = 1 << 22

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	f, err := os.Create("scratch.txt")
	exitOnError(err)
	defer f.Close()

	start := time.Now()
	// Write a bunch of ints
	for i := 1; i <= N; i++ {
		fmt.Fprintln(f, i)
	}
	f.Sync()
	fmt.Println("Time to insert", N, "integers:", time.Now().Sub(start))

	fStat, err := f.Stat()
	exitOnError(err)

	toBeInserted := fmt.Sprintln(0)

	start = time.Now()
	os.Truncate(f.Name(), fStat.Size()+int64(len(toBeInserted)))

	// Mmap! Let the OS do the work.
	sl, err := syscall.Mmap(int(f.Fd()), 0, int(fStat.Size())+len(toBeInserted),
		syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	exitOnError(err)

	// This feels like cheating :-P
	copy(sl, append([]byte(toBeInserted), sl...))
	f.Sync()
	syscall.Munmap(sl)
	fmt.Println("Time to prepend an int to the file:", time.Now().Sub(start))
}
```
