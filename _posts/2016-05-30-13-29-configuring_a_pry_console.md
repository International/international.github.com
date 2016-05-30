---
layout: inner
title: configuring a pry console
tags: ["ruby","debug","pry"]
---
{% highlight ruby %}
require "pry"

Pry.start
{% endhighlight %}

It's useful to know that if you do any kind of setup, or initialize any variables,
they will be available in the pry console as well.
