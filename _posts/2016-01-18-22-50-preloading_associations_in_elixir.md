---
layout: inner
title: preloading associations in elixir
tags: ["phoenix","elixir","ecto"]
---
{% highlight elixir %}
Repo.get(User, 1) |> Repo.preload(:association_name)
{% endhighlight %}

or, if an association has not been loaded, and you receive: `#Ecto.Association.NotLoaded<association :exercise_entries is not loaded>`
then, the following works too:

{% highlight elixir %}
u = Repo.get(User, 1)
p = Repo.preload(u, :association_name)
{% endhighlight %}
