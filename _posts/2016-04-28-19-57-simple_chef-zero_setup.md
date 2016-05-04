---
layout: inner
title: simple chef-zero setup
tags: ruby chef chef-zero
---
Simple Chef-Zero setup
===============
* create folder `my-chef`
* `cd my-chef`
* `mkdir .chef`
* `ssh-keygen -f local.pem -P ""`
* `ssh-keygen -f validation.pem -P ""`
* add this to knife.rb:

{% highlight ruby %}
current_folder = File.dirname(__FILE__)

chef_repo      = File.join(current_folder, "..")
chef_server_url  "http://127.0.0.1:9901"
node_name        "local"
client_key        File.join(current_folder, "local.pem")

cookbook_path    "#{chef_repo}/cookbooks"
cache_type       "BasicFile"
cache_options    :path => "#{chef_repo}/checksums"

{% endhighlight %}

* make sure to start chef-zero on port 9901:
<pre>chef-zero --port 9901</pre>

* make a folder `nodes` at the same level as .chef
* add some data in `nodes/lenode.json`

{% highlight json %}
{
  "name": "lenode",
  "chef_type": "node",
  "json_class": "Chef::Node",
  "chef_environment": "_default",
  "run_list": [
    "recipe[whatever]",
  ],
  ...
}
{% endhighlight %}

* cd to same level as nodes folder
* upload them:

{% highlight bash %}
knife upload nodes
{% endhighlight %}
