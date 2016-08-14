---
layout: inner
title: accessing shared folders in VirtualBox
tags: ["VirtualBox"]
---
To be able to view the contents of a shared folder in VirtualBox, your user needs
to be part of the `vboxsf` group:
<pre>sudo usermod -a -G vboxsf cloudera</pre>

The groups the user is part of do not apply right away. Either login again, or
restart the VM for this to take effect.
