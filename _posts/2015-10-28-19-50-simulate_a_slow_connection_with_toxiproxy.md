---
layout: inner
title: simulate a slow connection with toxiproxy
tags: ["ruby","proxy","toxiproxy"]
---

{% highlight ruby %}
Toxiproxy.populate([
  {
    name: "toxiproxy_test_redis_tags",
    listen: "127.0.0.1:22222",
    upstream: "127.0.0.1:6379"
  }
])

r = Redis.new(port: 22222)

Toxiproxy["toxiproxy_test_redis_tags"].downstream(:latency, latency: 5000).apply do
  puts Time.now
  puts "Value is", r.get("ok")
  puts Time.now
end

Toxiproxy["toxiproxy_test_redis_tags"].upstream(:timeout, timeout: 2000).apply do
  # whatever
end
{% endhighlight %}

The example above would give you a "slow" redis. I received the following error on my system:

{% highlight bash %}
Unit toxiproxy.service failed to load
{% endhighlight %}

when trying to start toxiproxy via:

{% highlight bash %}
sudo service toxiproxy start
{% endhighlight %}

But, you can just start it manually. Do a <b>updatedb</b> and then a <b>locate toxiproxy</b> in
case you don't know where your binary is, and then you can populate it like in the example above.
