---
layout: inner
title: "creating and debugging mongo indexes"
description: ""
category: ""
tags: ["mongo","index"]
---
To check index usage, start your mongod with -vv flags, example:
{% highlight bash %}
/usr/bin/mongod -vv --config /etc/mongod.conf
{% endhighlight %}

And then cat your logfile, probably in this location:

{% highlight bash %}
cat /var/log/mongodb/mongod.log
{% endhighlight %}

Look for the nscanned and nscannedObjects fields.
