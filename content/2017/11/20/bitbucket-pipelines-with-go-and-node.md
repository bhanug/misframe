---
title: Bitbucket Pipelines with Go and Node
date: "2017-11-20T22:30:00-05:00"
---

I wrote [GitLab CI with Go](/2016/12/26/gitlab-ci-go/) last year with
my **.gitlab-ci.yml** for building Go projects on GitLab, but I switched to Bitbucket and
Bitbucket Pipelines so I figured I should post my new setup.

This is how I'm building binaries and assets for [Transverse](https://github.com/Preetam/transverse).
I have two Go binaries, **web** and **metadata**, and a tar.gz with my static assets (images,
templates, CSS, the minified **app.min.js**, etc.). Once my artifacts are built they're uploaded
to S3 so I can fetch them later during deployment.

To get started, you need to grab a couple of things:

* The upload script **s3_upload.py** [here](https://bitbucket.org/awslabs/amazon-s3-bitbucket-pipelines-python/src/d1a2cd2355813b62621d8fedc9e100acf9adb228/s3_upload.py?at=master&fileviewer=file-view-default).
* My **preetamjinka/ci:golang** Dockerfile [here](https://github.com/Preetam/Dockerfiles/blob/5069c3afd98af1ef09f5c71624ccee42391b7a2f/golang/Dockerfile).
I need to use a custom image because I'm building both Go and Node.js stuff in one build. You probably
need to customize this or use your own Dockerfile.

Finally, here's my **bitbucket-pipelines.yml**:

```yaml
image: preetamjinka/ci:golang

pipelines:
  default:
    - step:
        script:
          - PACKAGE_PATH="${GOPATH}/src/bitbucket.org/${BITBUCKET_REPO_OWNER}/${BITBUCKET_REPO_SLUG}"
          - mkdir -pv "${PACKAGE_PATH}"
          - tar -cO --exclude-vcs --exclude=bitbucket-pipelines.yml . | tar -xv -C "${PACKAGE_PATH}"
          - cd "${PACKAGE_PATH}"
          - go test ./...                    # test Go stuff
          - cd web && npm i && npm test      # test Node.js stuff
          - cd "${PACKAGE_PATH}"
          - cd metadata && go build && cd .. # build "metadata" binary
          - cd web && go build && cd ..      # build "web" binary
          - cd web && make all && npm run build && tar czvf assets.tar.gz static templates && cd ..
          - python s3_upload.py "${S3_BUCKET}" ./metadata/metadata "transverse_metadata_${BITBUCKET_COMMIT}"
          - python s3_upload.py "${S3_BUCKET}" ./web/web "transverse_web_${BITBUCKET_COMMIT}"
          - python s3_upload.py "${S3_BUCKET}" ./web/assets.tar.gz "transverse_assets_${BITBUCKET_COMMIT}.tar.gz"
```
