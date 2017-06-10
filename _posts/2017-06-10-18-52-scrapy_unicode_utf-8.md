---
layout: inner
title: scrapy unicode utf-8
tags: []
---
Problem: want to write unicode strings encoded as utf-8 when exporting scrapy items.

Solution:

* found [here](https://stackoverflow.com/a/41346276/31610)
* add the following setting in your settings.py:

{% highlight python %}
FEED_EXPORT_ENCODING = 'utf-8'
{% endhighlight %}