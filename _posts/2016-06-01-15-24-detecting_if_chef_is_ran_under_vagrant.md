---
layout: inner
title: detecting if chef is ran under vagrant
tags: ["chef","vagrant"]
---
{% highlight ruby %}
if node.name =~ /vagrant/
  # do stuff
end
{% endhighlight %}
