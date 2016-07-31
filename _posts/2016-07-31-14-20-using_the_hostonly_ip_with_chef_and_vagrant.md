---
layout: inner
title: using the hostonly ip with chef and vagrant
tags: ["chef","ruby","vagrant"]
---
Example from a config file:

{% highlight ruby %}
listen_address: <%= node[:network][:interfaces][:eth1][:addresses].find{|k,v| v[:family] == "inet"}.first %>
{% endhighlight %}
