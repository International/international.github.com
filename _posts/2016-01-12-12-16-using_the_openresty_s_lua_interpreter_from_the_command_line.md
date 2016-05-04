---
layout: inner
title: using the openresty's lua interpreter from the command line
tags: ["nginx","openresty","lua"]
---
If you're using openresty, you can make use of the lua interpreter from the command
line if you need to test some code that will be ran in nginx. On my pc, I can access
the interpreter ( REPL mode ) like this:

{% highlight bash %}
./build/luajit-root/usr/local/openresty/luajit/bin/luajit-2.1.0-alpha
{% endhighlight %}

And if you need to execute a file, do:

{% highlight bash %}
./build/luajit-root/usr/local/openresty/luajit/bin/luajit-2.1.0-alpha file/path/goes/here
{% endhighlight %}
