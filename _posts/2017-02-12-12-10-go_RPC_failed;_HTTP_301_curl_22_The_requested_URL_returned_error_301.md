---
layout: inner
title: go RPC failed; HTTP 301 curl 22 The requested URL returned error 301
tags: ["go","golang","git"]
---
Problem: received this message when trying to install <b>gopkg.in/go-playground/validator.v9</b>

<pre>
RPC failed; HTTP 301 curl 22 The requested URL returned error: 301
</pre>

Solution: Found [here](https://github.com/niemeyer/gopkg/issues/50#issuecomment-273299592):

<pre>
git config --global http.https://gopkg.in.followRedirects true
</pre>
