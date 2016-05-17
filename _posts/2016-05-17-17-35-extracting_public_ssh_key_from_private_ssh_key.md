---
layout: inner
title: extracting public ssh key from private ssh key
tags: ["ssh"]
---
In which we find out how to extract a public key from a private key file
<!-- exce -->
{% highlight bash %}
ssh-keygen -y -f your_pem_file.pem
{% endhighlight %}
