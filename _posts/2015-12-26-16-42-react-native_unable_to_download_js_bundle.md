---
layout: inner
title: react-native unable to download js bundle
tags: ["ubuntu","android","react-native"]
---
Got this error message:

<pre>Unable to download JS bundle. Did you forget to start the development server or connect your device?</pre>

To fix, I had to run:

{% highlight bash %}
react-native start
{% endhighlight %}

and then click reload on the AVD.
