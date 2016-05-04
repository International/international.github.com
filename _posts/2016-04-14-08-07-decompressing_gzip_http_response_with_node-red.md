---
layout: inner
title: decompressing gzip http response with node-red
tags: ["node.js","node-red"]
---
Add zlib in settings.js, if you don't already have it:

{% highlight javascript %}
functionGlobalContext: {
   zlib: require('zlib'),
   ...
}
     
{% endhighlight %}

Then create a new function node with the following code:

{% highlight javascript %}
var buffer = new Buffer(msg.payload);
msg.payload = global.get('zlib').gunzipSync(buffer).toString();
return msg;
{% endhighlight %}

This won't yield the best performance, but it works in my case. Most likely,
if could be done better by using the callback version of gunzip, and then using
send to pass the msg to the next node.
