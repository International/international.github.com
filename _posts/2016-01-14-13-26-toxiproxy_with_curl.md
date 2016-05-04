---
layout: inner
title: toxiproxy with curl
tags: ["toxiproxy","curl"]
---
View all proxies:

{% highlight bash %}
curl http://localhost:8474/proxies | python -mjson.tool
{% endhighlight %}

Creating one:

{% highlight bash %}
$ curl -i -d '{"name": "toxic_mongo", "upstream": "localhost:27017", "listen": "localhost:22122"}' localhost:8474/proxies                            
HTTP/1.1 201 Created
Content-Type: application/json
Date: Thu, 14 Jan 2016 12:27:49 GMT
Content-Length: 617

{"name":"toxic_mongo","listen":"127.0.0.1:22122","upstream":"localhost:27017","enabled":true,"upstream_toxics":{"bandwidth":{"enabled":false,"rate":0},"latency":{"enabled":false,"latency":0,"jitter":0},"slicer":{"enabled":false,"average_size":0,"size_variation":0,"delay":0},"slow_close":{"enabled":false,"delay":0},"timeout":{"enabled":false,"timeout":0}},"downstream_toxics":{"bandwidth":{"enabled":false,"rate":0},"latency":{"enabled":false,"latency":0,"jitter":0},"slicer":{"enabled":false,"average_size":0,"size_variation":0,"delay":0},"slow_close":{"enabled":false,"delay":0},"timeout":{"enabled":false,"timeout":0}}}%  
{% endhighlight %}

Enabling a toxic:

{% highlight bash %}
curl -i -d '{"enabled":true, "latency":10000}' localhost:8474/proxies/toxic_mongo/downstream/toxics/latency
{% endhighlight %}


Deleting one:

{% highlight bash %}
curl -X DELETE localhost:8474/proxies/toxic_mongo
{% endhighlight %}
