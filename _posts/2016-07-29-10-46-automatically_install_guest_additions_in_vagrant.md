---
layout: inner
title: automatically install guest additions in vagrant
tags: ["vagrant"]
---
[This plugin](https://github.com/dotless-de/vagrant-vbguest) makes it all happen.
In my case, it was enough just to do a:
<pre>
vagrant plugin install vagrant-vbguest
</pre>

and guest additions were automatically installed.
