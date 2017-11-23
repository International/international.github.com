---
layout: inner
title: easy way to copy ssh key to another machine
tags: ["ssh"]
---
Install <b>ssh-copy-id</b> by running <b>brew install ssh-copy-id</b>. Then:

<pre>
ssh-copy-id -p PORT user@some_ip
</pre>

You need to use <b>PORT</b> if ssh is listening on something other than 22.
