---
layout: inner
title: vagrant private network and /etc/udev/rules.d/70-persistent-net.rules
tags: ["vagrant","networking"]
---
Encountered the following error when trying to setup private networking in vagrant:

{% highlight bash %}
The following SSH command responded with a non-zero exit status.
Vagrant assumes that this means the command failed!

ARPCHECK=no /sbin/ifup eth1 2> /dev/null

Stdout from the command:

Device eth1 does not seem to be present, delaying initialization.


Stderr from the command:
{% endhighlight %}

Additionally, the call to private_network had to be changed to this:

{% highlight ruby %}
box.vm.network "private_network", ip: "whatever_ip", :auto_config => false
{% endhighlight %}

The issue is discussed [here](https://github.com/jedi4ever/veewee/issues/970), and [this post](https://github.com/jedi4ever/veewee/issues/970) helped me fix the issue.
Basically, you have to check the contents of the file: `/etc/udev/rules.d/70-persistent-net.rules`.
The safest bet seems to be to delete that file, so that vagrant can assign device names
without MAC address issues. Thanks a lot, [ablecoder](http://able.cd/)!
