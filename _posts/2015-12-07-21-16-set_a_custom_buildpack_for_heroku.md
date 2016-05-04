---
layout: inner
title: set a custom buildpack for heroku
tags: ["heroku","build","deploy"]
---
For example, setting typescript as your buildpack:

{% highlight bash %}
heroku buildpacks:set https://github.com/pk11/heroku-buildpack-typescript.git -r heroku
{% endhighlight %}
