---
layout: inner
title: git export and apply patch
tags: ["git"]
---
Generating the patch from a commit:

{% highlight bash %}
git format-patch -1 SHA
{% endhighlight %}

Applying a patch:

{% highlight bash %}
git am < your.patch
{% endhighlight %}
