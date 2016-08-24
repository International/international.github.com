---
layout: inner
title: easy way to copy ssh key to another machine
tags: ["ssh"]
---
Install `ssh-copy-id` by running `brew install ssh-copy-id`. Then:
<pre>
ssh-copy-id -p PORT user@some_ip
</pre>

You need to use `PORT` if ssh is listening on something other than 22.
