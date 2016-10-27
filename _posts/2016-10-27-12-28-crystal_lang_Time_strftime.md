---
layout: inner
title: crystal lang Time strftime
tags: ["crystal"]
---
The equivalent for the ruby:
{% highlight ruby %}
Time.now.strftime("%Y-%m-%d-%H-%M")
{% endhighlight }

is
{% highlight ruby %}
Time.now.to_s("%Y-%m-%d-%H-%M")
{% endhighlight %}
