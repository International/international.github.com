---
layout: inner
title: chef runit example
tags: ["ruby","chef","runit"]
---
When working with the runit cookbook, you may define a service like this:

{% highlight ruby %}
runit_service "your_name" do
  env({"HOME" => "some value")
  check true
end
{% endhighlight %}

and then you'd have to create the following template files: `sv-your_name-run.erb` and
`sv-your_name-log-run.erb`. In the `sv-your_name-run` template you'd have something
like this:

{% highlight ruby %}
#!/bin/sh
exec 2>&1
exec chpst -e env echo $(date) >> log.txt
{% endhighlight %}

of course, it's kind of dumb to just echo date, and have that ran as a service, but, hey,
it's just an example. Basically, runit will make sure that the service will be restarted
if it goes down. So, in the case above, you will get a new line in log.txt about once a second.

Notice that in the runit_service I'm setting the env attribute. This has the effect of
passing environment variables to runit, which will be creted in the env folder. That's why
the `-e env` part is there in the run script.

The second template, I usually put it like this:

{% highlight ruby %}
#!/bin/sh
exec svlogd -tt ./main
{% endhighlight %}

and this will have the effect of sending the logs to `/etc/service/your_name/log/main/current`.
Of course, the logs will be sent to that file if you're logging to STDOUT. As an extra
attribute, I'm setting `check true`, which allows us to write a script to check if our
service is running. It's template file will be `sv-ruby-your_name-check.erb`. You can put anything
you want in it, as long as you exit with a 0 code. For example, to check that a ruby
script was running as a service, I did something like this:

{% highlight ruby %}
#!/bin/bash
exec &> /dev/null
exec ps aux | grep -v grep | grep -i ruby
{% endhighlight %}

If the check script writes output to stdout, it causes a write error. That's why I'm
redirecting to /dev/null.

To start/stop/restart/view logs:
{% highlight bash %}
sudo sv start your_name
sudo sv stop your_name
sudo sv restart your_name
cat /etc/service/your_name/log/main/current
{% endhighlight %}
