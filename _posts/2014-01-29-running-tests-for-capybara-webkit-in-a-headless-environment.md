---
layout: inner
title: "running tests for capybara webkit in a headless environment"
description: ""
category: ""
tags: ["ruby","capybara","rails","linux"]
---
{% highlight bash %}
sudo apt-get install xvfb
DISPLAY=localhost:1.0 xvfb-run rake spec
{% endhighlight %}
