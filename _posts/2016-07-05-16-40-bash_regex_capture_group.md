---
layout: inner
title: bash regex capture group
tags: ["regex","bash"]
---
capturing a group in bash ( regex <b>MUST NOT BE QUOTED!</b> )
{% highlight bash %}
$ [[ '${these}' =~ \${(.+)} ]] && echo ${BASH_REMATCH[1]}
these
{% endhighlight %}
