title: Router-on-a-stick
date: 2013-11-27 17:14:40
url: router-on-a-stick

(This is an old post from July 2012.)

![](http://media.tumblr.com/6ab3187ba70814940579f613dd54e156/tumblr_inline_mwy0cl4GTl1rs73cz.jpg)Attribution: <a href='http://www.flickr.com/photos/pdstahl/3903808739/'>Patrick Stahl</a>

This is just a quick configuration example to setup a “router-on-a-stick.” I’m using Vyatta for the router and a Brocade switch. I’ve changed the real addresses.

<h2>Brocade switch configuration</h2>

Ethernet 16 is the *trunk* port, in the Cisco lingo. It’s hooked up to the Vyatta router. It has to be tagged because we need to tag each frame with a VLAN ID so the switch knows where to send it. Ethernet 43 is an untagged (aka *access*) port which is connected to a server.

<pre>vlan 20 by port
 tagged ethe 16                                                   
 untagged ethe 43 
</pre>

<h2>Vyatta configuration</h2>

eth1 is connected to my provider. eth0 is going back to the switch. vif 20 refers to VLAN 20, which we already defined on the switch.

<pre>interfaces {
    ethernet eth0 {
        hw-id 90:e2:ba:1b:5a:54
        vif 20 {
            address 10.1.5.129/29
            address 2001:db8:5341:31::/64
        }
    }
    ethernet eth1 {
        address 10.1.5.42/29
        address 2001:db8:5341:8::2/126
        hw-id 00:1a:4d:53:f5:d4
    }
    loopback lo {
    }
}
protocols {
    static {
        route 0.0.0.0/0 {
            next-hop 10.1.5.41 {
            }
        }
        route6 ::/0 {
            next-hop 2001:db8:5341:8::1 {
            }
        }
    }
}
</pre>

<p>Not really a useful blog post, but I just wanted to post <em>something</em>. Vyatta is great so far!</p>

