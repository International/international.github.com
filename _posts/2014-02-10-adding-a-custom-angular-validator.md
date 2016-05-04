---
layout: inner
title: "adding a custom angular validator"
description: ""
category: ""
tags: ["angularjs"]
---
* Fun validator, check if an exercise is one of the big 3:
{% highlight coffeescript %}
app.directive "compoundExercise", ->
  restrict: "A"
  require: "ngModel"
  link: (scope, element, attrs, ctrl) ->
    exercises = ["Squat", "Deadlift", "Bench Press"]
    ctrl.$formatters.push (value) ->
      haveX = value in exercises
      ctrl.$setValidity("compoundExercise", haveX)
      return value

    ctrl.$parsers.push (value) ->
      haveX = value in exercises
      ctrl.$setValidity("compoundExercise", haveX)
      return if haveX then value else null

{% endhighlight %}

* If the field is not valid, then myForm.myFieldName.$error.compoundExercise will be true

* Checking if the field is valid:
{% highlight coffeescript %}
myForm.myFieldName.$valid
{% endhighlight %}
