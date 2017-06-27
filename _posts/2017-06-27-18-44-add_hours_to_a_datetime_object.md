---
layout: inner
title: add hours to a datetime object
tags: ["python","time","datetime"]
---
Problem: want to add hours to a datetime object

Solution: use <b>timedelta</b>

{% highlight python %}
from datetime import datetime, timedelta
print datetime.now() + timedelta(hours=5)
{% endhighlight %}