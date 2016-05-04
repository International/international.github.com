---
layout: inner
title: golang time in different timezones
tags: ["go","golang","timezone","time"]
---
{% highlight go %}
import "time"

// to get the current time in utc, iso8601
location, err := time.LoadLocation("Etc/UTC")
now := time.Now()
fmt.Println(now.In(location).Format(time.RFC3339)
{% endhighlight %}
