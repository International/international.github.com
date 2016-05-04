---
layout: inner
title: test that rails controller rendered a specific layout
tags: ['ruby','rails','rspec']
---
{% highlight ruby %}
get :index
expect(response).to have_rendered(layout: 'main')
{% endhighlight %}
