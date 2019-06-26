---
title: High availability with libab, health checks, and DNS
date: "2017-03-13T23:45:00-04:00"
---

This weekend project came up from something I thought about at [work](https://www.vividcortex.com).
I need to make a replicated service highly available across availability zones, and for consistency
reasons I need to make sure only one instance of the service is accessed.

I was looking for an external coordination service that...

1. Is highly available
2. Has service health checks
3. Picks *one* healthy service
4. Exposes it with DNS

#3 is the tricky part. It's basically like external leader election. I was looking at [Consul]
(https://www.consul.io/) but I don't think it can do what I'm looking for, so I decided to implement
something myself.

This reminds me of something I did last year with [libab and DigitalOcean's floating IPs]
(https://misfra.me/2016/04/29/ha-with-libab-and-digital-ocean/). As I mentioned in that post, libab
uses a bully algorithm and failure detectors to elect a single node in a cluster. libab's leader
election is used to make sure only one instance is updating DNS records. Otherwise, there may be
some weird race conditions.

Here's what the prototype architecture looks like:

<img src='/img/2017/03/failover.png' width=500/>

1. There are two identical services. I made these simple HTTP servers responding with "200 OK".
2. There's a single process using libab to do the actual health checking. It simply makes sure that
the primary service responds to HTTP requests. This is written in Go, and I have a simplified version
at the end of the post.
3. There are backup instances of the health check process doing failure detection on the elected node.
One of these will take over if the primary fails.
4. The client accessing the services. It sends requests to a CNAME entry pointing to one of the
actual services.
5. DNS service. In this prototype, I used DigitalOcean's DNS and API. The TTL for the CNAME was tiny.
I think it was 5 seconds or something like that. That way the client would connect to the right service
as soon as possible when a failover happens.

Here's a short video of health checking in action. There's only one health check instance running,
and it's running locally. You can see that my `curl` requests, which were happening every two seconds,
don't fail when I shut down the primary service.

<iframe width="560" height="315" src="https://www.youtube.com/embed/2Rtp9DDusio" frameborder="0" allowfullscreen></iframe>

This second video shows libab's leader election in action. There are two instances of libab running
in different data centers, and a third running on my laptop to serve as a tie-breaker.

<iframe width="560" height="315" src="https://www.youtube.com/embed/wSbaBOXbOzA" frameborder="0" allowfullscreen></iframe>

## Go source code

This is a simplified version of the health check program.

```go
package main

import (
	"flag"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Preetam/libab/go/ab"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type tokenSource string

func (s tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{AccessToken: string(s)}, nil
}

var (
	node   *ab.Node
	client *godo.Client
	nodeID int
)

type HealthChecker struct {
	lock       sync.Mutex
	active     bool
	failedOver bool
}

func (h *HealthChecker) OnAppend(node *ab.Node, commit uint64, data string) {}
func (h *HealthChecker) OnCommit(node *ab.Node, commit uint64)              {}
func (h *HealthChecker) LostLeadership(node *ab.Node) {
	log.Println("lost leadership")
	h.lock.Lock()
	h.active = false
	h.lock.Unlock()
}
func (h *HealthChecker) OnLeaderChange(node *ab.Node, leaderID uint64) {
	log.Println("leader is now node", leaderID)
}

func (h *HealthChecker) GainedLeadership(*ab.Node) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.active = true

	go func() {
		for _ = range time.Tick(1 * time.Second) {
			h.lock.Lock()
			if !h.active {
				h.lock.Unlock()
				return
			}
			h.lock.Unlock()

			log.Println("checking health")
			// (Do some health check)
			// if failed
			{
				if h.failedOver {
					log.Println("already failed over. not updating")
					continue
				}
				// (Update DNS)
				h.failedOver = true
			}
			// else
			{
				if h.failedOver {
					log.Println("recovering failed over service")
					// (Update DNS)
					h.failedOver = false
				}
			}
		}
	}()
}

func main() {
	flag.IntVar(&nodeID, "node-id", 0, "node ID")
	addr := flag.String("addr", "", "listen address")
	clusterSize := flag.Int("cluster-size", 3, "cluster size")
	peers := flag.String("peers", "", "comma-separated list of peers")
	token := flag.String("token", "", "DigitalOcean token")
	flag.Parse()

	var err error
	node, err = ab.NewNode(uint64(nodeID), *addr, &HealthChecker{}, *clusterSize)
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
