---
layout: inner
title: curl with proxy
tags: ["curl","proxy"]
---
Problem: want to make a curl request using a proxy.

Solution:

* Specify <b>-x</b> flag like in the example below ( assuming you have a proxy server running at localhost, on port 8080) :

{% highlight bash %}
curl -v -x http://localhost:8080 http://some.other.url
{% endhighlight %}