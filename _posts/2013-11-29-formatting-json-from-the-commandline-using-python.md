---
layout: inner
title: "formatting JSON from the commandline using Python"
description: ""
category: ""
tags: ["python","JSON"]
---
Say you're doing a CURL request that returns some JSON. If you want to 
pretty-print it, here's an easy way to do so:

{% highlight bash %}
curl http://some_host/action.json | python -m json.tool
{% endhighlight %}
