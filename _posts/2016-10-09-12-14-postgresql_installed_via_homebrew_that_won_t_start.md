---
layout: inner
title: postgresql installed via homebrew that won't start
tags: ["osx","postgres","homebrew"]
---
After a reboot, postgresql wouldn't start. Tried `brew services postgresql restart` and still nothing.
Checked the `/usr/local/var/postgres/server.log` and it was complaining about `/usr/local/var/postgres/postmaster.pid`.
My first thought was to delete it, but I ran a google search before. Luckily, I found [this link](http://superuser.com/a/553545/11904) where
it clearly says that deleting the pid file is not recommended. So, I ran a `cat` command and found the PID ( it's the first line in the file ),
and did a `kill -9 PID`. After that, running again `brew services postgresql restart` worked, and PG booted.
