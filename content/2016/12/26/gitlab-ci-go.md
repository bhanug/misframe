---
title: GitLab CI with Go
date: "2016-12-26T23:00:00-05:00"
bestof: true
---

This took a few tries to get right. I have a Go project that uses vendoring
and has a few binaries to save as artifacts. Here's the **.gitlab-ci.yml** file
you should start with.

```yml
image: golang:1.7

variables:
  REPO_NAME: gitlab.com/Preetam/my-project

before_script:
  - go version
  - echo $CI_BUILD_REF
  - echo $CI_PROJECT_DIR

stages:
  - test
  - build

test-project:
  stage: test
  script:
    - mkdir -p $GOPATH/src/$REPO_NAME
    - mv $CI_PROJECT_DIR/* $GOPATH/src/$REPO_NAME
    - cd $GOPATH/src/$REPO_NAME
    - go test $(go list ./... | grep -v /vendor/)

build-project:
  stage: build
  script:
    - mkdir -p $GOPATH/src/$REPO_NAME
    - mv $CI_PROJECT_DIR/* $GOPATH/src/$REPO_NAME/
    - mkdir -p $CI_PROJECT_DIR/artifacts
    - cd $GOPATH/src/$REPO_NAME
    - cd api && go build -o $CI_PROJECT_DIR/artifacts/api && cd ..
    - cd internal-api && go build -o $CI_PROJECT_DIR/artifacts/internal-api && cd ..
    - cd web && go build -o $CI_PROJECT_DIR/artifacts/web && cd ..
  artifacts:
    paths:
      - artifacts/api
      - artifacts/internal-api
      - artifacts/web
```

You should change the `REPO_NAME` variable to point to your own project.

I'm doing a couple of interesting things. First, I'm moving my project files
into my `GOPATH`. That's to make vendored packages work properly. Second, I'm
putting my binaries back in `$CI_PROJECT_DIR` because GitLab doesn't
support artifacts outside of the project directory.

I'm also setting up two stages in a pipeline so artifacts are only created if
the `test` stage passes.

<img src='/img/2016/12/gitlab-ci.png' width=454/>
