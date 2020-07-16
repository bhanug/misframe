---
title: Transverse Open Beta
date: "2017-11-13T21:30:00-05:00"
---

It's ready!

Two months ago I [introduced Transverse](/2017/09/12/introducing-transverse/), my app
to forecast goal progress. Everything you saw on that page was just a mockup. Now,
after actually building it and getting it in front of a few people, it's ready for
everyone! Note: **this is an open beta**. Stuff is messy. It's kind of ugly.
But it works.

I've been using it every day. Here's a real screenshot of my goals list:

[![Transverse goals list](/img/2017/11/transverse-goals.png)](/img/2017/11/transverse-goals.png)

It was *really* helpful to track my reading! I finished a record number of books this year.

Here's another real screenshot of my squat goal and its forecast.

[![Transverse squat goal](/img/2017/11/transverse-squat-goal.png)](/img/2017/11/transverse-squat-goal.png)

Finally, I'm going to share what every new user is getting in their registration email.
It's the best description of how Transverse works right now, and lists some of the
issues that I'm going to work on next.

<br>

---

<br>

### Transverse Beta Registration Notes

<p>Thank you for your interest and signing up! I'm really excited to hear what you think of my app.</p>

<h4>How it works &amp; best practices</h4>

<p>Transverse uses daily time series points to determine a trend and forecast future values. It's just drawing a line.</p>

<p>You enter data as a CSV (or TSV; the parser does some delimiter detection). Often times you can just copy+paste stuff out of a spreadsheet. Example:</p>

<pre>
11/1/2017, 10
11/2/2017, 11
11/3/2017, 12
</pre>

<p>You can enter data in manually or use the "quick add" form.</p>

<p>It's OK to skip days. Transverse will fill in gaps, but it uses linear interpolation so depending on your units, you have to manually enter in "0" values.</p>

<p>Transverse only works with daily points. It is also tuned for short-term forecasts (within a month). Anything more is unsupported for now. Transverse will also not generate forecasts longer than the period of time you've added data for.</p>

<p>Need ideas for goals? Here's what I use:</p>

<ul>
  <li>Progress in a book. My target is the total number of pages in the book and each point is the page number I've reached.</li>
  <li>Typing speed. My target is a WPM rate. I usually enter data in daily. I don't fill in gaps.</li>
  <li>Gaming time. I keep track of how long I play a video game every day and enter the duration in. If I skip a day, I need to add a "0" point.</li>
</ul>

<h4>Known issues</h4>

<p>This version is feature complete (for the most part) but lacks polish. Here are things you may run into:</p>

<ul>
  <li>If you want to add a password, you need to first log in with the token method and set a password in your profile.</li>
  <li>Errors aren't very clear when you're setting a password. Make sure you're using at least 8 characters.</li>
  <li>Sometimes there's a forecast even when you hit your goal value.</li>
  <li>You should pretend that the forecast bands (the two outer lines) don't mean anything. They'll mean something eventually, but that requires some intense math that will take a while.</li>
</ul>

<h4>Let me know what you think</h4>

<p>Please let me know what you think! Tell me how you're using it. Share your forecasts. Tell me what works well. Tell me what annoys you. My favorite use of Transverse, tracking book progress, came from my first user.</p>
