---
layout: inner
title: libv8 and ruby racer issues on el capitan
tags: ["ruby","os x", "gem","v8"]
---
Found solution for bundling problems [here](http://stackoverflow.com/a/36271579):

{% highlight bash %}
brew tap homebrew/versions
brew install v8-315

gem install libv8 -v '3.16.14.13' -- --with-system-v8
gem install therubyracer -- --with-v8-dir=/usr/local/opt/v8-315

bundle install
{% endhighlight %}
