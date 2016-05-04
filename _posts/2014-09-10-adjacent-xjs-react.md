---
layout: inner
title: "Adjacent XJS react"
description: ""
category: ""
tags: ["react", "javascript"]
---
The following error:

{% highlight bash %}
Error: Parse Error: Adjacent XJS elements must be wrapped in an enclosing tag
{% endhighlight %}

Is probably because you are not returning one element, but multiple. So, return:

{% highlight html %}
<div>
  <SomeComponent/>
  <SomeOtherComponent/>
</div>
{% endhighlight %}

instead of:


{% highlight html %}
  <SomeComponent/>
  <SomeOtherComponent/>
{% endhighlight %}
