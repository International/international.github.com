---
layout: inner
title: rbtrace ruby debug production mode
tags: ["ruby","rails","debug"]
---
This can be useful to debug something in production mode. Add rbtrace to your Gemfile, and start your servers in production mode. In my case, I needed to debug unicorn in production mode. I did a:

{% highlight bash %}
ps aux | grep -i unicorn
{% endhighlight %}

And found the pid of one of the unicorn worker process. After that:

{% highlight bash %}
bundle exec rbtrace -p YOUR_PID -e 'puts Rails.env' # or whatever code you want
{% endhighlight %}
