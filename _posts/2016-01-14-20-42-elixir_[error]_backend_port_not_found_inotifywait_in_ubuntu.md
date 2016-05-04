---
layout: inner
title: elixir [error] backend port not found :inotifywait in ubuntu
tags: ["ubuntu","elixir"]
---
If you get this error: `[error] backend port not found: :inotifywait`, then you
need to install inotify-tools, like this:

{% highlight bash %}
sudo apt-get install inotify-tools
{% endhighlight %}
