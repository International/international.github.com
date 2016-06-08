---
layout: inner
title: curl post json body from file
tags: ["curl"]
---
{% highlight bash %}
curl -H "Content-Type: application/json" -X POST -d @file.json url
{% endhighlight %}
