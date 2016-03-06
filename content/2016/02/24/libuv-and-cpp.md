---
title: Using libuv with C++
date: "2016-02-24T02:34:17.473Z"
---

libuv is the async I/O library built for Node.js. Node.js is essentially the glue between JavaScript
and libuv. How different would things be if you replaced JavaScript with C++? I kept thinking about
this as I started to work with C++14 and libuv. I wasn't doing this to try to make a "Node.cpp"; I
wanted to know how much of my previous experience would translate.

## Context

I think all of this stuff makes a lot more sense in context so I'll give you a simple C++ class
that involves everything covered here. `Peer` is a heavily simplified class coming from a
distributed systems project I've been working on. It represents a peer in a cluster. In this example
it doesn't do anything except have a method called periodically (for timeouts, etc), but you should
imagine other methods for network I/O, etc.

```c++
class Peer
{
public:
	Peer(const uv_loop_t* loop)
	: m_timer(std::make_unique<uv_timer_t>())
	{
		// Constructor
	}

	void
	periodic()
	{
		// Member function called by the timer
	}

	~Peer()
	{
		// Destructor
	}
private:
	std::unique_ptr<uv_timer_t> m_timer;
};
```

Each `Peer` owns its own libuv timer which is managed in a smart pointer.

## Lambdas

With JavaScript, you don't really have to worry about scoping within anonymous functions. It's more
explicit with C++ because you need to specify which variables will get captured (either by value or
reference). You can't use capturing lambdas as callbacks for libuv. Why? You need to consider how
capturing lambdas are implemented. Capturing lambdas have some stored state, and that state
has to exist somewhere. Consider the following capturing lambda:

```c++
int a = 5;
[&a]() { // capture a reference to a.
	std::cout << "a is " << a << std::endl;
}
```

*This lambda is not just a function* even though it looks like one. The compiler actually generates
a new class with a reference member that stores `&a`. When the lambda is executed, an instance of
that class is created within the usual C++ lifetime semantics. When the lambda instance goes out of
scope, it gets destructed. This means you can't use capturing lambdas as C callbacks.

Fortunately you *can* use non-capturing lambdas since they can be converted to function pointers,
and therefore be passed as callbacks to C libraries.

In our `Peer` class, the first thing we need to do is setup our `uv_timer_t` in the constructor.

```c++
Peer(const uv_loop_t* loop)
: m_timer(std::make_unique<uv_timer_t>())
{
	// Initialize the timer.
	uv_timer_init(loop, m_timer.get());
	// Start it.
	uv_timer_start(m_timer.get(), [](uv_timer_t* timer) {
		// This is a non-capturing lambda so we can use it as a callback.
	},
	// Repeat once a second.
	1000, 1000);
}
```

Since we can't capture `Peer` directly, we'll have to be clever and use the timer handle's `data`
field to get access to the `Peer` within the callback. Every libuv handle has a `void* data` field.
Here's how you use it:

```c++
// (Within the Peer() constructor)
m_timer->data = this; // Note: m_timer->data is a void*
uv_timer_start(m_timer.get(), [](uv_timer_t* timer) {
	auto self = (Peer*)timer->data;
	// Call the method.
	self->periodic();
}, 1000, 1000);
```

Without a lot of work I think this technique keeps things fairly clean.

Here is the final constructor:

```c++
Peer(const uv_loop_t* loop)
: m_timer(std::make_unique<uv_timer_t>())
{
	// Initialize the timer.
	uv_timer_init(loop, m_timer.get());

	// Set up data pointer.
	m_timer->data = this;

	// Start it.
	uv_timer_start(m_timer.get(), [](uv_timer_t* timer) {
		auto self = (Peer*)timer->data;
		// Call the method.
		self->periodic();
	},
	// Repeat once a second.
	1000, 1000);
}
```

You don't have to do anything fancy with the `periodic` method.

## Smart Pointers and Destructors

The most important part of this class is the destructor and how the timer gets cleaned up.

Asking libuv to stop the timer is easy. It's just

```c++
uv_timer_stop(m_timer.get());
```

But you still need to close the timer, as with all other libuv handles. The issue is that `uv_close`
is asynchronous, so if you call it within a destructor it will probably be called long after the
destructor finishes executing! That means the following **doesn't work.**

```c++
~Peer()
{
	// Destructor
	uv_timer_stop(m_timer.get());
	uv_close((uv_handle_t*)m_timer.get(), [](uv_handle_t* handle) {
	});
}

// Close may happen now, but the m_timer unique_ptr is gone!
// This becomes a "use after free" scenario.
```

The correct way to do this would be to release ownership from the smart pointer and call `delete`
manually.

```c++
~Peer()
{
	// Destructor
	uv_timer_stop(m_timer.get());

	// Release ownership.
	auto handle = m_timer.release();
	uv_close((uv_handle_t*)handle, [](uv_handle_t* handle) {
		delete handle;
	});
}
```

## Final thoughts

Using libuv with C++ is a little weird if you started off learning modern C++ like me. Once you have
some code to bridge the two together it's not that bad! I know I can avoid most of this by using
Boost.Asio, a pure C++ library, but I think libuv stuff is easier to read.

---

## Final Class Definition

```c++
class Peer
{
public:
	Peer(const uv_loop_t* loop)
	: m_timer(std::make_unique<uv_timer_t>())
	{
		// Initialize the timer.
		uv_timer_init(loop, m_timer.get());

		// Set up data pointer.
		m_timer->data = this;

		// Start it.
		uv_timer_start(m_timer.get(), [](uv_timer_t* timer) {
			auto self = (Peer*)timer->data;
			// Call the method.
			self->periodic();
		},
		// Repeat once a second.
		1000, 1000);
	}

	void
	periodic()
	{
		// Some periodic behavior
	}

	~Peer()
	{
		// Destructor
		uv_timer_stop(m_timer.get());

		// Release ownership.
		auto handle = m_timer.release();
		uv_close((uv_handle_t*)handle, [](uv_handle_t* handle) {
			delete handle;
		});
	}
private:
	std::unique_ptr<uv_timer_t> m_timer;
};
```
