---
layout: inner
title: run rspec test matching regex
tags: ["ruby","rspec","rails"]
---
I wanted to run just a test matching a specific pattern, from an integration suite. 
The SPEC_OPTS variable allowed me to do just that:

{% highlight bash %}
env SPEC_OPTS='-e translations' SELENIUM=1 bundle exec rspec spec/features/*
{% endhighlight %}
