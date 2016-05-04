---
layout: inner
title: "directive with controller in angular"
description: ""
category: ""
tags: ["angularjs","javascript"]
---
Been trying to get a directive to use a controller method. I had this:
{% highlight coffeescript %}
app.directive "myDirective", ->
  restrict: "A"
  controller: ($scope) ->
    sayHello: ->
      alert "Hello"
  template: '<div><a ng-click="sayHello()">Say hello</a></div>'
{% endhighlight %}

And it wouldn't react. The problem was that the method had to be dined on the scope:

{% highlight coffeescript %}
app.directive "myDirective", ->
  restrict: "A"
  controller: ($scope) ->
    $scope.sayHello = ->
      alert "Hello"
  template: '<div><a ng-click="sayHello()">Say hello</a></div>'
{% endhighlight %}
