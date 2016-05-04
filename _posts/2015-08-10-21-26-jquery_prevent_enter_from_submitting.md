---
layout: inner
title: jquery prevent enter from submitting
tags: ["jquery","coffeescript"]
---
When trying to catch enter in a submit, the following did the trick:

{% highlight coffeescript %}
    $(document).on 'keypress keydown keyup', '#elem', (e) ->
     if e.which == 13
       # whatever logic
       e.preventDefault()
{% endhighlight %}
