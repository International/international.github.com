---
layout: inner
title: "accessing an angular service from the browser console"
description: ""
category: ""
tags: ["angularjs","javascript"]
---
This works in 1.2.10:

{% highlight javascript %}
angular.injector(["yourAppName"]).get("serviceName")
{% endhighlight %}
