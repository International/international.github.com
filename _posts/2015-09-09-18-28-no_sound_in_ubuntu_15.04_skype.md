---
layout: inner
title: no sound in ubuntu 15.04 skype
tags: ['skype','ubuntu','linux','sound']
---
Had no sound after a fresh install, and everything in skype was showing virtual device.
This is what fixed it:

{% highlight bash %}
sudo apt-get install pulseaudio
sudo apt-get install libpulse0:i386
{% endhighlight %}
