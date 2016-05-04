---
layout: inner
title: environment variables in nginx
tags: ["nginx","environment","lua"]
---
Allowing environment variables in nginx config:

{% highlight nginx %}
env YOUR_BACKEND;

http {
  server {
      set_by_lua $your_backend 'return os.getenv("YOUR_BACKEND")';
      ...
      location ~ ^/to_your_backend/(.*)$ {
          set $url_full         '$1';
          ...
          proxy_pass             http://$your_backend/$url_full;
      }
  }
}
{% endhighlight %}
