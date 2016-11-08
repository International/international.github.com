---
layout: inner
title: elixir [error] backend port not found :inotifywait in ubuntu
tags: ["ubuntu","elixir"]
---
If you get this error: <b>[error] backend port not found: :inotifywait</b>, then you
need to install inotify-tools, like this:

{% highlight bash %}
sudo apt-get install inotify-tools
{% endhighlight %}
