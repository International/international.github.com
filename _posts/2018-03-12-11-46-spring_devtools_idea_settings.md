---
layout: inner
title: spring devtools idea settings
tags: ["java","spring"]
---
Problem: code not getting reloaded when something changes.

Solution: found [here](https://www.mkyong.com/spring-boot/intellij-idea-spring-boot-template-reload-is-not-working/)

* add devtools to your dependencies
* go to Preferences -> Build, execution, deployment -> Compiler and check "Build project automatically"
* press cmd-shift-a , and select registry
* find and check <b>compiler.automake.allow.when.app.running</b>

{% highlight javascript %}
// some code here
{% endhighlight %}