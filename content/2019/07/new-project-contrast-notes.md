---
title: 'New project: Contrast Notes'
date: "2019-07-15"
---

I'm taking more notes at work these days. They're mostly meeting notes and to-do lists.
I was trying to figure out the best app for me to take notes the way I do. I didn't like
the Notes app on macOS or a text editor, so I decided to write my own Markdown notes app.
It ended up being a great project idea!

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">I worked on a design for my notes app project. <a href="https://t.co/BC2fG29ZKQ">pic.twitter.com/BC2fG29ZKQ</a></p>&mdash; preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/1149722083917615105?ref_src=twsrc%5Etfw">July 12, 2019</a></blockquote>

I call it Contrast and it's available here: https://contrast.site. I'm still working
on the landing page. The code is on GitHub: https://github.com/Preetam/contrast. It uses [Mithril.js](https://mithril.js.org),
runs on [Cloud Run](https://cloud.google.com/run/) (my first serverless app!), and I used
[Figma](https://www.figma.com/) for the design.

It's a very simple app. It runs entirely in the browser (no communication with the server)
and uses local storage for saving. There's no syncing. Instead, there is a "save" button you can use
to save the Markdown contents of a note somewhere on your computer. I don't need syncing or sharing
support because these notes don't last very long. Once I draft something and clean it up, I
post it somewhere else like Slack or GitHub.

Besides actually using it to take notes, it's also an exercise in designing something, building
it, figuring out what doesn't work in practice, and redesigning. The local saving and full
rendered view features came after a few days of use and realizing they were something I really
wanted. I'll continue making small improvements over time, but I am still considering this project
_done_.

<!--more-->
