---
layout: inner
title: bash read line handle case of no newline at end of file
tags: ["bash","shell"]
---
Found the answer [here](http://stackoverflow.com/a/5010679/31610):
{% highlight bash %}
while IFS= read -r LINE || [[ -n "$LINE" ]]; do
  echo "$LINE"
done
{% endhighlight %}
