---
layout: inner
title: find program using port
tags: ["osx","port","lsof"]
---
Problem: find which process is using a specific port.

Solution: found [here](https://www.mkyong.com/mac/mac-osx-what-program-is-using-port-8080/)

{% highlight bash %}
lsof -i :8080 | grep LISTEN
{% endhighlight %}