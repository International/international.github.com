---
layout: inner
title: hipchat with software rendering
tags: ["hipchat"]
---
If you get this error message: "Hipchat could not initiate hardware for rendering video", you can start hipchat with software rendering like this:

{% highlight bash %}
LIBGL_ALWAYS_SOFTWARE=1 hipchat
{% endhighlight %}
