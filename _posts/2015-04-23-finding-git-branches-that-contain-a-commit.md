---
layout: inner
title: "finding git branches that contain a commit"
description: ""
category: ""
tags: ["git"]
---
Information found [here](http://stackoverflow.com/questions/1419623/how-to-list-branches-that-contain-a-given-commit):

{% highlight bash %}
git branch --contains SOMESHA
{% endhighlight %}

and to find only local ones:

{% highlight bash %}
git branch -r --contains SOMESHA
{% endhighlight %}
