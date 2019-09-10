---
title: Deploying projects to Cloud Run using GitHub Actions
date: "2019-09-09"
twitter_card:
  description: "Deploying projects to Cloud Run using GitHub Actions"
  image: "https://misfra.me/img/2019/08/github-actions-result.png"
---

Over the past month I made some small but significant changes to my notes app project called [Contrast](/2019/07/new-project-contrast-notes/).
It was always deployed as a Docker container on [Google Cloud Run](https://cloud.google.com/run/) and that didn't change,
but the process of getting it there is now completely different. The build process now uses multi-stage Docker builds, is built using
[GitHub Actions](https://github.com/features/actions) workflows, and is deployed automatically on pushes.

<!--more-->

## Multi-stage builds

My project is a Docker container with a single-page app and a tiny Go web server. My original Docker image simply
contained artifacts I built using Node and Go on my laptop that I then copied into the image. The `Dockerfile` was this:

```Dockerfile
FROM alpine
COPY ./serve/serve /bin/serve
COPY ./web/static /static/
CMD serve
```

The downside to this approach is that I need all of the dependencies to build those artifacts available outside of Docker.
I couldn't build it easily on CI/CD systems because I first had to have Go, Node.js, NPM, and anything else I need to build
the artifacts.

Later I learned about multi-stage builds in Docker, which let you build those artifacts from within Docker itself and carry
them over into a final image without having the dependencies to build them in the final image. This is what my `Dockerfile` became:

```Dockerfile
### Build Go server binary
FROM golang:alpine AS build-go

COPY ./serve /serve

RUN cd /serve && go build

### Build UI assets
FROM node AS build-node

COPY web /web

RUN cd /web/static && npm i
RUN cd /web/static && npm run build
RUN mkdir -p /web/static/css
RUN cd /web && node ./node_modules/clean-css-cli/bin/cleancss ./css/style.css -o ./static/css/style.min.css

### Prepare final image
FROM alpine

# Copy Go binary
COPY --from=build-go /serve/serve /bin/serve
RUN chmod +x /bin/serve

# Copy UI assets
COPY --from=build-node ./web/static /static/

CMD /bin/serve
```

The first section builds the Go binary, the second section builds the UI artifacts, and the last section copies the
binary and artifacts from the previous images into the final, minimal, deployable image.

That was an improvement but not significant on its own. The main benefit is that now my Docker image is built using only
`docker build`, and I can build it using GitHub Actions!

## GitHub Actions

[GitHub Actions](https://github.com/features/actions) is the new CI/CD system on GitHub. With a few minutes (but a lot
of tries 😅) I managed to get my project's Docker image built, pushed to Google Container Registry, and use the
available Google Cloud Action to deploy my project to Cloud Run on every push.

![GitHub Actions screenshot](/img/2019/08/github-actions-result.png)

The full workflow YAML is here: [GitHub Workflow](https://github.com/Preetam/contrast/blob/217e812fac24f0922425fa00785039f0aed4f8ad/.github/workflows/push.yml)

I don't have logic for PRs (so PR branches get deployed too!) but that requires only a little bit of tweaking.
Overall, this is a great starting point that I will keep building on for this project and others.

---

_Acknowledgements:_ Thanks to [Chaitanya](https://kchaitanya.com/) for sending me helpful pointers.
