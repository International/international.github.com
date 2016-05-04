---
layout: inner
title: iterating over the contents of a data bag
tags: ["ruby","chef"]
---
{% highlight ruby %}
data_bag("users").each do |user_name|
  user_data = data_bag_item("users", user_name)
  log "found #{user_data.id}"
end
{% endhighlight %}

or, using search:

{% highlight ruby %}
search("users","*:*").each do |user|
  log "found #{user.id}"
end
{% endhighlight %}
