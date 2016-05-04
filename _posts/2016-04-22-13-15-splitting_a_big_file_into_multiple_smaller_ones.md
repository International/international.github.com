---
layout: inner
title: splitting a big file into multiple smaller ones
tags: ["ubuntu","file","big"]
---
Used the following to split the cloudera zip into smaller ones:

{% highlight bash %}
split --bytes 300M --numeric-suffixes --suffix-length=3 cloudera-quickstart-vm-5.5.0-0-virtualbox.zip cloud.
{% endhighlight %}
