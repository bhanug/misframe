title: Don't panic! Type assertion safety.
date: 2013-12-18 18:36:00
url: dont-panic-type-assertion-safety

Go's `container/heap` package documentation has an [example](http://golang.org/pkg/container/heap/#example__intHeap) of an integer heap. Here's an interesting part of it:

	func (h *IntHeap) Push(x interface{}) {
		// Push and Pop use pointer receivers because they modify the slice's length,
		// not just its contents.
		*h = append(*h, x.(int))
	}

There's a little problem here. The point is, you shouldn't copy/paste that integer heap example into a package that you're going to use.

The last line (ignoring the brace) has a type assertion. `x` is an `interface{}` type, which means its type could be anything, but since `Push()` needs to append `x` to an array of `int`s, you have to assert that `x` is an `int`. If it's an `int`, it's all good. What if it's not an `int`?

> If the type assertion is false, a run-time panic occurs.
>
> — http://golang.org/ref/spec#Type_assertions

The thing about panics is that they bubble up. If you don't handle a panic by recovering, an entire goroutine will crash and, potentially, so will your entire program. Having a program crash isn't good (unless you want it to, of course)!

One of the more important things about writing good Go code is to not allow panics to escape package boundaries. It's not unlikely for someone to use a package without digging into its source code, so it should not be their responsibility to know about potentially panicky functions. Robert Griesemer said in a video with Erik Meijer that at the boundary, packages should return values and errors.

So what should you do?
----
Well, Go provides a way to check if type assertions are valid:

	v, ok := x.(T) // ok will be either true or false

But checking `ok` gets really annoying if you're doing many type assertions! Then just use `recover()`. Here's an example from the InfluxDB [source](https://github.com/influxdb/influxdb/blob/13c978abb1a25f56c89f6772e8056af97b91cb89/src/checkers/checkers.go#L16-L28):

	func (checker *inRangeChecker) Check(params []interface{}, names []string) (result bool, error string) {
		defer func() {
			if v := recover(); v != nil {
				result = false
				error = fmt.Sprint(v)
			}
		}()
		switch params[0].(type) {
		default:
			return false, "can't compare range for type"
		case int:
			p1 := params[0].(int)
			p2 := params[1].(int)
			p3 := params[2].(int)

What if I don't trust complicated code to handle panics?
-----
Sometimes you may be dealing with a complicated program that depends on a bunch of packages, and you have no idea whether or not there are potentially unhandled panics. One option to deal with this would be to use VividCortex's [Robustly](https://github.com/VividCortex/robustly). You can `robustly.Run()` a function and it will recover from panics should they occur.

