---
layout: inner
title: handling ctrl-c in a bash script
tags: ["bash","shell","signal"]
---
{% highlight bash %}
#!/bin/bash

trap ctrl_c INT

function ctrl_c() {
  # do whatever cleanup you want here
}

{% endhighlight %}
