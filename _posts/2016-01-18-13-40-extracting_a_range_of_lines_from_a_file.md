---
layout: inner
title: extracting a range of lines from a file
tags: ["unix", "sed"]
---
{% highlight bash %}
sed -n 10000,20000p initial_file.log > smaller.log
{% endhighlight %}

and to print all the lines after a specific line number:

{% highlight bash %}
sed -n '10000,$p'
{% endhighlight %}
