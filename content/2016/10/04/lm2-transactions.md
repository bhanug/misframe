---
title: lm2 Transactions
date: "2016-10-04T22:00:00-04:00"
---

Welcome back. In my [previous post](2016/09/06/state-of-the-state-iv-lm2/) about lm2, I wrote the
following:

> I believe it's possible to implement ACID transactions on top of anything that provides the
following two properties:

> 1. Atomic & durable writes
> 2. Snapshot reads

> lm2 provides both. So, can you implement serializable, ACID transactions on top of lm2? Yeah!

> In fact, one really easy way is to just use a mutex. It's that simple! I'll leave it as an exercise
for you to figure out how to do it.

Did you take some time to figure it out? Here's my approach. The full code is on [GitHub]
(https://github.com/Preetam/lm2-layers/tree/master/transactional).

First, I added a writer lock. This is only used by writers. That means writers will block
each other, but readers don't block writers or each other.

```go
type Collection struct {
	col        *lm2.Collection
	writerLock sync.Mutex
}
```

Next, there are two new wrapper functions called `View` and `Update`, which represent
read-only and read-write transactions, respectively. Both wrappers create a *snapshot
cursor*, which allows a transaction to see the collection as it was when the transaction
began.

### View

The `View` function doesn't do much; it simply creates a cursor and passes it along.
lm2's cursors already provide snapshot reads. I think it's important to mention that
creating an lm2 cursor will provide a snapshot view up to the latest committed
collection update, so stale reads are not a problem.

```go
func (c *Collection) View(f func(*lm2.Cursor) error) error {
	cursor, err := c.col.NewCursor()
	if err != nil {
		return err
	}
	return f(cursor)
}
```

### Update

The `Update` function is a little more interesting. If you don't do any writes,
this function works exactly like the `View` function. However, because we have
the writer lock, this transaction is also guaranteed to have *the latest* view of the
collection. Other writers are blocked until the previous writer is finished.

The `Update` function also creates and applies a `WriteBatch`, which is guaranteed by lm2
to be durable and atomic.

```go
func (c *Collection) Update(f func(*lm2.Cursor, *lm2.WriteBatch) error) error {
	c.writerLock.Lock()
	defer c.writerLock.Unlock()

	cursor, err := c.col.NewCursor()
	if err != nil {
		return err
	}
	wb := lm2.NewWriteBatch()
	err = f(cursor, wb)
	if err != nil {
		return err
	}
	_, err = c.col.Update(wb)
	return err
}
```

That means transactions have

* Atomicity
* Consistency
* Isolation
* Durability

ACID! Woo!

## Transactional Squares

I created a [test]
(https://github.com/Preetam/lm2-layers/blob/master/transactional/transactional_test.go) to verify
some of these ACID properties, mainly consistency.

First off, what does *consistency* mean in an ACID context, anyway? If you're familiar with
distributed systems and CAP, keep in mind that *consistency in ACID is not like consistency in CAP*.
ACID consistency is about maintaining *constraints*. For SQL databases, these constraints may be
uniqueness of indexes, or foreign key constraints.

In order to test consistency with lm2, which is a flat ordered key-value store, I had to come up
with my own constraint. I created a simple set of formulas that describe the relations between three
keys: `a`, `b`, and `c`. I call this constraint *transactional squares*. It's easy to implement
and verify.

```
a = a
b = a*a
c = b*b
```

As an example, here's a table with each row representing a valid set for `a`, `b`, and `c`:

| a | b | c  |
|---|---|----|
| 2 | 4 | 16 |
| 1 | 1 | 1  |
| 3 | 9 | 81 |

In the transactional squares test, I ran several goroutines that

* Pick a random number and set `a`, `b`, and `c` to be squares (writer goroutines)
* Look at a snapshot view of `a`, `b`, and `c` to verify the constraint (reader goroutines)

Here are some *weird* consistency bugs I found. :)

[![](/img/2016/10/consistency1.jpg)](/img/2016/10/consistency1.jpg)

* `a` is repeated, and has a different value! Keys are unique, so this is wrong.

---

[![](/img/2016/10/consistency2.jpg)](/img/2016/10/consistency2.jpg)

* `446 * 446` is `198916`, so this is wrong.

---

[![](/img/2016/10/consistency3.jpg)](/img/2016/10/consistency3.jpg)

* You can see that one result is flipped.

---

This was fun to work on :).
