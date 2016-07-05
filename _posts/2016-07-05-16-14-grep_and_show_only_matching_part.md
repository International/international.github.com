---
layout: inner
title: grep and show only matching part
tags: ["grep","bash"]
---
Here's how to grep for something and show only the matching part:

{% highlight bash %}
# echo '${these}${are}${vars}' | grep -o -E '\${.*?}'
${these}
${are}
${vars}
{% endhighlight %}
