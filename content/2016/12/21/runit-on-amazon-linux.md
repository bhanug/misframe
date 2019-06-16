---
title: runit on Amazon Linux with BusyBox
date: "2016-12-20T22:15:00-05:00"
---

[runit](https://smarden.org/runit/) is a great init scheme that you can use
along with your existing init system. It makes it really easy to create
service configurations for APIs and other long-running programs.

We're using runit with Fedora at the moment, but we're moving to Amazon Linux
and need to bring runit with us. Getting runit on Amazon Linux is a little
complicated. It's not something you can simply `yum install runit` since it's
not a part of the official Amazon Linux package list<sup>1</sup>.
[EPEL](https://fedoraproject.org/wiki/EPEL) doesn't have it either.

I figured I had two options left:

1. Compile it myself
2. Use https://github.com/imeyer/runit-rpm and build an RPM

Fortunately, I found ["Running runit on Amazon Linux AMI"](https://evasive.ru/50a3904206c52447aa1fa5d90a8382a3.html)
and that offered a great workaround: use BusyBox! I had no idea BusyBox included
runit. I guess that says something about how useful BusyBox is, and how minimal runit is.

The first thing you need to do is install BusyBox.

```sh
yum install busybox
```

Then, create some symlinks (as root) to let BusyBox handle runit programs.

```sh
busybox --list | awk '/runsv|chpst|svlog|^sv$/' | \
    xargs -I{} ln -sv /sbin/busybox /sbin/{}
```

After that, you can set up your runit services as usual.

---

1. Here is the Amazon Linux AMI 2016.09 packages list:  
   https://aws.amazon.com/amazon-linux-ami/2016.09-packages/
