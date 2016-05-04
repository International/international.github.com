---
layout: inner
title: "git read tree to move files around"
description: ""
category: ""
tags: ["git"]
---
I wanted to copy a directory with all of it's files from a branch to the current one. 
This seems to have done the trick:

{% highlight bash %}
git read-tree --prefix=lib/zombie -u devel:lib/zombie
{% endhighlight %}

The prefix switch specifies where to place it, and the -u represents what I want to copy.
