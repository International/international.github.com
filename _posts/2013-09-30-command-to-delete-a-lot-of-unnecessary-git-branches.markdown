---
layout: inner
title: command to delete a lot of unnecessary git branches
tags: ["git"]
---
Assuming you have a lot of branches, and you only want to keep devel and master:

{% highlight bash %}
gb -a | grep -i origin | sed -E "s/^\s+remotes\/origin\///" | grep -vi master | grep -vi devel | while read br; do echo "deleting $br"; git push origin :$br; done
{% endhighlight %}
