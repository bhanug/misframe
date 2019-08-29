---
title: 'Deploying projects to Cloud Run using GitHub Actions'
date: "2019-08-27"
---

Over the past month I made some small but significant changes to my [notes app project](/2019/07/new-project-contrast-notes/).
It was always a Docker container on [Google Cloud Run](https://cloud.google.com/run/), but now it uses multi-stage
Docker builds, is built using [GitHub Actions](https://github.com/features/actions) workflows, and is deployed automatically to Google Cloud!

<!--more-->

## Multi-stage builds

My project is a Docker container with a single-page app and a tiny Go web server. My original Docker image simply
contained artifacts I built in my regular development environment that I copied into the image. The `Dockerfile` was this:

```Dockerfile
FROM alpine
COPY ./serve/serve /bin/serve
COPY ./web/static /static/
CMD serve
```

The downside to this approach was that I needed all of the dependencies to build those artifacts available outside of Docker.
I couldn't build it easily on CI/CD systems because I first had to have Go, Node.js, NPM, and anything else I need to build
the artifacts.

Later I learned about multi-stage builds in Docker, which let you build those artifacts from within Docker itself. This means
you don't need any other dependencies besides Docker to build my project image. This is what my `Dockerfile` became:

```Dockerfile
FROM golang:alpine AS build-go

COPY ./serve /serve

RUN cd /serve && go build

FROM node AS build-node

COPY web /web

RUN cd /web/static && npm i
RUN cd /web/static && npm run build
RUN mkdir -p /web/static/css
RUN cd /web && node ./node_modules/clean-css-cli/bin/cleancss ./css/style.css -o ./static/css/style.min.css

FROM alpine
COPY --from=build-go /serve/serve /bin/serve
RUN chmod +x /bin/serve
COPY --from=build-node ./web/static /static/
CMD /bin/serve
```

The first section builds the Go binary, the second section builds the UI artifacts, and the last section copies the
binary and artifacts from the previous images into the final, minimal, deployable image.

All that is straightforward on its own, but the fact that I could now build my images using only `docker build`
meant that I could start building them using GitHub Actions!

## GitHub Actions

TODO

![GitHub Actions screenshot](/img/2019/08/github-actions-result.png)

[GitHub Workflow](https://github.com/Preetam/contrast/blob/master/.github/workflows/push.yml)
