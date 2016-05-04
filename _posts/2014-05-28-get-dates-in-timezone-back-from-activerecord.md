---
layout: inner
title: "get dates in timezone back from ActiveRecord"
description: ""
category: ""
tags: ["ruby","rails","activerecord","date","time"]
---
If you want that date/times come back in your timezone from the database, instead of
ActiveRecord::Base.default_timezone, make you sure you activate the following:

{% highlight ruby %}
ActiveRecord::Base.time_zone_aware_attributes = true
{% endhighlight %}
