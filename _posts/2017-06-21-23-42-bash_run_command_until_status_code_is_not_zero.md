---
layout: inner
title: bash run command until status code is not zero
tags: ["bash","shell"]
---
Problem: want to run a command until the exit code is not zero

Solution: found [here](https://stackoverflow.com/a/12967264)

{% highlight bash %}
while ./runtest; do :; done
{% endhighlight %}