---
layout: inner
title: accessing github over port 443
tags: ["git","github","firewall"]
---
Added the following to <b>.ssh/config</b>, cause 22 was blocked by the network firewall:
<pre>
Host github-work
  HostName ssh.github.com
  IdentityFile path_to_key
  User your_username
  Port 443
</pre>
