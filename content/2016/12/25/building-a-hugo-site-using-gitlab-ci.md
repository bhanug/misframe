---
title: Building a Hugo site using GitLab CI
date: "2016-12-25T15:00:00-05:00"
---

Using GitLab CI and a Hugo Docker image, you can build static sites automatically and create
tar.gz artifacts to deploy straight to production. GitLab CI can use images from various registries,
including Docker Hub. I'm using the [jojomi/hugo](https://hub.docker.com/r/jojomi/hugo/builds/bhtpbmfinvukrupmyq49j6a/)
Docker image.

Here's the **.gitlab-ci.yml** file you need:

```yml
image: jojomi/hugo:latest

stages:
  - build

build-site:
  stage: build
  script:
    - hugo
    - tar -cvzf site.tar.gz public
  artifacts:
    paths:
      - site.tar.gz
```
