---
layout: inner
title: "handling uninitialized constant ActiveRecord::RecordNotFound"
description: ""
category: ""
tags: ["ruby", "activerecord"]
---
I kept getting "uninitialized constant ActiveRecord::RecordNotFound" in a Rack app.
The fix was to add this at the beginning of the file:
{% highlight ruby %}
require "active_record/errors"
{% endhighlight %}
