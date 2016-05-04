---
layout: inner
title: "install nokogiri with sys libraries"
description: ""
category: ""
tags: ["ruby","gem","nokogiri"]
---
Been trying to install nokogiri without much success on ruby 2. I already had it
working with ruby 1.9.3, and perhaps the fact that the previous version was using 
system's libxml2 influenced it somehow. Long story short, to make it install under
ruby 2, I did this:
{% highlight bash %}
NOKOGIRI_USE_SYSTEM_LIBRARIES=1 bundle install
{% endhighlight %}
