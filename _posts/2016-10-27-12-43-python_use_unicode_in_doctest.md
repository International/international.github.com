---
layout: inner
title: python use unicode in doctest
tags: ["python"]
---
Found this useful information [here](http://stackoverflow.com/a/1734503/31610):
{% highlight python %}
# -*- coding: utf-8 -*-
def mylen(word):
  u"""        <----- SEE 'u' HERE
  >>> mylen(u"áéíóú")
  5
  """
  return len(word)

print mylen(u"áéíóú")
{% endhighlight %}
