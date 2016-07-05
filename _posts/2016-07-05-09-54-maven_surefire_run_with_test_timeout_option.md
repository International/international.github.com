---
layout: inner
title: maven surefire run with test timeout option
tags: ["maven","java","timeout"]
---
Assuming you have the `maven-surefire-plugin`, [this](https://maven.apache.org/surefire/maven-surefire-plugin/test-mojo.html#forkedProcessTimeoutInSeconds) 
option is useful:

{% highlight bash %}
mvn -Dsurefire.timeout=60 test
{% endhighlight %}
