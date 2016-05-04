---
layout: inner
title: javascript arguments array
tags: ["coffeescript", "javascript"]
---

It may be useful to know that each javascript function has access to a variable/property called arguments. It will hold all the arguments that the function was called with. An example where this may prove useful:

- imagine you're doing a variable number of ajax requests in the same time, and you need to do something with the results, after they've all finished. I'm using the Deferred and Promise objects from jQuery, which I'm resolving once I've parsed the response of the server. Typically, the then function would take as many parameters as there were requests ... and to write .then (data) it would only pass the first result to the then function. And ... this is where arguments comes into play.

{% highlight coffeescript %}
urls   = ["url1","url2", ...] # it may have any number
deferreds = urls.map (e) -> 
  dfrd    = $.Deferred()
  $.get e, (server_response) ->
    your_val  = $("some_selector",server_response).html()
    dfrd.resolve(your_val)
  dfrd.promise()

$.when.apply(this, deferreds).then ->
  console.log "All AJAX requests finished"
  console.log arguments
{% endhighlight %}
