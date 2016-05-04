---
layout: inner
title: "accessing the target of a click in angularjs"
description: ""
category: ""
tags: ["angularjs", "javascript"]
---
Here's how to access the target of a click event:

{% highlight html %}
<a ng-click="doSomething($event)">Click me</a>
{% endhighlight %}

And in your method, you can access the target via $event.target
