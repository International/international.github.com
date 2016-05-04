---
layout: inner
title: "debug cors with curl"
description: ""
category: ""
tags: ["curl"]
---
Example take from [here](http://stackoverflow.com/questions/12173990/how-can-you-debug-a-cors-request-with-curl)
{% highlight bash %}
curl -H "Origin: http://example.com" --verbose \
  https://www.googleapis.com/discovery/v1/apis?fields=
{% endhighlight %}
