---
layout: inner
title: debugging elixir/phoenix tests
tags: ["elixir","phoenix","iex","pry", "debug"]
---
{% highlight elixir %}
iex -S mix test path/to/your/test.exs
{% endhighlight %}

Also make sure that you have the proper require in your test:

{% highlight elixir %}
defmodule YourApp.YourTest do
  use ExUnit.Case

  require IEx

  # in case 60s is not enough
  @tag timeout: 120000
  test "it should work" do
    # whatever
    IEx.pry
    # inspect
  end
end

{% endhighlight %}
