---
layout: inner
title: random examples of working with knife zero
tags: ["chef","ruby","knife","knife-zero","chef-zero"]
---
This is the config I've used with knife zero:

{% highlight ruby %}
current_dir = File.dirname(__FILE__)
chef_zero.enabled        true
local_mode               true

log_level                :info
log_location             STDOUT

node_name                "local"
client_key               "#{current_dir}/local.pem"

chef_server_url          "http://127.0.0.1:9901"
cache_type               'BasicFile'
cache_options( :path => "#{ENV['HOME']}/.chef/checksums" )

cookbook_path [
  File.join(current_dir, '..', 'cookbooks'),
  File.join(current_dir, '..', 'site-cookbooks'),
  File.join(current_dir, '..', 'berks-cookbooks')
]

{% endhighlight %}

* configure a kitchen.yml file with a private network ( I will assume 192.168.33.15 is the IP ):
<pre>
kitchen create
</pre>

* bootstrap it with `knife zero`:
<pre>
knife zero bootstrap 192.168.33.15 --ssh-user vagrant --sudo
</pre>

* this will ask you for your password, then proceed with the installation of chef
* after the command is done, you should have a new entry in the nodes folder