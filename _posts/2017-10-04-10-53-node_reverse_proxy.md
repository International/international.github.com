---
layout: inner
title: node reverse proxy
tags: ["node", "express", "proxy", "reverse proxy"]
---
Problem: want to setup a reverse proxy for some urls

Solution: use [http-proxy-middleware](https://github.com/chimurai/http-proxy-middleware):

{% highlight javascript %}
const express = require('express');
const proxy = require('http-proxy-middleware');
const app = express();

const whereRequestsShouldBeProxiedTo = 'http://your.host'

const shouldProxy = req => {
    const prefixes = ['/api']

    return prefixes.find(e => req.url.startsWith(e))
}

const revProxy = proxy({
    target: whereRequestsShouldBeProxiedTo,
    changeOrigin: true,
})

app.use((req, res, next) => {
    if(shouldProxy(req)) {
        revProxy(req, res, next)
    } else {
        next()
    }
})

app.listen(3005)
{% endhighlight %}