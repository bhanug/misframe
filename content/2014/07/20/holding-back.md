---
title: Holding back.
date: "2014-07-20"
url: /holding-back
---


(This is basically a brain dump. No revisions or drafts!)

For many people, their GitHub profile is their coding portfolio. I wanted mine to be like a portfolio too. I guess you can say the same thing about Misframe. I don't write about *everything* I think about. Sometimes I let things simmer around in my head, and when I'm not feeling lazy I just collect those thoughts together and digest it into a post. So in the end, I'm more or less left with sh** that roughly makes sense.

The really annoying downside is that this arbitrary quality threshold or whatever filters out a lot of good stuff. Here's a little PHP script I found that takes a disk image of Ubuntu 11.10 (so you know it's from way back :P) and provisions a KVM:

https://gist.github.com/PreetamJinka/42fd851980ef9a04e7ab

Now, this one was saved as "deploy.php". I also had one called "crudeDeploy.php". I glanced at it and it's basically the "deploy.php" version except the parameters are hardcoded. It's funny because the root password was set to an expletive. Right now, I'm going, "that's AWESOME." That would've never made it to my GitHub, but I know I probably wrote that because I was super pissed one evening and I wanted the darn thing to work.

What the actual thing was doing is pretty interesting. First it creates a logical volume and copies the disk image over. The disk image is compacted, so the rootfs was around like 2 GB and there was no swap partition. The script sets up loop devices and mounts filesystems. It then resizes the rootfs to the size of the logical volume, then writes sets up a new root password, adds a swap partition, and all that jazz. Oh, and it also sets up the networking config within the VM so the IP config is statically set on boot. There's also some stuff that goes on in the host machine to create a shell environment for the VM.

Did I mention all of this happens in around a minute? DigitalOcean never really impressed me ;).

Now, all of this was an internal thing. I never showed this to anyone. I was just messing around with stuff during my last two years of high school.

I don't see myself doing stuff like this right now. It's not "portfolio" quality. That really sucks. There's lots of interesting stuff in that single file.

Something like this, which reads a partition table off a file and modifies a block device's partition table to match, is awesome. I don't even know how long it took me to figure this out.

```
cat ./partitionmaps/$plan | sfdisk /dev/vps/$kvmID --force
```

I'm gonna start using GitHub like I was doing a few years ago: writing stupidly simple, sh***y code that does whatever I wanted it to do at the time. I'm also going to try to stop thinking so much before I write (code and for Misframe). Having a mental threshold slows me down a lot.
