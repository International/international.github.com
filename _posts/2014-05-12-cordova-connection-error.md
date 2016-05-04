---
layout: inner
title: "cordova connection error"
description: ""
category: ""
tags: ["javascript","cordova","ionic"]
---
To handle this error:

Application Error
A network error occurred. (file:///android_asset/www/index.html)

The fix is to add a javascript redirect to the real index file:

{% highlight javascript %}
window.location = "./real_index.html";
{% endhighlight %}
