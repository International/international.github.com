---
layout: inner
title: selenium webdriver javascript save screenshot
tags: ["selenium","javascript","node"]
---
Problem: obtaining a screenshot from selenium webdriver
Solution: The string that comes back as a result of <b>takeScreenshot</b> is a base64 png image, so
we decode it before writing it to disk:

{% highlight javascript %}
let browser = new webdriver.Builder().forBrowser('chrome').build()
...
browser.takeScreenshot().then(screenShotContents => {
  const fs = require('fs')
  fs.writeFileSync("screenshot.png", Buffer.from(screenShotContents, 'base64'))
})
{% endhighlight %}
