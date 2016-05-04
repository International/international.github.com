---
layout: inner
title: openresty deny request with access_by_lua
tags: ["nginx","openresty", "lua"]
---
Example config:

{% highlight lua %}
location ~ ^/some_location {
  access_log on;

  access_by_lua '
    if not some_condition then
      ngx.exit(ngx.HTTP_FORBIDDEN);
    end
  ';
  ...
  proxy_pass             http://$your_backend;
}
{% endhighlight %}

This way, if `some_condition` is not true, nginx will return a 403 and not do the proxy pass.
