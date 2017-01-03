---
layout: inner
title: selenium webdriver click select option
tags: ["selenium", "javascript"]
---
Problem: select a specific option in a select dropdown

Solution: found [here](http://stackoverflow.com/a/28157070/31610):

{% highlight javascript %}
driver.findElement(wd.By.css('#month>option[title=\'November\']')).click();
{% endhighlight %}
