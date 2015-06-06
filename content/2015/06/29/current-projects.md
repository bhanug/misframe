---
title: Current Projects
date: "2015-06-28"
---

My interests generally seem to change every few months, but I tend to hover around the same topics. These days, I've been focusing on peer-based communication, consensus, failure detectors, and C++.

I’m a fan of thinking big but starting small<a class='ref' href='#endnote-ref-1'>[1]</a>. The "big" project I have right now is a failure detector. A failure detector simply checks for failures among nodes in a distributed system using some sort of ping<a class='ref' href='#endnote-ref-2'>[2]</a>. I am starting to write this using C++. I've worked with C++ before, but I never actually "learned" it. It's been quite a while since I worked with an object-oriented language, and that was back when I was first learning how to program. I'm writing C++ wrappers for POSIX utilities in order to build up my intuition about C++. It turns out that this is also helping me learn a lot more about POSIX calls.

As a Go programmer, I'm generally spoiled by the standard library. There are so many things that I don't have to think about, and this is true for other languages as well. For example, you don't have to think about parsing IP address strings or setting up a socket to create a TCP listener. For my wrappers, I'm doing things like creating an IP address class<a class='ref' href='#endnote-ref-3'>[3]</a> and making small command-line flag parsing utilities<a class='ref' href='#endnote-ref-4'>[4]</a>. I feel like I'm reimplementing the Go standard library in C++.

So far, it's been an incredible way of learning C++. I'm just getting into smart pointers, but I've already gotten a taste of RAII (which is amazing), templates, move semantics, references vs pointers, and so on. I'm looking forward to learning more advanced topics as I get to them.

<br/>

<a class='endnote' name='endnote-ref-1'>[1]</a> [Thinking big, starting small](https://medium.com/@preetamjinka/thinking-big-starting-small-b0719ac9604)  
<a class='endnote' name='endnote-ref-2'>[2]</a> [FailureDetectors](http://www.cs.yale.edu/homes/aspnes/pinewiki/FailureDetectors.html)  
<a class='endnote' name='endnote-ref-3'>[3]</a> [ip.hpp](https://github.com/Preetam/cpplibs/blob/39a9eee62112edb08bc2a3388e6bd5b5e1cd2242/include/ip.hpp)  
<a class='endnote' name='endnote-ref-4'>[4]</a> [flags.hpp](https://github.com/Preetam/cpplibs/blob/39a9eee62112edb08bc2a3388e6bd5b5e1cd2242/include/flags.hpp)
