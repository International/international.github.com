---
layout: inner
title: timeout command from the shell
tags: ["shell","bash","timeout"]
---
{% highlight bash %}
[vagrant@localhost ~]$ timeout 2 bash -c "sleep 3"
[vagrant@localhost ~]$ echo $?
124
{% endhighlight %}
