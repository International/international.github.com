---
layout: inner
title: facebook flow comment types
tags: ["javascript","flow"]
---
Problem: you want to add flow, but you don't want to complicate your build process.

Solution: use comment types.

{% highlight javascript %}
const myInt /*: number */ = 10;
const someFnc = (param /*: number */) => {
  // do something
}
{% endhighlight %}