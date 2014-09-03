title: Streaming HTTP
date: 2014-09-02 22:45:00
url: streaming-http

I just came back from this month's
[PyCHO](https://wiki.python.org/moin/CharlottesvillePythonUserGroup) meetup,
and one of the things that we discussed was how to handle many simultaneous
HTTP clients that send data in streams.

Someone had a project, written in Python of course, that accepted large file
uploads from clients, computed a hash, and sent it to S3. The goal is to
do everything concurrently with minimal resource usage.

I'm not a Python programmer, but as a Go programmer this is very simple to do.
You don't have to worry about threads, processes, coroutines, futures, or
any other concurrency primitives. In fact, the following example doesn't even
expose goroutines directly.

This example is an extremely simplified version of the project that I whipped
up in a couple of minutes. I wouldn't even call it "correct," but it's enough
to get the point across.

```go
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
)

// HD5Handler computes an MD5 checksum of the HTTP
// body and sends it as a response.
func MD5Handler(w http.ResponseWriter, req *http.Request) {
	hash := md5.New()

	// Copy data from req.Body (an io.Reader) to
	// hash (an io.Writer).
	io.Copy(hash, req.Body)

	req.Body.Close()

	// Write a response
	fmt.Fprintf(w, "%x", hash.Sum(nil))
}

func main() {
	// Route "/" to the MD5Handler
	http.HandleFunc("/", MD5Handler)

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```

And here is some sample output.

```
$ go run main.go
```

```
$ # This creates a 100 MB file of random data
$ dd if=/dev/urandom of=test bs=1M count=100
100+0 records in
100+0 records out
104857600 bytes (105 MB) copied, 7.73638 s, 13.6 MB/s

$ time curl -i -X POST -d @test localhost:8080
HTTP/1.1 100 Continue

HTTP/1.1 200 OK
Date: Wed, 03 Sep 2014 02:28:12 GMT
Content-Length: 32
Content-Type: text/plain; charset=utf-8

97ccb418d79d7f69c71d145bd88c643c
real	0m0.360s
user	0m0.106s
sys	0m0.086s
```
