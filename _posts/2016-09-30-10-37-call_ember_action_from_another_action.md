---
layout: inner
title: call ember action from another action
tags: ["javascript","ember"]
---
If you need to call a different action from one of your actions, you can use:

{% highlight javascript %}
this.send("nameOfOtherAction", someArguments)
{% endhighlight %}
