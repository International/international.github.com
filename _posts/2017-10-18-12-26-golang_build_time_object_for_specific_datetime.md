---
layout: inner
title: golang build time object for specific datetime
tags: ["go","golang","time","date","datetime"]
---
Problem: want to have a datetime for a specific day, at a specific hour

Solution: one possible solution [here](https://stackoverflow.com/q/25254443/31610):

{% highlight go %}
func Bod(t time.Time) time.Time {
    year, month, day := t.Date()
    return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
{% endhighlight %}