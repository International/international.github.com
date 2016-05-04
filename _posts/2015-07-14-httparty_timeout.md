---
layout: inner
title: HTTParty timeout
tags: ["ruby", "httparty"]
---
{% highlight ruby %}
HTTParty.post("some_url", :body => {}, :timeout => 120)
{% endhighlight %}
