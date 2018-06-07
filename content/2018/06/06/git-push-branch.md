---
title: git push-branch
date: "2018-06-06T20:30:00-07:00"
---

One of my top git commands is `push-branch`, which is a custom command that I
configured to push whatever branch I have checked out.

My usual git workflow is this:

1. Checkout a branch: `git checkout -b preetam/my-branch`
2. Make some changes and commit
3. Push this new branch and nothing else:  
With vanilla git: `git push origin preetam/my-branch`  
With push-branch: `git push-branch`

With `push-branch` I don't have to keep typing the branch name.

You can add `push-branch` to git by creating an executable script
named  
`git-push-branch` somewhere in your `PATH` with the following:

```sh
#!/bin/bash
BRANCH=$(git name-rev HEAD 2> /dev/null | awk "{ print \$2 }")
git push origin $BRANCH "$@"
```

The nice thing about git is that if you type in an unknown command like
`git foo`, it'll search your path for a file called `git-foo` and run it.
You can make lots of custom commands like this to save you time, and they
don't even have to be git things. I once had a command to print the weather!
