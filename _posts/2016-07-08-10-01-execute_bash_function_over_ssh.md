---
layout: inner
title: execute bash function over ssh
tags: ["bash","shell","function","ssh"]
---
{% highlight bash %}
#!/bin/bash

function show_uptime() {
  uptime
}

PEM_FILE=some/path
ssh -i ${PEM_FILE} $1 "$(typeset -f show_uptime); show_uptime"
{% endhighlight %}
