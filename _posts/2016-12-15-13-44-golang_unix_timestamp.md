---
layout: inner
title: golang unix timestamp
tags: ["go","golang"]
---
Found answer [here](http://stackoverflow.com/a/9539644)
{% highlight go %}
int32(time.Now().Unix())
{% endhighlight %}
