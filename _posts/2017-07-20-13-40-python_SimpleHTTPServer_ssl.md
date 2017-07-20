---
layout: inner
title: python SimpleHTTPServer ssl
tags: ["python","ssl"]
---
Problem: wanted to serve some assets over https using SimpleHTTPServer.

Solution: used info from [here](https://www.piware.de/2011/01/creating-an-https-server-in-python/) and [here](https://alexanderzeitler.com/articles/Fixing-Chrome-missing_subjectAltName-selfsigned-cert-openssl/)

* following snippet combines both of those links:

{% highlight python %}
import BaseHTTPServer, SimpleHTTPServer
import ssl

httpd = BaseHTTPServer.HTTPServer(('localhost', 8090), SimpleHTTPServer.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket (httpd.socket, certfile='~/ssl/server.crt', 
  keyfile='~/ssl/server.key', server_side=True)
httpd.serve_forever()
{% endhighlight %}