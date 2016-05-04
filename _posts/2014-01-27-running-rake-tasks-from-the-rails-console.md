---
layout: inner
title: "running rake tasks from the rails console"
description: ""
category: ""
tags: ["ruby","rake","rails","console"]
---
{% highlight ruby %}
require "rake"
Your::Application.load_tasks
Rake::Task["some_task"].invoke # you'll also need reenable if calling a task multiple times
{% endhighlight %}
