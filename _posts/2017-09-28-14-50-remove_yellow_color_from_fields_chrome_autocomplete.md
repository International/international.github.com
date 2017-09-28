---
layout: inner
title: remove yellow color from fields chrome autocomplete
tags: ["css","chrome","html"]
---
Problem: want to get rid of the yellow color that appears when chrome autofills a field.

Solution: found [here](https://stackoverflow.com/a/40397889/31610)

{% highlight css %}
input:-webkit-autofill,
input:-webkit-autofill:hover,
input:-webkit-autofill:focus,
input:-webkit-autofill:active {
  -webkit-box-shadow: 0 0 0px 1000px white inset !important;
}
{% endhighlight %}