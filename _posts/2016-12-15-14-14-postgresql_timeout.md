---
layout: inner
title: postgresql timeout
tags: ["postgresql"]
---
Found a really interesting bit of info in [this heroku article](https://blog.heroku.com/postgres-essentials?utm_source=postgresweekly&utm_medium=email):

{% highlight sql %}
SET statement_timeout TO '30s';
{% endhighlight %}

Which would ensure your queries will be timed out after 30 seconds.
