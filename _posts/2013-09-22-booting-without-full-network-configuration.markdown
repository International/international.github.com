---
layout: inner
title: booting without full network configuration
tags: ["linux"]
---

So that I didn't have to wait almost 2 min for my Ubuntu to book, I changed my /etc/network/interfaces file to this:

{% highlight bash %}
# interfaces(5) file used by ifup(8) and ifdown(8)
auto lo
iface lo inet loopback
{% endhighlight %}

