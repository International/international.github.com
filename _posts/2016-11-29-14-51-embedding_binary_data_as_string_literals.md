---
layout: inner
title: embedding binary data as string literals
tags: ["golang","go"]
---
The problem: had a gzipped file that I wanted to include as a string literal in code.
The solution:

* used the following tool [go-bindata](https://github.com/jteeuwen/go-bindata):
<pre>
go-bindata -nocompress *.gz
</pre>

* found the <b>[]byte("")</b> representation of the file
* included it in the code
* done!
