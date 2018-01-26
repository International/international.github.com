---
layout: inner
title: postgres reassign owned
tags: ["postgres","pg","sql"]
---
Problem: wanted to change everything owned by one user to be owned by another user.

Solution: found [here](https://stackoverflow.com/a/13535184/31610)

{% highlight sql %}
reassign owned by old_user_name to new_user_name;
{% endhighlight %}