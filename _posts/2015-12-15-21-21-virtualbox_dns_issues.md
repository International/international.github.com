---
layout: inner
title: virtualbox dns issues
tags: ["virtualbox"]
---
Found that running this, solves my "virtualbox temporary failure in name resolution" errors.

{% highlight bash %}
VBoxManage modifyvm Linux --natdnshostresolver1 on
{% endhighlight %}
