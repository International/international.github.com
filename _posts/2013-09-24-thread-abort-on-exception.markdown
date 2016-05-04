---
layout: inner
title: thread abort on exception
tags: ["ruby","thread"]
---
It might be useful to enable this while working with threads ( at least in early development phases ):

{% highlight ruby %}
Thread.abort_on_exception = true
{% endhighlight %}
