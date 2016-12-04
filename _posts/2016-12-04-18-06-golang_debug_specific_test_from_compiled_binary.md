---
layout: inner
title: golang debug specific test from compiled binary
tags: ["go","golang","debug"]
---
Problem: want to debug a failing test, from a compiled binary containing a lot of
other tests.

Solution:

* compile tests with <b>go test -c</b>
* run dlv: <b>dlv exec ./my.test -- -test.run SomePattern</b>
* debug!
