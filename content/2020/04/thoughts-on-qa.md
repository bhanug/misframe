---
title: Thoughts on QA
date: "2020-04-27"
---

A former colleague used to tell me that when it comes to QA, there's no replacement for someone manually clicking around and trying different things in your product. After thinking about that for a few years, I agree with it more than ever.

Automated unit tests, integration tests, end-to-end (E2E) tests are all important. They make sure things work as expected. They catch any breaking changes you may accidentally introduce. You gotta have them. But they're only as good as your test cases. You can cover a lot of ground with just a few test cases, but getting to 100% coverage often requires unreasonable amounts of time. Not worth it. That's OK—done is better than perfect.

You know what QA approach I think has a great return on investment? Just clicking around and trying different inputs. You don't need to think so much about what to test. Just play around with things. It's often very clear when things don't work as expected.

You can start by opening up your documentation and following what it says. This is how users start using your new features anyway, right? It's important to make sure those cases work flawlessly.

Demo prep is also a great way to find issues outside of automated tests. Those are usually the first times I try out a feature with an actual use case or workflow instead of something in isolation. Plus, the issues I find aren't just bugs that slipped past automated tests but also usability issues, things that make workflow implementation really hard or impossible.

The final test for quality comes from having a person trying out what you built. Better for you to run that test before a customer.

<!--more-->
