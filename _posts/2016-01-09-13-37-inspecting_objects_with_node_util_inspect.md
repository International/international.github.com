---
layout: inner
title: inspecting objects with node util inspect
tags: ["node","javascript", "debug"]
---
Default depth for <b>util.inspect</b> is 2, to make it infinite:

{% highlight javascript %}
util.inspect(some_object,{depth: null});
{% endhighlight %}
