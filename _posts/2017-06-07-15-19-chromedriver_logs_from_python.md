---
layout: inner
title: chromedriver logs from python
tags: ["python","selenium","chrome","chromedriver"]
---
Problem: wanted to see why chrome might fail to start.

Solution: specify a logging file when building your driver object:

{% highlight javascript %}
driver = webdriver.Chrome(service_args=["--verbose", "--log-path=/tmp/your.log"])
{% endhighlight %}