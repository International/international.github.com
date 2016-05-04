---
layout: inner
title: "memcached SASL ubuntu"
description: ""
category: ""
tags: ["linux", "ubuntu", "memcached"]
---
{% highlight bash %}
sudo apt-get -f install libsasl2-2 sasl2-bin libsasl2-2 libsasl2-dev libsasl2-modules
{% endhighlight %}

Create a sasl folder in your user's home directory, with the following content:
{% highlight bash %}
mech_list: plain
log_level: 5
sasldb_path: /home/john/sasl/sasldb2
{% endhighlight %}

Create the database file:

{% highlight bash %}
saslpasswd2 -c -a memcached -f /home/john/sasl/sasldb2 your_user_name
{% endhighlight %}

Start memcached with SASL enabled:
{% highlight bash %}
SASL_CONF_PATH=~/sasl memcached -vvv -S 
{% endhighlight %}
