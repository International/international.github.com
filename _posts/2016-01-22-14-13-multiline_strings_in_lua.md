---
layout: inner
title: multiline strings in lua
tags: ["lua"]
---
{% highlight bash %}
LuaJIT 2.1.0-alpha -- Copyright (C) 2005-2014 Mike Pall. http://luajit.org/
JIT: ON SSE2 SSE3 SSE4.1 fold cse dce fwd dse narrow loop abc sink fuse
> a=[[
>> one
>> two
>> three
>> ]]
> print(a)
one
two
three

>

{% endhighlight %}
