---
layout: inner
title: adding a bin/console to your gem
tags: ["ruby","rubygems","gem","irb","pry"]
---
{% highlight ruby %}
#!/usr/bin/env ruby

require "bundler/setup"
require "your_lib"

require "irb"
IRB.start
{% endhighlight %}

In case you prefer pry as a console, you could replace the irb lines with this:

{% highlight ruby %}
require "pry"
Pry.start
{% endhighlight %}
