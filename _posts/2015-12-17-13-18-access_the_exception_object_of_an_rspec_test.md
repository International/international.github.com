---
layout: inner
title: access the exception object of an rspec test
tags: ['ruby','rspec','rails']
---
Add an after hook like this:

{% highlight ruby %}
after(:each) do |ex|
  if ex.example.exception
    # do smth here
  end
end
{% endhighlight %}
