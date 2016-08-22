---
layout: inner
title: WARNING REMOTE HOST IDENTIFICATION HAS CHANGED
tags: ["ssh"]
---
If you get this message:
<pre>WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!</pre>

when trying to use ssh to login, you can run this command:
<pre>ssh-keygen -R <host_or_ip></pre>

If ssh-keygen fails, check your `.ssh/known_hosts` file for the entry and delete it.
By the way, do this only if you're sure of what you're doing. The reason I wrote this,
is that I've encountered this message a few times with some VirtualBox machines. 
You might want to double check this ( or probably not do it at all ), if this
happens for a machine connected to the internet ( think staging/prod servers )
