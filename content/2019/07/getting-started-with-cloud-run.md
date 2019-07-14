---
title: Getting started with Cloud Run
date: "2019-07-14"
---

[Cloud Run](https://cloud.google.com/run/) is a service on Google Cloud that lets you run
Docker Containers in a serverless way. Billing works a lot like AWS Lambda; you're charged
for CPU and memory in 100 ms increments. Unlike AWS Lambda, you can run _any_ Docker image
that listens on an HTTP port.

Here's a really simple example with a Go program that runs a simple HTTP server:

```
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, world!\n")
	}))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
```

Here's the `Dockerfile` to package that program into a simple image:

```
FROM alpine
COPY hello-world /bin/hello-world
CMD hello-world
```

Cloud Run only seems to work with images hosted on Google Cloud Container Registry.
If you only push to Docker Hub, you need to do a few more things to get your image ready
for Cloud Run. Fortunately you can do the rest from the Google Cloud Shell in your browser.

You need to configure Docker in your Cloud Shell environment and then re-tag your image for GCR.

```sh
# Set up Docker
gcloud auth configure-docker

# Pull image from Docker Hub
docker pull preetamjinka/golang-hello-world

# Re-tag image for Google Cloud Container Registry
docker tag preetamjinka/golang-hello-world gcr.io/my-project-180915/golang-hello-world

# Push
docker push gcr.io/my-project-180915/golang-hello-world
```

Once that's done, you can create your Cloud Run service with just a few clicks and
have a serverless website up and running.

<!--more-->
