---
layout: inner
title: adding Println statements in golang stdlib
tags: ["go","golang","debug"]
---
* install go bootstrap: https://storage.googleapis.com/golang/go1.4-bootstrap-20161024.tar.gz
* build it ( note the directory it says at the end of installation), in my case:

<pre>Installed Go for darwin/amd64 in /Users/geo/bootstrapped-go</pre>

* clone go repo: https://go.googlesource.com/go
* check out the version you're interested in, in my case go1.7.4
* go to src folder
* GOROOT_BOOTSTRAP=/Users/boostrapped-go ./make.bash
* add your Println statements where needed
* any time you modify something in stdlib, you need to recompile
* repeat until problem found
