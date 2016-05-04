---
layout: inner
title: ruby command line filtering
tags: ["ruby","command line"]
---
I wanted to be able to filter the API calls that took longer than 1200ms from a logentries log. Lines are in the following format:

{% highlight bash %}
477 <158>1 2015-06-30T11:24:39.500691+00:00 heroku router - - at=info method=GET path="/api/users.json" host=somehost.name request_id=9bb472d6-c5d7-4f91-95ef-e795cf5ce17f fwd="some.ip.here" dyno=web.1 connect=2ms service=5815ms status=200 bytes=65518
{% endhighlight %}

The following does the trick:

{% highlight ruby %}
ruby -n -e 'print $_ if $_[/service=(\d+)/,1].to_f > 1200' < log_entries.log
{% endhighlight %}
