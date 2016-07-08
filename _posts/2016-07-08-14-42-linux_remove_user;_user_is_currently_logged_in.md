---
layout: inner
title: linux remove user user is currently logged in
tags: ["bash","linux"]
---
<pre>
$ sudo userdel username
userdel: user username is currently logged in
$ sudo pkill -u username
$ sudo userdel -r username
</pre>
