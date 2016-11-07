---
layout: inner
title: compiling crystal from sources on OSX El Capitan
tags: ["crystal"]
---
This is easy if you have <b>homebrew</b>:
* brew install llvm
* clone the crystal repo:
<pre>
git clone git@github.com:crystal-lang/crystal.git
</pre>
* cd to that folder
* export PATH=$PATH:/usr/local/opt/llvm/bin
* run <b>make</b>
