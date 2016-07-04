---
layout: inner
title: printing to output with groovy in a jenkins postbuild step
tags: ["groovy","jenkins"]
---
Found answer [here](http://stackoverflow.com/a/30704758/31610):

{% highlight groovy %}
manager.listener.logger.println("Some string")
{% endhighlight %}
