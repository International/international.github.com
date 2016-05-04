---
layout: inner
title: find port where hdfs is listening
tags: ["hdfs","hadoop"]
---
{% highlight bash %}
hdfs getconf -confKey fs.default.name
{% endhighlight %}
