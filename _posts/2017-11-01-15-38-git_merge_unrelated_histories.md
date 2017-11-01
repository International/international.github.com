---
layout: inner
title: git merge unrelated histories
tags: ["git"]
---
Problem: want to merge 2 branches that don't share a common history.

Solution:

{% highlight shell %}
git merge --allow-unrelated-histories somebranch
{% endhighlight %}