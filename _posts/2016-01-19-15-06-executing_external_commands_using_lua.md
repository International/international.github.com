---
layout: inner
title: executing external commands using lua
tags: ["lua"]
---
{% highlight lua %}
io.popen"pwd":read'*l'
{% endhighlight %}
