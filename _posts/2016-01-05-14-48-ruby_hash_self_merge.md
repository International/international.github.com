---
layout: inner
title: ruby hash self merge
tags: ["ruby","rails"]
---
Useful trick, if you have a hash like this:

{% highlight ruby %}
a = {1 => [2,3,4], 2 => [3,4,5], 3 => [4,5,6,7]}                                                                                {% endhighlight %}

and you would like to get to a hash containing the same keys, but the count of the values:
{% highlight ruby %}
a = {1 => 3, 2 => 3, 3 => 4}
{% endhighlight %}

{% highlight ruby %}
a.merge(a) {|k,v| v.count }                                                                                                     {% endhighlight %}

[hash#merge](http://docs.ruby-lang.org/en/2.0.0/Hash.html#method-i-merge) takes the following params:

{% highlight ruby %}
merge(other_hash){|key, oldval, newval| block} → new_hash
{% endhighlight %}
