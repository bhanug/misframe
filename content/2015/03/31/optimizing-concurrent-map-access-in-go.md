---
title: Optimizing Concurrent Map Access in Go
date: "2015-03-31"
url: /optimizing-concurrent-map-access-in-go
summary: "7x performance with 4 lines changed"
---

One of the more contentious sections of code in [Catena](https://github.com/PreetamJinka/catena), my time series storage engine, is the function that fetches a `metricSource` given its name. Every insert operation has to call this function at least once, but realistically it will be called potentially hundreds or thousands of times. This also happens across multiple goroutines, so we'll have to have some sort of synchronization.

The purpose of this function is to retrieve a pointer to an object given its name. If it doesn't exist, it creates one and returns a pointer to it. The data structure used is a `map[string]*metricSource`. The key fact to remember is that elements are *only inserted* into the map.

Here is a simple implementation. I have excluded the function header and return statement to save space.
```go
var source *memorySource
var present bool

p.lock.Lock() // lock the mutex
defer p.lock.Unlock() // unlock the mutex at the end

if source, present = p.sources[name]; !present {
	// The source wasn't found, so we'll create it.
	source = &memorySource{
		name: name,
		metrics: map[string]*memoryMetric{},
	}

	// Insert the newly created *memorySource.
	p.sources[name] = source
}
```

I have a benchmark that inserts time series points into the database. Again, each insert has to
call this function to get the pointer to the metric source it has to update.

This one gets about **1,400,000 inserts / sec** with four goroutines running in parallel
(i.e. `GOMAXPROCS` is set to 4). This may seem fast, but it's actually *slower* than having
one goroutine do all the work. If you're thinking lock contention, you're right.

So, what's the problem here? Let's consider a simplified case where there are no
inserts into the map. Suppose goroutine 1 wants to get source "a" and goroutine 2 wants
to get "b", and assume "a" and "b" are already in the map. With the given implementation,
the first one will grab the lock, get the pointer, unlock, and move on. Meanwhile, the other
goroutine is stuck waiting to grab the lock. Waiting on that lock seems like a pretty bad use of time!
This gets worse and worse as you add more goroutines.

One way to make this faster is to remove the lock and make sure only one goroutine accesses the map.
That's simple enough but you have to give up scalability. Here's an alternative that's just as simple
*and* maintains thread-safety.

This change only takes one more line and an additional character, but will keep getting faster as
you scale up.

```go
var source *memorySource
var present bool

if source, present = p.sources[name]; !present { // added this line
	// The source wasn't found, so we'll create it.

	p.lock.Lock() // lock the mutex
	defer p.lock.Unlock() // unlock at the end

	if source, present = p.sources[name]; !present {
		source = &memorySource{
			name: name,
			metrics: map[string]*memoryMetric{},
		}

		// Insert the newly created *memorySource.
		p.sources[name] = source
	}
	// if present is true, then another goroutine has already inserted
	// the element we want, and source is set to what we want.

} // added this line

// Note that if the source was present, we avoid the lock completely!
```

**5,500,000 inserts / sec.** This is **3.93 times** as fast. Recall that I had four goroutines
running in parallel, so this increase makes sense.

This works because we're never deleting sources, and the addresses don't change. If we have
a pointer address in CPU cache, we can use it safely even if the map is changing below us.
Notice how we still need the mutex. If we didn't have it, there would be a race condition
where one goroutine will realize that it has to create the source and insert it, but another
may insert it in the middle of that sequence. This way, we only hit the lock during inserts into
the map, but those are relatively rare.

My colleague [John Potocny](https://twitter.com/JohnPotocny1) suggested that I remove the `defer`
because it has nontrivial overhead. He was right. One more *very* minor change and I was amazed
at the result.

```go
var source *memorySource
var present bool

if source, present = p.sources[name]; !present {
	// The source wasn't found, so we'll create it.

	p.lock.Lock() // lock the mutex
	if source, present = p.sources[name]; !present {
		source = &memorySource{
			name: name,
			metrics: map[string]*memoryMetric{},
		}

		// Insert the newly created *memorySource.
		p.sources[name] = source
	}
	p.lock.Unlock() // unlock the mutex
}

// Note that if the source was present, we avoid the lock completely!
```

This version gets **9,800,000 inserts / sec**. That's **7 times** faster
with only about 4 lines changed.

### Edit:

Is this correct? Unfortunately, no! There is still a race condition, and it's easy to find
using the race detector. We can't guarantee the integrity of the map for readers while there
is a writer.

Here is the race-free, thread-safe, "correct" version. Using an RWMutex, readers won't block each other
but writers will still be synchronized.
```go
var source *memorySource
var present bool

p.lock.RLock()
if source, present = p.sources[name]; !present {
	// The source wasn't found, so we'll create it.
	p.lock.RUnlock()
	p.lock.Lock()
	if source, present = p.sources[name]; !present {
		source = &memorySource{
			name: name,
			metrics: map[string]*memoryMetric{},
		}

		// Insert the newly created *memorySource.
		p.sources[name] = source
	}
	p.lock.Unlock()
} else {
	p.lock.RUnlock()
}
```
This version is **93.8%** as fast as the previous one, so still very good. Of course, the previous version
isn't correct, so there shouldn't even be a comparison.
