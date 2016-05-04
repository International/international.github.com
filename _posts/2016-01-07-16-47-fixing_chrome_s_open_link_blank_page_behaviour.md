---
layout: inner
title: fixing chrome's open link blank page behaviour
tags: ["chrome","ubuntu"]
---
Whenever clicking a link that should open chrome, like for example a link
in the terminal, I always got a new chrome window with a blank page.
Turns out it was because the file in $HOME/.local/share/applications/google-chrome.desktop
had this as Exec:

{% highlight bash %}
Exec=/opt/google/chrome/chrome
{% endhighlight %}

instead of:
{% highlight bash %}
Exec=/opt/google/chrome/chrome %U
{% endhighlight %}

Found the answer [here](http://askubuntu.com/questions/689449/15-10-chrome-external-links-are-opened-as-blank-tabs-in-new-browser-window)
