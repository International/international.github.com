---
layout: inner
title: angular1 controllerAs with ngRoute
tags: ["javascript","angular","angular1"]
---
Problem: want to use <b>controller as</b> with <b>ngRoute</b>

Solution:

* found [here](http://stackoverflow.com/a/23861159/31610)
{% highlight javascript %}
$routeProvider
    .when('/somePath', {
        template: htmlTemplate,
        controller: 'myController as ctrl'
    });
{% endhighlight %}