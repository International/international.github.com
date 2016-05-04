---
layout: inner
title: push local database to heroku
tags: ["heroku","postgres"]
---
In case you already have data, you have to do the reset. Make sure you know what
you're doing, because the reset is a destructive operation.

{% highlight bash %}
$ heroku pg:reset DATABASE_URL -r heroku
$ heroku pg:push your_db DATABASE_URL -r heroku      
{% endhighlight %}
