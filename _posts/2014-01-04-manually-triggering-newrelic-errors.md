---
layout: inner
title: "manually triggering newrelic errors"
description: ""
category: ""
tags: ["ruby","rails","newrelic"]
---
I triggered an error manually from the console like this:
{% highlight ruby %}
irb(main):008:0> NewRelic::Agent.manual_start
=> [NewRelic::Agent::Samplers::CpuSampler, NewRelic::Agent::Samplers::MemorySampler, NewRelic::Agent::Samplers::ObjectSampler, NewRelic::Agent::
Samplers::DelayedJobSampler]
irb(main):009:0> NewRelic::Agent.notice_error(RuntimeError.new("no way"))
{% endhighlight %}

Of course when running in an environment where the agent is already started,
you'd have to call only notice_error.
