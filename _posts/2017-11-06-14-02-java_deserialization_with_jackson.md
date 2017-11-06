---
layout: inner
title: java deserialization with jackson
tags: ["java","jackson","spring","spring boot"]
---
Been debugging a deserialization issue for 3 hours. Turns out I had this in my code:

{% highlight java %}
import org.codehaus.jackson.annotate.JsonProperty;
import org.codehaus.jackson.annotate.JsonSetter;
{% endhighlight %}

When, it should have been these:

{% highlight java %}
import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
{% endhighlight %}