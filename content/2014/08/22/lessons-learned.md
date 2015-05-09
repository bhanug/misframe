---
title: Lessons learned.
date: "2014-08-22"
url: /lessons-learned
---


I've finished another summer of interning. Here are some incredibly useful
"lessons learned." Some of these came from hours of frustration. Some are related to others.
And by the way, just because you can read this list does not mean you're "learning" these lessons. ;)

1. Intuition is incredibly powerful.
2. Having the right tools is the difference between trying for hours or days and failing, and getting the answer within minutes.
3. Tools can either empower you or burden you. It's critical to find ones that work and if you can't, build your own!
4. If you don't know what you're doing, you're most likely going to waste your time.
5. Ask for help if you're stuck, but remember that even making a tiny bit of progress is still progress.
6. Nothing is perfect, and most things aren't close to perfect. But that's okay!
7. It's impossible to be good at everything if you're working in a team. People get good at what they do, and not everyone's doing the same thing.
8. Assumptions can be dangerous. Ask yourself, "is this correct?" Don't answer that question until you're sure.
9. You get better every day even if you don't notice it.
10. You're probably right more often than you think you are.

---

Some of the public things I've worked on...

1. Worked on replacing LevelDB in [InfluxDB](http://influxdb.org/) with BoltDB. [Link.](https://github.com/VividCortex/influxdb/tree/bolt-storage-engine) And no, you probably don't want to use a B-tree storage engine for time series data ;). Just working on this was an interesting project. Some lessons learned here ;).
2. Created [Trace](https://github.com/VividCortex/trace). You can see where I sprinkled those lines into the InfluxDB code to diagnose race conditions.
3. A few annoying compatibility things to [golibpcap](https://github.com/VividCortex/golibpcap/commits/master). For some reason a few structs are defined as nested unions or whatever on the recent versions of Ubuntu. Go doesn't like that. Go also doesn't like functions defined as preprocessor macros (OS X does this).
