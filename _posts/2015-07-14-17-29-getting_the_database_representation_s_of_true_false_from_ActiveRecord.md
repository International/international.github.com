---
layout: inner
title: getting the database representation's of true/false from ActiveRecord
tags: ["ruby","rails","activerecord"]
---
Use this:

{% highlight ruby %}
ActiveRecord::Base.connection.quote(true)
{% endhighlight %}

or 
{% highlight ruby %}
ActiveRecord::Base.connection.quoted_true
{% endhighlight %}

or 

{% highlight ruby %}
ActiveRecord::Base.connection.quoted_false
{% endhighlight %}
