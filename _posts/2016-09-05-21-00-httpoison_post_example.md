---
layout: inner
title: httpoison post example
tags: ["elixir","httpoison"]
---
{% highlight elixir %}
post_body           = %{"events" => shows} |> Poison.encode!
{:ok, response}     = HTTPoison.post(huginn_url, post_body, %{"Content-Type" => "application/json"})
%{status_code: 201} = response
{% endhighlight %}
