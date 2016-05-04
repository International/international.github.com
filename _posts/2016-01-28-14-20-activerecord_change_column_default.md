---
layout: inner
title: activerecord change column default
tags: ["ruby","rails","activerecord","migration"]
---
{% highlight ruby %}
def change
  change_column_default :users, :name, "anonymous"
end
{% endhighlight %}
