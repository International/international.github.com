---
layout: inner
title: elixir simple regex capture
tags: ["elixir","regex"]
---
{% highlight elixir %}
iex(5)> q = "one 2 three 4 five"
"one 2 three 4 five"
iex(6)> Regex.match?(~r/\d+/, q)
true
iex(10)> Regex.scan(~r/(\d+)/, q)               
[["2", "2"], ["4", "4"]]
iex(11)> Regex.scan(~r/(\d+)/, q, capture: :all_but_first)
[["2"], ["4"]]
iex(12)> Regex.scan(~r/(\d+)/, q, capture: :all_but_first) |> List.flatten
["2", "4"]
{% endhighlight %}
