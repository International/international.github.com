---
layout: inner
title: ruby configuration info
tags: ["ruby"]
---
A lot of useful information can be found in the `RbConfig::Config` hash. For example,
the host os:

{% highlight ruby %}
[2] pry(main)> RbConfig::CONFIG['host_os']
=> "darwin15.4.0"
{% endhighlight %}
