---
layout: inner
title: ecto simple update
tags: ["elixir","ecto"]
---
{% highlight elixir %}
u = Repo.get(User, 1)
# and to change it's updated_at
Repo.update!(u |> Ecto.Changeset.change(%{updated_at: Ecto.DateTime.utc}))
{% endhighlight %}
