---
layout: inner
title: git multiple worktrees
tags: ["git"]
---
This checks out the branch called `devel` in a folder called `dev-code`. It
makes it easy to keep track of multiple branches of your project without having
to have multiple clones, or to keep switching from one branch to the other :

{% highlight bash %}
git worktree add dev-code devel
{% endhighlight %}
