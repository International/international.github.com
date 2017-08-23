---
layout: inner
title: osx capture output of running process
tags: ["osx","mac","stdout","process","capture"]
---
Problem: want to capture output of an already running process

Solution: found [here](https://stackoverflow.com/a/33844061/31610)

{% highlight bash %}
capture() {
    sudo dtrace -p "$1" -qn '
        syscall::write*:entry
        /pid == $target && arg0 == 1/ {
            printf("%s", copyinstr(arg1, arg2));
        }
    '
}
{% endhighlight %}

after, run:

* save above code in <b>capture.sh</b>
* <b>source capture.sh</b>
* capture PID