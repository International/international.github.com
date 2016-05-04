---
layout: inner
title: typescript instantiate generic type
tags: ["typescript","node","javascript"]
---
{% highlight javascript %}

function buildObject<T>(type: { new(): T } ): T {
  return new type();
}

{% endhighlight %}
