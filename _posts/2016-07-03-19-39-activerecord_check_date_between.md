---
layout: inner
title: activerecord check date between
tags: ["ruby","rails","activerecord"]
---
Was writing a raw query, when I remembered about range support:

{% highlight ruby %}
User.where(created_at: Time.zone.now-2.hours..Time.zone.now)
{% endhighlight %}
