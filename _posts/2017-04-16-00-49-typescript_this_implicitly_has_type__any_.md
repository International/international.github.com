---
layout: inner
title: typescript this implicitly has type 'any'
tags: ["javascript","typescript"]
---
Problem: received this error in typescript: **'this' implicitly has type 'any' because it does not have a type annotation.**

Solution:

* Found [here](http://stackoverflow.com/a/41945742/31610)

{% highlight javascript %}
foo.on('error', function(this: Foo, err: any) {
{% endhighlight %}