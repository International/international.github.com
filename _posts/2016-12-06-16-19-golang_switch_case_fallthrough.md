---
layout: inner
title: golang switch case fallthrough
tags: ["go","golang"]
---
TIL: in go, a switch cases don't automatically fallthrough. So, this:

{% highlight go %}
switch something {
  case 0:
  case 1:
    fmt.Println("handling both")
}
{% endhighlight %}

won't work as intended. Instead it should be:
{% highlight go %}
switch something {
  case 0:
    fallthrough
  case 1:
    fmt.Println("handling both")
}
{% endhighlight %}
