---
layout: inner
title: gunzip without overwriting interactively
tags: ["gunzip","shell"]
---
This will keep the archive after decompressing, and in case a file already exists, instead of seeing the prompt with `do you with to overwrite`,
`n` will instead be sent. Found solution [here](http://stackoverflow.com/a/24012314/31610)
<pre>yes n | gzip -v -dk *.gz</pre>
