---
layout: inner
title: selenium webdriver chrome headless
tags: []
---
Problem: wanted to use chrome headless in selenium tests.

Solution: found [here](https://stackoverflow.com/a/44328140/31610). Since I'm using chrome 61, I don't need to set the binary to chrome.

{% highlight javascript %}
    const options = new chromeDriver.Options()
    options.addArguments(
        'headless',
        // Use --disable-gpu to avoid an error from a missing Mesa library, as per
        // https://chromium.googlesource.com/chromium/src/+/lkgr/headless/README.md
        'disable-gpu',
    )
    const browser = new webdriver.Builder()
        .forBrowser(targetBrowser)
        .setChromeOptions(options)
        .build();
    // ... do your thing here    
{% endhighlight %}