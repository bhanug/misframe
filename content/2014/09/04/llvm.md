---
title: LLVM
date: "2014-09-04"
url: /llvm
---


For some reason, instead of studying for my math quiz today (it's past
midnight right now), I decided to play around with LLVM. One of my
homework exercises was to solve the Fibonacci difference equation,

F(n) = F(n-1) + F(n-2) for n >= 2,

so that's why I wrote a
recursive Fibonacci function.

I think it's pretty cool considering I never formally learned
assembly.

```
declare i32 @printf(i8* noalias nocapture, ...)

@numPrintStr = constant [27 x i8] c"#%d Fibonacci number is %d\00"

define void @printNumber(i32 %a) {
	%f = call i32 @fib(i32 %a)
	call i32 (i8*, ...)* @printf(
		i8* getelementptr([27 x i8]* @numPrintStr, i32 0, i32 0),
		i32 %a, i32 %f)

	ret void
}

define i32 @fib(i32 %a) {
entry:
	switch i32 %a, label %recur [ i32 1, label %base
	                              i32 2, label %base ]

base:
	ret i32 1

recur:
	%prev1 = sub i32 %a, 1
	%prev2 = sub i32 %a, 2
	%prev1val = call i32 @fib(i32 %prev1)
	%prev2val = call i32 @fib(i32 %prev2)
	%sum = add i32 %prev1val, %prev2val
	ret i32 %sum
}

define i32 @main() {
entry:
	call void @printNumber(i32 40)
    ret i32 0
}
```

Output:
```
$ llc fib.ll -o fib.s
$ gcc fib.s -o fib
$ ./fib
#40 Fibonacci number is 102334155
```

There's probably a way to generate a binary without using
gcc, but it's late and I'm too lazy to figure it out.
