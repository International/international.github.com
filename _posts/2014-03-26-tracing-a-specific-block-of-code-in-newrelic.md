---
layout: inner
title: "tracing a specific block of code in NewRelic"
description: ""
category: ""
tags: ["ruby"]
---
Info available [here](https://docs.newrelic.com/docs/ruby/ruby-custom-metric-collection).
Make sure you extend the NewRelic tracer in your class:

{% highlight ruby %}

extend ::NewRelic::Agent::MethodTracer

self.class.trace_execution_scoped("Custom/Some/TransactionName") do
  # traced code
end

{% endhighlight %}
