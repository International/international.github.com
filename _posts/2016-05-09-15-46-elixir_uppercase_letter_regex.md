---
layout: inner
title: elixir uppercase letter regex
tags: ["elixir","regex","unicode"]
---
How to match uppercase unicode characters in Elixir
<!-- exce -->

Part of the solution of the 2nd exercise on exercism.io

{% highlight elixir %}
String.match?("УХОДИ", ~r/\p{Lu}{2,}/u)
{% endhighlight %}

