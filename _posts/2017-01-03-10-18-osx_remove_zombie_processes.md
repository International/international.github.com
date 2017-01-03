---
layout: inner
title: osx remove zombie processes
tags: []
---
Problem: zombie process appeared in <b>ps</b>

Solution: find it's parent process,and kill it

<pre>
ps -xo pid,ppid,stat,command | grep your_zombie_process_name
kill -9 ppid_of_zombie_process
</pre>

Found this useful information [here](http://superuser.com/a/506048/11904)
