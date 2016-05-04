---
layout: inner
title: performing a count operation in ecto
tags: ["ecto","elixir"]
---
{% highlight elixir %}
iex(3)> import Ecto.Query
nil
iex(4)> Repo.all(from f in User, select: count(f.id))
[debug] SELECT count(u0."id") FROM "users" AS u0 [] OK query=67.8ms queue=10.5ms
[1]
{% endhighlight %}

or alternatively:

{% highlight elixir %}
iex(5)> Repo.one(from f in User, select: count(f.id))
1
[debug] SELECT count(u0."id") FROM "users" AS u0 [] OK query=0.6ms
{% endhighlight %}
