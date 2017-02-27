---
title: Signed SSH Certificates
date: "2017-02-26T19:45:00-05:00"
---

We've used LDAP at [work](https://www.vividcortex.com/) to manage SSH access for a while.
I won't go into the detailed pros and cons of LDAP for SSH access, but I'll just say that
centralized authentication like that can get really annoying. If your LDAP server goes down,
you can't get into any system. It's happened to us multiple times and it's not fun.

Instead of coming up with clever ways of making it highly available or whatever, we just
decided to get rid of it and switch to a decentralized approach. We went with a basic version
of Facebook's approach that they explained in [Scalable and secure access with SSH](https://code.facebook.com/posts/365787980419535/scalable-and-secure-access-with-ssh/).

Basically...

* There's a CA certificate added to all instances.
* Instances accept any valid certificate signed by the CA.
* Each person has their SSH key signed by the centralized CA.

Instead of having a single SSH authority that manages who has access, we now have a decentralized
system where each instance checks the signature of a certificate and a local revocation list
without contacting anything else.

There are also a couple of other changes:

* We don't create user accounts anymore. Everyone logs in as the same `ec2-user` user.
* We don't use AWS key pairs anymore.

As the Facebook post mentions, we can still keep track of logins by individuals using their SSH
certificate IDs. In **/var/log/secure**:

```
Feb 26 17:53:40 ip-10-10-10-103 sshd[11745]: Accepted publickey for ec2-user from 10.10.1.110 port 58734 ssh2: RSA-CERT ID preetam (serial 2000) CA RSA ...
```

I'm starting to do this for my own servers as well. It'll make key rotations a lot easier.
