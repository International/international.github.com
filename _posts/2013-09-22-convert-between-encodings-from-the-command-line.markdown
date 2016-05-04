---
layout: inner
title: convert between encodings from the command line
tags: ["linux", "encoding"]
---

Here's how to change a file's encoding from windows-1250 to utf-8:

{% highlight bash linenos %}
iconv -f windows-1250 -t utf-8 some_file.txt
{% endhighlight %}

