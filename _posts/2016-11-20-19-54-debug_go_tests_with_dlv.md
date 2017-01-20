---
layout: inner
title: debug go tests with dlv
tags: ["go","debug","delve","golang"]
---
Useful info [here](https://github.com/derekparker/delve/issues/27):

* run:
<pre>go test -c</pre>

* then
<pre>dlv exec ./yourfile.test -- -test.run TestYourTestName</pre>

* debug
