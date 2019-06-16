---
title: Attaching lots of EBS volumes
date: "2016-12-20T22:00:00-05:00"
---

When you attach an EBS volume to an EC2 instance, you'll see a prompt like this
if you're using the console:

[![Attach volume dialog](/img/2016/12/attach-ebs.png)](/img/2016/12/attach-ebs.png)

What if you have already used /dev/sdf to /dev/sdp<sup>1</sup>? You can still use other
names!

[Here](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html) is the AWS
documentation for device naming. I've copied the relevant table below.

[![Available EBS Names](/img/2016/12/ebs-names.png)](/img/2016/12/ebs-names.png)

/dev/sd[f-p] are recommended, but you can use the rest of the letters (except for whatever
is reserved for root).

---

1. If you are using *that* many devices, you're probably doing something wrong...
