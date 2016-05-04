---
layout: inner
title: "check if database exists in postgresql"
description: ""
category: ""
tags: ["postgresql"]
---
{% highlight sql %}
select * from pg_database where datname='some-name'
{% endhighlight %}
