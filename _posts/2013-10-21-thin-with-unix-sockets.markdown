---
layout: inner
title: thin with unix sockets
tags: ["ruby", "thin"]
---
To start thin with a unix socket, run this:
{% highlight bash %}
thin start --socket /tmp/some.sock
{% endhighlight %}
