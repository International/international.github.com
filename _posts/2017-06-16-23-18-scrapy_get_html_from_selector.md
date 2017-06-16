---
layout: inner
title: scrapy get html from selector
tags: ["python","scrapy"]
---
Problem: need to access the html for a selector

Solution: found [here](https://stackoverflow.com/a/36249255/31610):

{% highlight python %}
from lxml import etree
etree.tostring(selector._root)
{% endhighlight %}