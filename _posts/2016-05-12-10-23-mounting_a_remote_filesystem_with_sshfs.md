---
layout: inner
title: mounting a remote filesystem with sshfs
tags: ["sshfs"]
---
How to mount a remote filesystem with sshfs
<!-- exce -->

{% highlight ruby %}
sudo sshfs -o allow_other,defer_permissions root@your-ip:/ /your/local/mount/point
{% endhighlight %}
