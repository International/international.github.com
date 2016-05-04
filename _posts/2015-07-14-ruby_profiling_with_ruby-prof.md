---
layout: inner
title: ruby profiling with ruby-prof
tags: ["ruby","profiling","optimization"]
---
Useful snippet to profile a piece of code:

{% highlight ruby %}
RubyProf.start

results = RubyProf.stop

File.open "./#{Time.now.strftime("%H-%M-%S")}_res.html", 'w' do |file|
  RubyProf::GraphHtmlPrinter.new(results).print(file)
end
{% endhighlight %}

After this, you'll have a timestamped file containing benchmark results. Really useful for optimization.
