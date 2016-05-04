---
layout: inner
title: doing a request with basic auth
tags: ["ruby"]
---

{% highlight ruby linenos %}
require "net/http"

url = URI.parse('http://localhost:3000/users/1.json')
req = Net::HTTP::Get.new(url.path)
req.basic_auth 'user', 'pass'

resp = Net::HTTP.new(url.host, url.port).start {|http| http.request(req) }
puts resp

{% endhighlight %}
