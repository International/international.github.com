---
layout: inner
title: fixing hibernation on ubuntu 16.04
tags: ["ubuntu","linux","gnome-shell"]
---
After updating to 16.04, hibernation wasn't working anymore. Found the fix
[here](http://askubuntu.com/questions/761820/suspend-not-working-on-ubuntu-16-04-for-dell-3537):

{% highlight bash %}
$ cd /tmp
$ wget \
kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.8-wily/linux-headers-4.4.8-040408_4.4.8-040408.201604200335_all.deb \
kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.8-wily/linux-headers-4.4.8-040408-generic_4.4.8-040408.201604200335_amd64.deb \
kernel.ubuntu.com/~kernel-ppa/mainline/v4.4.8-wily/linux-image-4.4.8-040408-generic_4.4.8-040408.201604200335_amd64.deb

$ sudo dpkg -i linux-headers-4.4*.deb linux-image-4.4*.deb
{% endhighlight %}

Reboot, and then it should work again.
