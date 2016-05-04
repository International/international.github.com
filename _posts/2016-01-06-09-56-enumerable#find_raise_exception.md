---
layout: inner
title: enumerable#find raise exception
tags: ["ruby","rails"]
---
The following will raise an exception, and you won't have to deal with nil:

{% highlight ruby %}
[1,2,3].find(-> { raise "Expected to find 4 in the array"}) {|e| e == 4}
{% endhighlight %}
