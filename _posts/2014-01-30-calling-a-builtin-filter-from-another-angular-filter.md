---
layout: inner
title: "calling a builtin filter from another angular filter"
description: ""
category: ""
tags: ["javascript","angularjs","coffeescript"]
---
I needed to call dateFilter from one of my filters. This is how I've done it:

{% highlight coffeescript %}
app.filter "myCustomFilter", ["dateFilter", (dateFilter) ->
  (input) ->
    # call dateFilter with whatever args
]
{% endhighlight %}
