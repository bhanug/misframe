---
title: Checking disk activity using iostat
date: "2016-12-28T23:30:00-05:00"
---

I often use a [monitoring system](https://www.vividcortex.com/) to look at disk activity,
but sometimes it's nice to have a CLI tool to get stats in a different format. I use
iostat(1) for that. iostat has a bunch of different options, but I usually stick to
`iostat -dxy`.

Here's the description for each option (from the man page):

```
-d     Display the device utilization report.
-x     Display extended statistics.
-y     Omit first report with statistics since system boot, if displaying multiple records at given interval.
```

Example output:

```txt
$ iostat -dxy 1
Linux 4.4.0-53-generic	 	12/28/2016 	_x86_64_	(1 CPU)

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00    9.90     0.00    39.60     8.00     0.00    0.00    0.00    0.00   0.00   0.00

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00    9.90     0.00    39.60     8.00     0.00    0.00    0.00    0.00   0.00   0.00

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00   10.10     0.00    40.40     8.00     0.00    0.00    0.00    0.00   0.00   0.00

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00    20.79    0.00   11.88     0.00   130.69    22.00     0.00    0.00    0.00    0.00   0.00   0.00

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00   83.84     0.00   420.20    10.02     0.23    2.80    0.00    2.80   0.10   0.81

Device:         rrqm/s   wrqm/s     r/s     w/s    rkB/s    wkB/s avgrq-sz avgqu-sz   await r_await w_await  svctm  %util
vda               0.00     0.00    0.00   10.00     0.00    40.00     8.00     0.00    0.00    0.00    0.00   0.00   0.00
```

The `1` in the example is the interval. It's optional, but if you don't provide it, iostat will
show you stats since system boot, which isn't very useful. You can also pass an optional count
after the interval. For example, `iostat -dxy 3 2` will print 3-second averages twice and exit.

By the way, if you want to read up on `%util` (the last column), check out this blog post:
[Why %util number from iostat is meaningless for MySQL capacity planning]
(https://www.percona.com/blog/2014/06/25/why-util-number-from-iostat-is-meaningless-for-mysql-capacity-planning/).
The comments are good too.
