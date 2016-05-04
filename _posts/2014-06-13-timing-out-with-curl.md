---
layout: inner
title: "timing out with curl"
description: ""
category: ""
tags: ["curl"]
---
To "timeout" a request after a number of seconds, use the -m switch:

{% highlight bash %}
curl -m 5 some_url
{% endhighlight %}

And in case the web server doesn't respond after 5 seconds, curl will exit with an error code.
