---
layout: inner
title: exposing a host service to a vagrant instance
tags: ["vagrant","ssh"]
---
I wanted to allow vagrant to connect to host's postgres:
{% highlight bash %}
vagrant ssh -- -R 5432:localhost:5432
{% endhighlight %}
