---
layout: inner
title: upgrade pg client to 9.4 in trusty
tags: ["ubuntu","postgres"]
---
{% highlight bash %}
su
add-apt-repository "deb https://apt.postgresql.org/pub/repos/apt/ trusty-pgdg main"
wget --quiet -O - https://postgresql.org/media/keys/ACCC4CF8.asc | apt-key add - 
apt-get update
apt-get install postgresql-client-9.4
{% endhighlight %}
