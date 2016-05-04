---
layout: inner
title: git diff when you have a file with the same name as a branch
tags: ["git"]
---
Let's say you have a file called potato, and a branch called potato. If you want to see a diff against the branch called potato you'd run this:

{% highlight bash %}
git diff potato --
{% endhighlight %}

If you want to see a diff of the file potato:

{% highlight bash %}
git diff -- potato
{% endhighlight %}
