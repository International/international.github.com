---
layout: inner
title: "bash pipe exit code"
description: ""
category: ""
tags: ["bash","linux","error"]
---
If you pipe a command to another, and want to keep the exit code, you can use one of these 2 ways:

{% highlight bash %}
# now you can et the error code in $?
set -o pipefail
{% endhighlight %}

Or

{% highlight bash %}
echo ${PIPESTATUS[0]}
{% endhighlight %}
