---
title: Powered By
date: "2014-04-19"
url: /powered-by
---


A long, *long* time ago, Misframe used to run on Wordpress. That didn't last very long :-). Then, in 2011, I learned enough about Node.js and CouchDB to build my first blogging "platform" from scratch. That was a really fun project. I essentially wrote most of it in one evening during a Thanksgiving break. Occasionally I browse through the [source code](https://github.com/Preetam/Misframe-Platform) to see how horrible I coded back then.

Eventually it became a hassle. I kept tweaking and adding features, and I found myself spending more time coding than writing. I gave up on maintaining my own blogging code and moved to Tumblr.

Tumblr was fine but I had two concerns. First, as before, I found myself writing less. I'm not sure why that happened. I also felt insecure about my content. It was just somewhere "in the cloud." It was also not easily accessible. There's no easy way to get all of my posts in a ZIP file, for example.

I also tried out [Pelican](http://blog.getpelican.com/), a static site generator. It worked fine... sort of. It still felt a little too bloated. I felt like I could write something myself. I had already converted all of my posts to Markdown in order to migrate to Pelican. So, in somewhat of a fit one evening, I sat down and started coding something that I felt would be simple and *elegant*.

On March 4th 2014 at around 1 AM, I finished a simple Go program that ran Misframe. It had less than 150 lines, including the HTML and CSS, and ran on my BeagleBone Black. I've cleaned up the code a little and now, on April 19th 2014, I have the next iteration of this simple program that's still running on my BeagleBone.

There's no fancy database, complicated routing, or any ridiculous dependencies. I have a single binary with a few text files. I'm spending more time writing :-).

The entire source code of this version of Misframe is available on GitHub: [https://github.com/Preetam/misframe](https://github.com/Preetam/misframe)
