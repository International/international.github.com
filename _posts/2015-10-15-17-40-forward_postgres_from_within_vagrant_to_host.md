---
layout: inner
title: forward postgres from within vagrant to host
tags: ["postgresql","vagrant","forward"]
---
To be able to connect from outside to postgresql running in vagrant, make sure that postgresql.conf has this:

<b>listen_addresses = '*'</b>

And then restart the server.
