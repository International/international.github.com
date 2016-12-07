---
layout: inner
title: golang goqu ignore struct field
tags: ["golang", "go", "database"]
---
Problem: unexported struct field was used in a goqu query. The field did not have
an equivalent db column.

Solution: mark it like in JSON:
{% highlight go %}
type someName struct {
  myVirtualField int `db:"-"`
}
{% endhighlight %}
