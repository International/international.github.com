---
layout: inner
title: twisted debugging
tags: ["python","twisted","scrapy"]
---
Problem: had to work with some defer.Deferred and defer.DeferredList() with scrapy, and started to receive this kind of errors: <b>twisted.internet.defer.AlreadyCalledError</b>

Solution: this seems promising, add it after you import defer from <b>twisted.internet</b>

{% highlight python %}
defer.setDebugging(True)
{% endhighlight %}