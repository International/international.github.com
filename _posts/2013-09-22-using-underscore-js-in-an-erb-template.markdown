---
layout: inner
title: Using underscore js in an erb template
tags: coffeescript javascript underscore
---


This allows for mustache style templates, interpolation done with {{= }} and logic with {{ }}
 
{% highlight coffeescript %}
_.templateSettings = 
  interpolate: /\{\{\=(.+?)\}\}/g
  evaluate: /\{\{(.+?)\}\}/g

{% endhighlight %}
