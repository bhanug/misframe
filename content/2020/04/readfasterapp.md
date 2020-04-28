---
title: ReadFaster.app
date: "2020-04-28"
---

I have a new project! It's called [ReadFaster.app](https://www.readfaster.app). I'm launching it May 1, 2020.

When I usually talk about my projects, I'm either thinking about starting something or already have something in progress. This time is different: I'm already done!

![ReadFaster.app](/img/2020/rfa.png)

I started this project for [Startup School](https://www.startupschool.org/companies/readfasterapp). I needed something to build and decided to work on something I've wanted for a while.<sup>1</sup> Plus this idea is simple enough to implement that it gave me time to focus on other aspects.

This project was more about the process and execution, and less about the idea itself. It was my first time taking the Startup School course, attending group sessions, getting feedback about the idea, and talking to users... _all before writing a single line of code._ Usually I start with code 🙃. I learned a lot this way, like how to really focus on the _problem_ and the _why_.

I did most of the implementation in the last couple of weeks. It was a nice way of taking advantage of all of the free time I have during this COVID-19 pandemic.

Here are some implementation details. I did a few new things this time:

* Go API
* Preact UI (first time using Preact instead of Mithril.js)
* PostgreSQL (not my own K-V store this time, phew!)
* Docker containers on a DigitalOcean droplet
* CloudFront (CDN for everything, including APIs!)
* Built using GitHub Actions
* Deployed using [Ansible in a GitHub Action](/2019/10/using-ansible-with-github-actions/)

Now that this is all done... I should to get back to reading 😂.

---

1. I was using my old Transverse project for this use case but it was a poor experience.

<!--more-->
