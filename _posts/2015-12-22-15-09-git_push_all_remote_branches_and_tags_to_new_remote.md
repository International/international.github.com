---
layout: inner
title: git push all remote branches and tags to new remote
tags: ['git']
---
{% highlight bash %}
git push new_remote_name refs/remotes/origin/*:refs/heads/*
git push new_remote_name refs/tags/*:refs/tags/*
{% endhighlight %}
