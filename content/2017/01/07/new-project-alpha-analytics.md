---
title: "New project: Alpha analytics"
date: "2017-01-07T23:00:00-05:00"
---

I have a new project! It's called Alpha. Since the last project I started is called
[Epsilon](2016/11/05/epsilon/) (for <strong>e</strong>vents), I decided to call this Alpha for...
<strong>a</strong>nalytics! Alpha is an extremely simplified version of Google Analytics. I know
very little about website / browser analytics, so building something like this is a great way to
learn.

It's up and running right now. If you're reading this post on the Misframe site, you'll see this
snippet on the bottom of the page:

```js
<script>
(function() {
  var e = document.createElement('script');
  e.type = 'text/javascript';
  e.src = 'https://alpha.infinitynorm.com/t.js';
  e.async = true;
  document.body.appendChild(e);
})();
</script>
```

Here's what I want to get out of working on this project:

- Another use case for Epsilon
	- I'm coming up with query ideas that I didn't think about before. For example, I want a
	  "count distinct" query, but it doesn't exist yet.
- Get better at D3.js. My [visualizations list](https://preetam.github.io/d3-visualizations/) isn't
  very extensive right now.
- Build another web app. Hopefully I can get this thing to a point where others can use it too.

Here's a preview of the web app I'm writing for Alpha:

<blockquote class="twitter-tweet" data-lang="en"><p lang="en" dir="ltr">Progress... <a href="https://t.co/rCmmyw592M">pic.twitter.com/rCmmyw592M</a></p>&mdash; Preetam (@PreetamJinka) <a href="https://twitter.com/PreetamJinka/status/815981434942918658">January 2, 2017</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

It's not great, but it's a start! The time series are generated using an Epsilon query. Because
Alpha stores raw events in Epsilon, I get to slice and dice this data any way I want. Everything
is stored using [lm2](https://github.com/Preetam/lm2).

---

### Note 1

Did you know Googlebot runs JavaScript? I ran the equivalent of this SQL on Epsilon:

```sql
SELECT COUNT(*) FROM events GROUP BY user_agent
```

```js
{
  "data": {
    "query": {
      "columns": [
        {
          "aggregate": "count",
          "name": "_id"
        }
      ],
      "descending": false,
      "group_by": [
        "user_agent"
      ],
      "order_by": [
        "count(_id)"
      ],
      "time_range": {
        "end": 9223372036854775807,
        "start": 0
      }
    },
    "summary": [
      {
        "count(_id)": 44,
        "user_agent": "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
      },
      {
        "count(_id)": 90,
        "user_agent": "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36"
      }
	  // Others that I manually removed since Epsilon doesn't have a regexp filter yet :).
    ]
  }
}
```

### Note 2

Another version of this project started as a way to track my own website usage. I used a browser
extension to inject the tracking code into every page I visited. I ended up finding a security
issue in [VividCortex](https://www.vividcortex.com/) this way 😬 . Make sure you're using
[CSP](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) policies in your applications!
