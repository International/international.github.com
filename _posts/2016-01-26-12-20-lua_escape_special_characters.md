---
layout: inner
title: lua escape special characters
tags: ["lua"]
---
Found a very good way of escaping punctuation characters in lua [here](http://stackoverflow.com/a/13454747/31610)

{% highlight lua %}
str:gsub ("%p", "%%%0")
{% endhighlight %}
