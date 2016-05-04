---
layout: inner
title: "debugging go with gdb"
description: ""
category: ""
tags: ["golang", "gdb"]
---
{% highlight bash %}
go build -gcflags "-N -l" your_file.go
{% endhighlight %}

and after that:

{% highlight bash %}
gdb ./your_file
{% endhighlight %}

Continue with info from [here](http://golang.org/doc/gdb)
