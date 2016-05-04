---
layout: inner
title: migrating stuff from the commandline
tags: ["ruby","activerecord","migration"]
---
Can be ran from the rails console:

{% highlight ruby %}
ActiveRecord::Migration.remove_column(:users,:email)
{% endhighlight %}
