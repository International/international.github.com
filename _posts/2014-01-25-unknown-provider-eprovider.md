---
layout: inner
title: "unknown provider eProvider"
description: ""
category: ""
tags: ["angularjs","rails","minification","javascript"]
---
After minification, I received the error "unknown provider eProvider". 
The solution was to switch from something like this:

{% highlight coffeescript %}
controller: ($scope) ->
  $scope.closeModal = ->
    ModalService.close()
{% endhighlight %}

to something like this:

{% highlight coffeescript %}
controller: ["$scope"
  ($scope) ->
    $scope.closeModal = ->
      ModalService.close()
]
{% endhighlight %}

I found the solution [here](http://stackoverflow.com/questions/13459452/rails-3-angularjs-minification-does-not-work-in-production-unknown-provider)
