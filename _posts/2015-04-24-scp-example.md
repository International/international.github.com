---
layout: inner
title: scp example
tags: ["scp"]
---

Here's how to scp a file, using a pem file:

{% highlight bash %}
scp -i ~/my_key.pem ubuntu@some.host:/tmp/some_file.dump .
{% endhighlight %}

