---
layout: inner
title: Parse data from another agent using WebsiteAgent in huginn
tags: ["huginn","ruby"]
---
For one use case, I had to use a PostAgent, which returned html, and I wanted to parse
it's created events with a WebsiteAgent. Here's the code that made it happen:

PostAgent:

{% highlight javascript %}
{
  "post_url": "URL to which you're posting goes here",
  "expected_receive_period_in_days": "1",
  "content_type": "form",
  "method": "post",
  "payload": {
    "param1": 1,
    "param2": "2"
  },
  "headers": {},
  "emit_events": "true"
}
{% endhighlight %}

The events created by the previous agent will be in the
following format:

{% highlight javascript %}
{
  "body": "<html>...</html>",
  "headers": {
    "Date": "Fri, 01 Apr 2016 06:02:40 GMT",
    "Pragma": "no-cache",
    "Expires": "Thu, 01 Jan 1970 00:00:00 GMT",
    // other headers cut for brevity
  },
  "status": 200
}
{% endhighlight %}

And here's the WebsiteAgent that parses the events created ( make sure to set
the previous agent as source ):

{% highlight javascript %}
{
  "expected_update_period_in_days": "2",
  "data_from_event": "{{body}}",
  "type": "html",
  "mode": "on_change",
  "extract": {
    "status": {
      "css": "table tbody td:nth-child(3)",
      "value": ".//text()"
    }
  }
}
{% endhighlight %}
