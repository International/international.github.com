---
layout: inner
title: schedule a sidekiq worker to run in a few minutes
tags: ["ruby","sidekiq"]
---
{% highlight ruby %}
MyWorker.perform_in(5.minutes, worker_args)
{% endhighlight %}
