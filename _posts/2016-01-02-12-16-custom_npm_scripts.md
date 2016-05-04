---
layout: inner
title: custom npm scripts
tags: ["node","javascript","npm"]
---

You can define your own scripts in package.json, like:

{% highlight javascript %}
  "scripts": {
    "fix-tags": "node app.js tags/index.html"
  }
{% endhighlight %}

and you can run them like this:

{% highlight bash %}
npm run-script fix-tags
{% endhighlight %}
