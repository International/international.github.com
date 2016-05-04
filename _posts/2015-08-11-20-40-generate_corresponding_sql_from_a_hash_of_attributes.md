---
layout: inner
title: generate corresponding sql from a hash of attributes
tags: ["ruby","rails","activerecord"]
---
{% highlight ruby %}
User.send(:sanitize_sql_for_conditions, User.first.attributes)
{% endhighlight %}
