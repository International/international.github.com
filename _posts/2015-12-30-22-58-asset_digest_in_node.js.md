---
layout: inner
title: asset digest in node.js
tags: ["node","javascript","asset"]
---
[node-static-asset](https://github.com/bminer/node-static-asset) can help with this.

Like the docs say, add this to your app:

{% highlight javascript %}

var express = require('express');
var app = express();
var staticAsset = require('static-asset');
app.use(staticAsset(__dirname + "/public/") );
app.use(express.static(__dirname + "/public/") );

{% endhighlight %}

And then in the view, use it like this:

{% highlight jade %}
script(type="text/javascript", src=assetFingerprint("/javascripts/jquery.min.js") )
{% endhighlight %}
