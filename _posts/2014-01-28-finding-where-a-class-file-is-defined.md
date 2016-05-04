---
layout: inner
title: "finding where a class file is defined"
description: ""
category: ""
tags: ["ruby"]
---
Get a handle on one of the methods:
{% highlight ruby %}
m = YourClass.method(:some_method)
puts m.source_location
{% endhighlight %}
