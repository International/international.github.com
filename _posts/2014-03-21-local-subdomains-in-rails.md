---
layout: inner
title: "local subdomains in rails"
description: ""
category: ""
tags: ["ruby","rails","subdomain"]
---
If you want to configure subdomains, using /etc/hosts, this will not work:

{% highlight bash %}
127.0.0.1 api.localhost
{% endhighlight %}

Instead, you need to use

{% highlight bash %}
127.0.0.1 api.localhost.local
{% endhighlight %}

As mentioned [here](http://therailworld.com/posts/30-Issue-with-localhost-subdomains-in-Rails-3), Rails 3 treats api.localhost as the
domain.
