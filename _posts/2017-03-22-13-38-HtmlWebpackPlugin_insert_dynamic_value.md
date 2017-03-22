---
layout: inner
title: HtmlWebpackPlugin insert dynamic value
tags: ["javscript","webpack", "react"]
---
Problem: add a value in `index.html` depending on a condition

Solution:

* modify webpack file, to have this ( notice the <b>someValue</b> key ):

{% highlight javascript %}
new HtmlWebpackPlugin({
      someValue: someValue,
      filename: 'index.html',
    }),
{% endhighlight %}

* then modify index.html with your logic, and make use of the value, like in the following snippet:

{% highlight html %}
<%=  htmlWebpackPlugin.options.someValue %>
{% endhighlight %}
