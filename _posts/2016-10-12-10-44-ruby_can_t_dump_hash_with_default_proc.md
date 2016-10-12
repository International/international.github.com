---
layout: inner
title: ruby can't dump hash with default proc
tags: ["ruby"]
---
Received the error <b>ruby can't dump hash with default proc</b>. It was raised by a 3rd party lib, that was using `Marshal.dump`.
Turns out, I had a hash defined like this:

{% highlight ruby %}
hash = Hash.new{|h,k| h[k] = []}
{% endhighlight %}

The fix was to remove the initialization block, and instead use it like this in my code:
{% highlight ruby %}
hash[some_value] ||= []
hash[some_value] << some_other_value
{% endhighlight %}
