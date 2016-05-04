---
layout: inner
title: "comparing binary files"
description: ""
category: ""
tags: ["diff","binary"]
---
Here's how to compare two binary files:
{% highlight bash %}
diff <(xxd some.bin) <(xxd some_other.bin)
{% endhighlight %}
