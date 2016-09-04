---
layout: inner
title: elixir supervised tasks
tags: ["elixir","otp","supervisor", "task"]
---
Useful information found [here](http://stackoverflow.com/a/36606238/31610)
{% highlight elixir %}
{:ok, tas} = Task.Supervisor.start_link(restart: :transient, max_restarts: 4)
Task.Supervisor.start_child(tas, fn ->
  IO.puts "starting child"
  1 = 2 
end)
{% endhighlight %}
