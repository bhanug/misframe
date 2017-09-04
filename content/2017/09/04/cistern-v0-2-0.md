---
title: Cistern v0.2.0
date: "2017-09-04T12:35:00-04:00"
---

<img src='/img/2017/09/cistern-v0.2.0.png' width=400/>

It's out! You can go download a binary on the GitHub [release page](https://github.com/Cistern/cistern/releases/tag/v0.2.0)
and follow the [Getting Started](https://cistern.github.io/docs/#getting-started) instructions.

As mentioned in the [previous post](/2017/08/31/whats-coming-in-cistern-v020/), the major features are

* Query language
* UI
* Generic JSON CloudWatch Logs support

There's also a new website for the project at [cistern.github.io](https://cistern.github.io/) with
some [documentation](https://cistern.github.io/docs/) about how it works. There will be more
documentation coming soon.

## Up next

Some things I want to see in the next version are

* More filter operators (greater than, less than, regular expressions)
* GROUP BYs with functions. That way you can GROUP BY 5xx status codes, for example.
* Cleaner logging from the binary
* Error reporting from the UI
