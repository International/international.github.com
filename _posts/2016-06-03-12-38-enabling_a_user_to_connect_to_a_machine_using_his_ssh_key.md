---
layout: inner
title: enabling a user to connect to a machine using his ssh key
tags: ["ssh"]
---
Assuming your user name is john, pasting the contents of the user's public key
in `/home/john/.ssh/authorized_keys` on the server, will allow him to login using only his
ssh key.
