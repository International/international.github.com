---
layout: inner
title: "setting a global gitignore file"
description: ""
category: ""
tags: ["git","ignore","gitignore"]
---
If you want to set a global gitignore file, that would apply to all your repos,
you can do so like this:

{% highlight bash %}
git config --global core.excludesfile '~/.gitignore'
{% endhighlight %}
