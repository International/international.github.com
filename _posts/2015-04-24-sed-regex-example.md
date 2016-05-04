---
layout: inner
title: sed regex replace backreference example
tags: ["sed", "regex", "backreference"]
---

I had a file containing something like this:
{% highlight ruby %}
ENV['SOME_VAR'] = "some_value"
ENV['SOME_OTHER_VAR'] = "some_other_value"
{% endhighlight %}

and I wanted to turn it into a shell script that would export those vars. This comes pretty close to it, and shows an example of
using replace using sed backreferences:

{% highlight bash %}
cat unicorn.conf  | grep ENV | sed "s/ENV\['\(.*\)'.*\]/export \1/g" | sed 's/ = /=/g'
{% endhighlight %}
