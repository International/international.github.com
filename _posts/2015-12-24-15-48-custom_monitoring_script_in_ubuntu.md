---
layout: inner
title: custom monitoring script in ubuntu
tags: ["ubuntu","linux","monitor"]
---
I wanted to have a way to run something repeteadly and have visual output. Luckily,
[indicator-sysmonitor](https://github.com/fossfreedom/indicator-sysmonitor) allows you to
configure a custom script. You can get a full description on how to do it, by following
[this link](http://askubuntu.com/questions/100306/is-an-internet-connectivity-indicator-applet-available-for-the-unity-panel/453975#453975).
I find it very useful to add the path to a shell script, and to call whatever you want from it.
So, I've configured the path to be <i>~/code/sysmenu.sh</i>, with the following contents:

{% highlight bash %}
#!/bin/bash
source ~/.rvm/scripts/rvm
cd /home/geo/code/sidekiq-queue
ruby extractor.rb
{% endhighlight %}

Of course, you could run node, python, c, or whatever your heart desires. Just make sure your
script produces output, since that output is what's displayed by your indicator. For example, I use
it to keep track of the jobs that are being processed by sidekiq, how many jobs are running, how
many are queued, and how many scheduled. 
