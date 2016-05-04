---
layout: inner
title: stopping an infinite stream in elixir
tags: ["elixir","erlang"]
---
Suppose you have an infinite stream. You can stop it as shown [here](http://stackoverflow.com/a/25173872/31610)

{% highlight elixir %}
try do
  for tweet <- ExTwitter.stream_filter(track: terms) do
    process_tweet tweet
    if logic_to_determine_halt?, do: throw :halt
  end
catch 
  :halt -> #finished
end
{% endhighlight %}

