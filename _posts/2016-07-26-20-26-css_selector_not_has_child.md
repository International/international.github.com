---
layout: inner
title: css selector not has child
tags: ["css","jquery","javascript"]
---
Example for selecting a's that don't have img as children:

{% highlight css %}
table:nth-of-type(14) td:nth-of-type(2) > a:not(:has(>img))
{% endhighlight %}
