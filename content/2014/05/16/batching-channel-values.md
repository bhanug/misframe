---
title: Batching channel values.
date: "2014-05-16"
url: /batching-channel-values
---


This was an interesting use of channels in Go from today. The goal was to
listen on a channel for a stream of values as they arrived one-by-one. We
then had to process and batch them up in order to get them into a certain
format, and then send them on another channel.

The following code shows a simpler example: listening for a stream of <span class='mono'>ints</span>
and batching them into slices. A batch is sent after 15 elements are gathered
or 100 ms after the last send.

See it in action at the [Go Playground](http://play.golang.org/p/uaBxvMw1x1)!

Thanks to [John Berryman (@JnBrymn)](https://twitter.com/JnBrymn) for providing a sanity
check and the original skeleton program.

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    from := make(chan int)
    to := make(chan []int)

    go func() {
        for i := 0; i < 500; i++ {
            from <- i
            // vary the time between sends
            time.Sleep(time.Duration(rand.Intn(15)) * time.Millisecond)
        }
    }()

    go func() {
        for {
            fmt.Println(<-to)
        }
    }()

    var wait = time.After(time.Millisecond * 100)
    var buffer = make([]int, 0)

    for {
        select {
        case <-wait:
            if len(buffer) > 0 {
                to <- buffer
                buffer = buffer[:0]
            }
            wait = time.After(time.Millisecond * 100)
        default:
            buffer = append(buffer, <-from)
            if len(buffer) == 15 {
                to <- buffer
                buffer = buffer[:0]
                wait = time.After(time.Millisecond * 100)
            }
        }
    }
}
```
