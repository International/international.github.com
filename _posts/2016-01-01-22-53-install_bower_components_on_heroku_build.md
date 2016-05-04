---
layout: inner
title: install bower components on heroku build
tags: ["node","javascript","bower","heroku"]
---
Make sure you have bower added as a dependency in package.json, and then add the following
script as postinstall:

{% highlight javascript %}
"scripts": {
  "postinstall": "./node_modules/bower/bin/bower install"
}
{% endhighlight %}
