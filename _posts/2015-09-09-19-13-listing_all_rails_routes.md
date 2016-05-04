---
layout: inner
title: listing all rails routes
tags: ['ruby','rails','rspec']
---
Found myself needing a list of routes in an rspec controller test. This did the trick:

{% highlight ruby %}
Rails.application.routes.routes.map{|e| e.path.spec.to_s }
{% endhighlight %}
