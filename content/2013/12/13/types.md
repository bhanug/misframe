---
title: Types...
date: "2013-12-13"
url: /types
---


What does this print?

----

	package main

	import "fmt"

	type Foo interface {
		foo()
	}

	type A int
	type B A

	func (a A) foo() {}
	func (b B) foo() {}

	func main() {
		hashmap := make(map[Foo]string)
		a := A(1)
		b := B(1)
		hashmap[a] = "foo"
		hashmap[b] = "bar"
		fmt.Println(hashmap)
	}

----

Let's see... there's a map that takes keys of type `Foo`, which is an interface. `a` and `b` both implement `foo()`, so they're of the `Foo` interface. So the output should be `map[1:bar]`, right?

Nope. It's actually `map[1:foo 1:bar]`.

I think this is weird. :-/

https://play.golang.org/p/AlmujxUWRa

