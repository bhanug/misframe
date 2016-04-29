---
title: High Availability with libab and DigitalOcean
date: "2016-04-29T01:56:15.193-04:00"
---

I've been using DigitalOcean more lately and reading up on [Floating IPs](https://www.digitalocean.com/company/blog/floating-ips-start-architecting-your-applications-for-high-availability/).
Floating IPs are a lot like Elastic IPs on AWS. You can allocate a floating IP and assign it to any
droplet within the same data center. To redirect traffic to another droplet, you simply reassign
the IP. This reassignment isn't automatic. When the droplet with the floating IP fails, you'll have
to reassign the IP through the DO interface or use the API.

Then I thought of something interesting: using [libab](https://github.com/Preetam/libab)'s leader
election and the client callbacks to assign a floating IP.

libab allows you to set callback for certain events. One is gaining leadership. When the current
node becomes a leader, libab will call the `gained_leadership` function that you provide. When that
happens you can set the floating IP to point to the droplet the leader node is running on.

## Bully algorithm

libab uses a [bully algorithm](https://en.wikipedia.org/wiki/Bully_algorithm) to elect a leader.
Of all possible nodes, only the one with the lowest ID will become the leader. In the following
diagram, that's Node 1.

![Architecture diagram](/img/2016/04/do-ha.png)

A nice property of the bully algorithm is that you can be sure some nodes will never be elected.
That's because you always need a majority to elect a leader, and the leader is the one with the
lowest ID. In the diagram above, Node 3 can never become the leader (unless you add more nodes).
That's why I put it in another data center. Since it won't be a leader, it won't be assigned the
floating IP, so I don't need to worry about it being in the same data center.

This method is inspired by [CARP](https://en.wikipedia.org/wiki/Common_Address_Redundancy_Protocol).
I used CARP when I did web hosting to setup redundant firewalls using OpenBSD. It uses a similar
bully algorithm approach to assign an IP, but in that case it's the gateway IP for a network.

## Testing it

To test this out, I wrote a small Go program that uses the libab cgo bindings and the DigitalOcean
API. I was able to correctly reassign a floating IP when I terminated the program and the droplet
itself.

The only problem I noticed was that it took a while for the IP to get rerouted. I kept `ping`
running and I noticed the floating IP was unresponsive to pings for 300 seconds during the
transition. That's quite unreasonable for a HA setup, but I'm sure it's going to be much better.

## Another approach to HA

Vultr, a DigitalOcean clone operated by Choopa, also offers a floating IP feature, but they take it
to the next level. They allow you to [use BGP to setup HA](https://www.vultr.com/docs/high-availability-on-vultr-with-floating-ip-and-bgp).
This approach doesn't need a failure detector since you can use anycast to send traffic to an IP
to multiple instances at the same time, but that's beyond the scope of this post.

---

## Go source code

Again, this program doesn't do much since it's a test and it just runs a libab node. I can imagine
this being a part of another web server or some important network service that you'd want to make
highly accessible. It's nice to finally use libab for something useful!

```go
package main

import (
	"flag"
	"log"
	"strings"

	"github.com/Preetam/libab/go/ab"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type tokenSource string

func (s tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: string(s)}, nil
}

var node *ab.Node
var client *godo.Client

var dropletID int
var floatingIP string

type handler struct{}

func (h handler) OnAppend(round uint64, commit uint64, data string) {}
func (h handler) OnCommit(round uint64, commit uint64)              {}
func (h handler) LostLeadership()                                   {}
func (h handler) OnLeaderChange(leaderID uint64)                    {}

func (h handler) GainedLeadership() {
	client.FloatingIPActions.Assign(floatingIP, dropletID)
}

func main() {
	flag.IntVar(&dropletID, "droplet-id", 0, "droplet ID")
	addr := flag.String("addr", "", "listen address")
	clusterSize := flag.Int("cluster-size", 3, "cluster size")
	peers := flag.String("peers", "", "comma-separated list of peers")
	flag.StringVar(&floatingIP, "floating-ip", "", "floating IP address")
	token := flag.String("token", "", "DigitalOcean token")
	flag.Parse()

	var err error
	node, err = ab.NewNode(uint64(dropletID), *addr, handler{}, *clusterSize)
	if err != nil {
		log.Fatal(err)
	}
	for _, peer := range strings.Split(*peers, ",") {
		node.AddPeer(peer)
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext,
		tokenSource(*token))
	client = godo.NewClient(oauthClient)
	node.Run()
}
```
