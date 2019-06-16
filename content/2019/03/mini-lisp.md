---
title: Mini Lisp
date: "2019-03-06T07:42:00.558Z"
---

I implemented a small Lisp interpreter over the weekend. You can find it here:
https://github.com/Preetam/mini-lisp. It's about 400 lines of Go code so far.

I started writing an interpreter a couple of months ago using the [mal - Make a
Lisp](https://github.com/kanaka/mal/) guide. After step 4 my implementation felt
really messy, and I felt I was just doing what the guide told me without learning
too much about how things actually worked. Later I found Peter Norvig's
[(How to Write a (Lisp) Interpreter (in Python))](https://norvig.com/lispy.html)
which is a much simpler version, and that was the inspiration to start over.

Here's what you can do with it. Let's start with a simple factorial function:

<!--more-->

```lisp
;; Factorial
(define fact
  (lambda (n)
    (if (<= n 1)
      1
      (* n (fact (- n 1)))
    )
  )
)

(print (str (fact 5)))
; 120
```

There are mathematical operations, `if`, printing, and lambas.

Using [(An ((Even Better) Lisp) Interpreter (in Python))](https://norvig.com/lispy2.html),
I also added tail recursion optimization. I was almost at that point in the mal guide too.
Tail recursion optimization makes it really cheap to execute certain functions implemented
recursively. Something like the following function, which takes advantage of tail recursion
optimization, can call itself practically infinitely without running out of stack space.

```lisp
;; sum2 sums numbers up to n.
;; It uses a tail recursion optimization.
(define sum2
  (lambda (n acc)
    (if (= n 0)
      acc
      (sum2 (- n 1) (+ n acc))
    )
  )
)

(print (str (sum2 1000 0)))
; 500500
```

Finally, the most interesting part I got to (which also took the most time) is
[call/cc](https://en.wikipedia.org/wiki/Call-with-current-continuation). call/cc
is used to implement more complicated control flows and continuation objects.
It's very hard to implement call/cc entirely, so I decided to only implement the
simplified version that Norvig wrote about. It behaves like `try/catch` so I
decided to call mine `catch!`.

```lisp
;; catch! examples

(catch! (lambda (throw)
  (+ 5 (* 10 (catch! (lambda (escape) (* 100 (throw 3)))))))
)
; 3

(catch! (lambda (throw)
  (+ 5 (* 10 (catch! (lambda (escape) (* 100 (escape 3)))))))
)
; 35
```

Finally, I can pass it a file name to interpret so I can start writing
scripts with it.

```
#!/usr/bin/env mini-lisp
; In ./lisp_script

(print (str "hello, world!"))
```

```
$ ./lisp_script
"hello, world!"
```

That's it! There isn't a lot you can do with it now, but over time I can add more
built-in functions. I'll still consider this "done" for now. Not a bad result for
a weekend project!
