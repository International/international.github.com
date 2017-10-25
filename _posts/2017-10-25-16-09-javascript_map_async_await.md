---
layout: inner
title: javascript map async await
tags: ["javascript","map","async","await","node","selenium","webdriver"]
---
Problem: want to obtain the text from an array of elements

Solution:

* <b>Promise.all</b> and async await map:

{% highlight javascript %}
const childrenTable = await driver.findElement(By.css('table'))
const childrenEntries = await childrenTable.findElements(By.css('tr'))
const childrenText = await Promise.all(childrenEntries.map(async (cd) => cd.getText()))
{% endhighlight %}