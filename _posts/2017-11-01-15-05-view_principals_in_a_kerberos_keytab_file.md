---
layout: inner
title: view principals in a kerberos keytab file
tags: ["kerberos"]
---
Problem: want to see principals in a keytab file.

Solution:

* start <b>ktutil</b>

{% highlight bash %}
ktutil
{% endhighlight %}

* read your keytab file:

{% highlight bash %}
ktutil:  read_kt /path/to/some.keytab
{% endhighlight %}

* list them:

{% highlight bash %}
ktutil: list
{% endhighlight %}