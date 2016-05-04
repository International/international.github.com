---
layout: inner
title: tracing ruby method calls
tags: ["ruby","debug"]
---
The following allowed me to identify the cause of an infinite loop:

{% highlight ruby %}
tp = TracePoint.new(:call) do |x|
  puts x.inspect
end
tp.enable
# add your own code here
tp.disable
{% endhighlight %}

When running this, you'll see all method calls in between enable/disable of the tracepoint
