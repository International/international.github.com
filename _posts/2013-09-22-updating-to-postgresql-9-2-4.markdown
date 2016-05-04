---
layout: inner
title: updating to postgresql 9.2.4
tags: ["postgresql","pg","linux"]
---

I needed to upgrade pg_dump from 9.1.9 to 9.2.4, so I needed to upgrade postgresql as well.
I found the solution [here](http://askubuntu.com/questions/287786/how-to-install-postgresql-on-ubuntu-13-04).

{% highlight bash linenos %}
wget --quiet -O - http://apt.postgresql.org/pub/repos/apt/ACCC4CF8.asc | sudo apt-key add -
sudo apt-get update
sudo apt-get install postgresql-common -t raring
sudo apt-get install postgresql-9.2
{% endhighlight %}

After doing this I symlinked the 9.2.4 pg_dump version:

{% highlight bash %}
sudo ln -s /usr/lib/postgresql/9.2/bin/pg_dump /usr/bin/pg_dump
{% endhighlight %}
