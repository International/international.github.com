---
layout: inner
title: setting headers on the capybara driver
tags: ["capybara","ruby","rails"]
---
{% highlight ruby %}
Capybara.current_session.driver.header("X-Ninja", "true")
{% endhighlight %}
