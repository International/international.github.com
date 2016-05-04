---
layout: inner
title: git worktree is already checked out at
tags: ['git']
---
I received errors like `fatal: 'devel' is already checked out at some_location`.
This went away after removing the folder:

{% highlight bash %}
rm -rvf .git/worktrees/some_location
{% endhighlight %}
