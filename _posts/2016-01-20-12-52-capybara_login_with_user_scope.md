---
layout: inner
title: capybara login with user scope
tags: ["ruby","rails","rspec","devise"]
---
If you have more than one user type, you can login as it from your tests using:

{% highlight ruby %}
login_as(your_user, :scope => :admin)
{% endhighlight %}

The example above assumes that you also have an admin mapping, besides your User mapping.
