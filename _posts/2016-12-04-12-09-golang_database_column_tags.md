---
layout: inner
title: golang database column tags
tags: ["go","golang","database"]
---

Comes in handy when you want a struct field to map to a db column, which may have
a different name. For example, I have <b>CreatedAt</b> field, but in the db, it's
<b>created_at</b>:

{% highlight go %}

type Agent struct {
	Name      string
	Config    string
	CreatedAt time.Time `db:"created_at"`
}

{% endhighlight %}
