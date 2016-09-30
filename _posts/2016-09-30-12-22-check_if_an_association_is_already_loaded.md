---
layout: inner
title: check if an association is already loaded
tags: ["ruby","rails"]
---
Following can be useful in tracking N+1 queries, imagine you have a User with many Photos:

{% highlight ruby %}
user.photos.loaded?
{% endhighlight %}

the code above would tell you if the association has been already loaded. It proved useful when trying to
track down a N+1 query when serializing to JSON.
