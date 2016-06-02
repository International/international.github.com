---
layout: inner
title: preparing with automatic deployments for a new server with chef
tags: ["chef","ruby"]
---
To be able to do a pull on a newly created machine, without you having to do an initial
ssh to github, or some other host, you can do this in a recipe:

{% highlight ruby %}
ssh_known_hosts_entry 'github.com'
{% endhighlight %}

This comes from the [ssh_known_hosts](https://github.com/chef-cookbooks/ssh_known_hosts) cookbook
