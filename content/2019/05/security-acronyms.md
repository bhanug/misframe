---
title: Security acronyms
date: "2019-05-24T04:17:53.890Z"
twitter_card:
  description: "List of runtime and static analysis security acronyms and what they mean"
---

I joined [ShiftLeft](https://www.shiftleft.io/) a year ago, and since it's my first time in the security space I had to familiarize myself with the following terminology. I use this as a reference. Hopefully it's useful for you too!

<!--more-->

## Runtime

Runtime security involves inspection and protection against vulnerabilities while an application is running in production. Runtime security usually involves configuring a set of rules that apply to your application or architecture and enforcing them against incoming traffic.

### WAF: Web Application Firewall

A WAF is a type of firewall that inspects HTTP traffic at the edge and blocks attacks.

**Example:** [ModSecurity](https://en.wikipedia.org/wiki/ModSecurity) is an open-source WAF module for web servers like Apache.

### RASP: Runtime Application Self-Protection

RASP products and tools modify and instrument a running application to protect it against attacks. For example, a RASP could hook into SQL database library functions to block potential SQL injection attacks.

## Static analysis

Static analysis involves inspection of applications by looking at source code or byte code, or running automated scans against a running application to uncover any potential attack surfaces. Static analysis happens earlier on in the software development lifecycle, before you deploy your applications to production.

### SCA: Software Composition Analysis

SCA refers to identifying vulnerabilities by looking at an application's open-source dependencies. This can be done by inspecting manifest files like `package.json` and referring to a vulnerability database.

**Example:** GitHub can inspect your project's open-source dependencies to find known vulnerabilities. See https://help.github.com/en/articles/about-security-alerts-for-vulnerable-dependencies.

### SAST: Static Application Security Testing

SAST involves analyzing a program's source code (or byte code) to find vulnerabilities. For example, a SAST product or tool could check for SQL injection vulnerabilities by looking for unsanitized strings from an external source that end up in a SQL query.

### DAST: Dynamic Application Security Testing

DAST is a black-box security testing approach. This involves using a scanner like the [Burp Suite](https://portswigger.net/burp) which crawls an application and attempts various attacks.

### IAST: Interactive Application Security Testing

IAST is like DAST, but also instruments an application to find vulnerabilities for a more focused scanning approach.