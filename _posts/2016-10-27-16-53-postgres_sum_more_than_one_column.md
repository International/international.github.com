---
layout: inner
title: postgres sum more than one column
tags: ["postgres","sql"]
---
Found [here](http://stackoverflow.com/a/15719644/31610):
{% highlight sql %}
SELECT COALESCE(col1,0) + COALESCE(col2,0)
FROM yourtable
{% endhighlight %}
