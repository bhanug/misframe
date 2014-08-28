title: Packet Sniffing.
date: 2014-08-27 00:00:00
url: packet-sniffing
draft: true

If you search for "packet sniffing" or "how to sniff packets" on
Google, chances are you'll find many results on the general concepts
behind packet sniffing and software commonly used to sniff packets.

If you want to do some packet sniffing of your own, you can get Wireshark
or program something of your own using libpcap. It's easy enough, and
there are many resources online that can help you do this. But eventually
you might ask yourself, as I did, how do these tools actually work?


Example
---

```go
package main

import (
	"log"
	"fmt"
	"syscall"
)

// htons converts an int (assuming 16-bit) to network order (big endian).
func htons(n int) int {
	return int(int16(byte(n))<<8 | int16(byte(n>>8)))
}

func main() {
	fd, err := syscall.Socket(syscall.AF_PACKET,
		syscall.SOCK_RAW, htons(syscall.ETH_P_ALL))

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 65536)

	for {
		n, _, err := syscall.Recvfrom(fd, buf, 0)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%x\n", buf[:n])
	}
}
```
