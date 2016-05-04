---
layout: inner
title: fixing X11/extensions/XTest.h No such file or directory in ubuntu
tags: ["ubuntu","x11","node"]
---
If you get the error: `X11/extensions/XTest.h: No such file or directory`, install:

{% highlight bash %}
sudo apt-get install libxtst-dev
{% endhighlight %}
