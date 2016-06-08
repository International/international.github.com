---
layout: inner
title: amazon sns http endpoint post
tags: ["amazon","aws","sns","ruby"]
---
Turns out the requests coming from Amazon SNS have `Content-Type: text/plain`, even
though they are JSON. One way to handle them in a rails app is this:

{% highlight ruby %}
amazon_params = JSON.parse(request.raw_post)
{% endhighlight %}

Notice that we're using `raw_post`, because params won't have the body parsed.
