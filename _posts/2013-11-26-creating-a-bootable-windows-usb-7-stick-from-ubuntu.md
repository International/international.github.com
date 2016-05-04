---
layout: inner
title: "creating a bootable windows usb 7 stick from ubuntu"
description: ""
category: ""
tags: ["ubuntu", "windows", "usb"]
---
Simple way to create a bootable windows 7 usb stick:

{% highlight bash %} 
$ sudo apt-get update 
$ sudo apt-get install gparted 
{% endhighlight %}

* insert usb stick
* format it as ntfs using gparted
* in case it's not set, go to Manage flags, and check boot

{% highlight bash %}
$ sudo mount -o loop your.iso /some/mount/point 
{% endhighlight %}

* insert usb stick again so your ubuntu can remount it

{% highlight bash %} 
$ cp -rvf /some/mount/point/* /media/wherever/ubuntu/mounted_it 
{% endhighlight %}

* wait
* reboot, and select your usb stick as your first boot device
* done!