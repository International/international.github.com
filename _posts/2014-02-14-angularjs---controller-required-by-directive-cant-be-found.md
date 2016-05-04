---
layout: inner
title: "angularjs   controller, required by directive, can't be found"
description: ""
category: ""
tags: ["angularjs"]
---
This kind of errors:
{% highlight javascript %}
Controller 'ngModel', required by directive 'ngSubmit', can't be found!
{% endhighlight %}

Can be fixed by following the steps in [this](http://stackoverflow.com/questions/15099204/angularjs-error-no-controller-ngmodel?rq=1)
stackoverflow post. Basically, configure your directive require like this:

{% highlight coffeescript %}
require: "^?ngModel"
{% endhighlight %}

instead of:
{% highlight coffeescript %}
require: "ngModel"
{% endhighlight %}
