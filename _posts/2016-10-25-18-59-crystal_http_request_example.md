---
layout: inner
title: crystal http request example
tags: ["crystal", "http"]
---
{% highlight ruby %}
require "http/client"

response   = HTTP::Client.get("https://some.url")

if response.status_code == 200
  # do something with the response
  content    = response.body
end
{% endhighlight %}
