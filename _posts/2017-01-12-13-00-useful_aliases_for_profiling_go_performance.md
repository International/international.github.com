---
layout: inner
title: useful aliases for profiling go performance
tags: ["go","golang","profiling"]
---
Problem: I never remember the syntax for <b>go tool pprof</b>

Solution: have these aliases saved.

{% highlight bash %}
alias cpuprof="go test -v -run TestMyPerf -cpuprofile cpu.out; go tool pprof mypkg.test cpu.out"
alias memprof="go test -v -run TestMyPerf -memprofile mem.out; go tool pprof --alloc_space mypkg.test mem.out"
{% endhighlight %}
