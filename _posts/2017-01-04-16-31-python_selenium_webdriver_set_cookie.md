---
layout: inner
title: python selenium webdriver set cookie
tags: ["python","selenium","webdriver"]
---
Problem: use provided cookies with selenium python bindings

Solution: found [here](http://stackoverflow.com/a/37578697/31610):

{% highlight python %}
from selenium import webdriver

driver = webdriver.PhantomJS()
driver.get('http://stackoverflow.com/')

cookies = driver.get_cookies()

driver.delete_all_cookies()

for cookie in cookies :
    driver.add_cookie({k: cookie[k] for k in ('name', 'value', 'domain', 'path', 'expiry')})
{% endhighlight %}
