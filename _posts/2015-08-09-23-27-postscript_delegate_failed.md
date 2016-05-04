---
layout: inner
title: postscript delegate failed
tags: ["ruby","rmagick"]
---
Solution for error messages such as:
{% highlight bash %}
Postscript delegate failed `/home/ubuntu/workspace/project/spec/testfiles/example.pdf': No such file or directory @ error/pdf.c/ReadPDFImage/677
{% endhighlight %}

Is fixed by:
{% highlight bash %}
sudo apt-get install ghostscript
{% endhighlight %}
