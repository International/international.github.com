---
layout: inner
title: generate pem file from crt and key
tags: ["ssl","pem","crt"]
---
Problem: have a <b>.key</b> and a <b>.crt</b> and want to generate a <b>.pem</b> file from them.

Solution:

* a simple cat will do

{% highlight bash %}
cat server.crt server.key > server.includespk.pem
{% endhighlight %}