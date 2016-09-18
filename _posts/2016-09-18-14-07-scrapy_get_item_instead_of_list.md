---
layout: inner
title: scrapy get item instead of list
tags: ["python","scrapy"]
---
ItemLoader `.add_css` returns a list, instead of a single value. To get the first value, use this in your item:
{% highlight python %}
import scrapy
from scrapy.contrib.loader.processor import TakeFirst

class YourItem(scrapy.Item):
    description = scrapy.Field(output_processor=TakeFirst())
{% endhighlight %}
