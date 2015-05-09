---
title: Break-in
date: "2015-02-08"
url: /break-in
summary: Someone broke into my server.
---

Someone broke into my server.

I was at beSwarm yesterday with my "social networking" setup.
<blockquote class="twitter-tweet" lang="en"><p>Social networking! <a href="http://t.co/fdApIwlKyy">pic.twitter.com/fdApIwlKyy</a></p>&mdash; Preetam Jinka (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/564090574009929728">February 7, 2015</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

I was demoing [Cistern](http://preetamjinka.github.io/cistern/) in some form. Cistern doesn't expose much to the user right now since most of my time was spent on very core features. So, what most people usually saw was the terminal log output. It's still a little interesting because you can see it do some basic host discovery using SNMP, and it prints flow data as it arrives.

It looks like this. Sorry about the wrapping.
![](http://static.misfra.me/images/posts/break-in/cistern-log.png)

I grabbed that screenshot as I was demoing it to someone. It's live flow data. If you look carefully, you'll see that most of the lines show my blog's IP sending UDP packets to some IP's port 22. This only a sample (1 in 1024, in fact) of the packets. Therefore, there are lots of packets going out.

My blog, which runs on a BeagleBone Black, does not produce this kind of traffic. Something's up.

I grabbed my laptop and went aside to check up on things. I logged in, and yep. The load average was about 4 on this single core ARM machine and there were weird processes running as the "debian" user and taking up lots of CPU time. I *know* I don't run anything as that user. I checked tcpdump and there were lots of IRC packets going over the wire. I knew immediately that some script kiddie got in and made my server part of an IRC controlled botnet.

I took care of the issue. How did this happen? I made a silly mistake. The [Debian image](http://elinux.org/BeagleBoardDebian) for the BeagleBone Black comes set up with a "debian" user with the default password "temppwd". I always use root with key-based authentication, so I forgot about this user. I apparently did not change the password. Leaving a default combination like that on a publicly accessible server is not good.

By the way, the script they ran is at the following address: [http://mui3.ucoz.com/maxx.txt](http://mui3.ucoz.com/maxx.txt). Notice how it disguises itself (`my @ps = . . .`).

----

This made me upset, but it also made me excited. I saw, for myself, that my tool is useful. It doesn't even do much right now but it has helped me already. I'm excited to think that this can be automated and become a valuable tool. Someone asked me yesterday if I had plans to sell this and I said, "no, it's all open source." It'll stay open. My plan is to disrupt network monitoring by making the most technically advanced software possible and keeping it completely open.

![](http://static.misfra.me/images/posts/break-in/plots.png)

Something awesome is coming.
