---
title: Go projects and the BSD license
date: "2016-02-28T22:40:21.130Z"
---

I relicensed many of my projects on GitHub to use the BSD license a little over a year ago. They
used to be under the MIT license because I started using GitHub when I was a Node user, and most
projects in that ecosystem use the MIT license.

I switched because of this clause in the BSD license:

> 2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation and/or
other materials provided with the distribution.

This means I should get credit whether my code is distributed in its original source code form or if
it's included in other software in binary form. It may be a little annoying to make sure these
copyright notices are included with every binary, but considering how much work work goes into these
projects I think it's fair to require a small notice when they are used. I simply want credit for my
work even if it's not in source code form.

## Go and dependency licenses

The Go project also uses a BSD license so *every* distributed Go binary should include the Go
copyright notice. Because imports get added to Go binaries you also need to include copyright
notices for any dependencies that use the BSD license. I think this is where things can get really
complicated. Your dependencies may include other projects that are BSD licensed, and if you're not
careful you may forget to include the appropriate copyright notices when you distribute binaries.

If you're familiar with `npm` you know that each module may require a ridiculous number of
dependencies. Just look at what gets installed for the `request` module:

```
   ∂ tmp: npm install request
/private/tmp
└─┬ request@2.69.0 
  ├── aws-sign2@0.6.0 
  ├─┬ aws4@1.2.1 
  │ └── lru-cache@2.7.3 
  ├─┬ bl@1.0.3 
  │ └─┬ readable-stream@2.0.5 
  │   ├── core-util-is@1.0.2 
  │   ├── inherits@2.0.1 
  │   ├── isarray@0.0.1 
  │   ├── process-nextick-args@1.0.6 
  │   ├── string_decoder@0.10.31 
  │   └── util-deprecate@1.0.2 
  ├── caseless@0.11.0 
  ├─┬ combined-stream@1.0.5 
  │ └── delayed-stream@1.0.0 
  ├── extend@3.0.0 
  ├── forever-agent@0.6.1 
  ├─┬ form-data@1.0.0-rc3 
  │ └── async@1.5.2 
  ├─┬ har-validator@2.0.6 
  │ ├─┬ chalk@1.1.1 
  │ │ ├─┬ ansi-styles@2.2.0 
  │ │ │ └── color-convert@1.0.0 
  │ │ ├── escape-string-regexp@1.0.5 
  │ │ ├─┬ has-ansi@2.0.0 
  │ │ │ └── ansi-regex@2.0.0 
  │ │ ├── strip-ansi@3.0.1 
  │ │ └── supports-color@2.0.0 
  │ ├─┬ commander@2.9.0 
  │ │ └── graceful-readlink@1.0.1 
  │ ├─┬ is-my-json-valid@2.13.1 
  │ │ ├── generate-function@2.0.0 
  │ │ ├─┬ generate-object-property@1.2.0 
  │ │ │ └── is-property@1.0.2 
  │ │ ├── jsonpointer@2.0.0 
  │ │ └── xtend@4.0.1 
  │ └─┬ pinkie-promise@2.0.0 
  │   └── pinkie@2.0.4 
  ├─┬ hawk@3.1.3 
  │ ├── boom@2.10.1 
  │ ├── cryptiles@2.0.5 
  │ ├── hoek@2.16.3 
  │ └── sntp@1.0.9 
  ├─┬ http-signature@1.1.1 
  │ ├── assert-plus@0.2.0 
  │ ├─┬ jsprim@1.2.2 
  │ │ ├── extsprintf@1.0.2 
  │ │ ├── json-schema@0.2.2 
  │ │ └── verror@1.3.6 
  │ └─┬ sshpk@1.7.4 
  │   ├── asn1@0.2.3 
  │   ├─┬ dashdash@1.13.0 
  │   │ └── assert-plus@1.0.0 
  │   ├── ecc-jsbn@0.1.1 
  │   ├── jodid25519@1.0.2 
  │   ├── jsbn@0.1.0 
  │   └── tweetnacl@0.14.1 
  ├── is-typedarray@1.0.0 
  ├── isstream@0.1.2 
  ├── json-stringify-safe@5.0.1 
  ├─┬ mime-types@2.1.10 
  │ └── mime-db@1.22.0 
  ├── node-uuid@1.4.7 
  ├── oauth-sign@0.8.1 
  ├── qs@6.0.2 
  ├── stringstream@0.0.5 
  ├── tough-cookie@2.2.1 
  └── tunnel-agent@0.4.2
  ```

Imagine if these were Go packages and each one was BSD licensed. That's a *lot* of copyright notices
to include, and I wouldn't be surprised if you didn't include all of them!

I noticed that some Go projects don't list the licenses of their dependencies. That makes it really
hard for people to build and distribute binaries. Getting dependencies for a project may be as easy
as running `go get`, and building a binary is another `go build` after that, but distributing that
binary isn't going to be as easy if any of those external packages use the BSD license and we don't
realize it.

I propose that each repository should include a full list of licenses and copyright notices based on
what is used to generate a binary. It's not that hard to do, makes things a lot better for
downstream users, and gives people the credit they want.
