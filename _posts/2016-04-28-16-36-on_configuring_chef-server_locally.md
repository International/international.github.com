---
layout: inner
title: on configuring chef-server locally
tags: ["ruby","chef", "chef-server"]
---
After reading [this article](http://misheska.com/blog/2014/11/25/chef-server-12/), not all the steps applied in my case. I've pasted the steps I had to take on my system, to get <b>chef-dk 0.11.0-1_amd64.deb</b> to work on the listed CentOs.

Steps to configure chef-server locally
=======================================
- use this kitchen file:

{% highlight yaml %}
---
driver:
  name: vagrant

provisioner:
  name: chef_zero

# Uncomment the following verifier to leverage Inspec instead of Busser (the
# default verifier)
# verifier:
#   name: inspec

platforms:
- name: centos65
  driver:
    network:
      - ["private_network", {ip: "192.168.33.7"}]
    box: learningchef/centos65
    box_url: learningchef/centos65

suites:
  - name: default
    run_list:
      - recipe[chef-server::default]
    attributes:
{% endhighlight %}

- with this recipe:

{% highlight ruby %}
#
# Cookbook Name:: chef-server
# Recipe:: default
#
# Copyright (c) 2016 The Authors, All Rights Reserved.
default['chef-server']['url'] = 'https://opscode-omnibus-packages.s3.amazonaws.com/el/6/x86_64/chef-server-11.1.4-1.el6.x86_64.rpm'

package_url = node['chef-server']['url']
package_name = ::File.basename(package_url)
package_local_path = "#{Chef::Config[:file_cache_path]}/#{package_name}"

# omnibus_package is remote (i.e., a URL) let's download it
rpm_package package_name do
  source package_local_path
end

package package_local_path
# reconfigure the installation
execute 'chef-server-ctl reconfigure'
{% endhighlight %}

- do a <b>kitchen converge</b>
- open https://192.168.33.7 in your browser
- login as admin/p@assw0rd1
- change password
- create user geo with whatever password, make him an admin
- on the screen there will be a private key shown, copy that and save it locally
to a file named geo.pem
- in <b>chef-repo</b> create a <b>.chef</b> folder
- copy geo.pem there
- we need to copy the chef-validator from the host <b>scp root@192.168.33.7:/etc/chef-server/chef-validator.pem .</b>
- add a new entry to <b>/etc/hosts</b> with the following content:

<pre>
192.168.33.7 default-centos65
</pre>

- in .chef, create a file called <b>knife.rb</b> with the following content:
{% highlight ruby %}
current_dir = File.dirname(__FILE__)
log_level :info
log_location STDOUT
node_name "geo"
client_key "#{current_dir}/geo.pem"
validation_client_name "chef-validator"
validation_key "#{current_dir}/chef-validator.pem"
chef_server_url "https://default-centos65:443"
cache_type "BasicFile"
cache_options( :path => "#{ENV['HOME']}/.chef/checksums" )
cookbook_path ["#{current_dir}/../cookbooks"]
{% endhighlight %}

- the content of .chef should be:
<pre>
.chef/
  geo.pem
  knife.rb
  chef-validator.pem
</pre>

- test that server is accessible:
<pre>
$ knife client list
chef-validator
chef-webui
</pre>

Potential problems you may encounter
====================================
* I encountered this initially, when running <b>knife client list</b>
<pre>
➜  chef-repo git:(master) ✗ knife client list
ERROR: SSL Validation failure connecting to host: default-centos65.vagrantup.com - hostname "default-centos65.vagrantup.com" does not match the server
certificate
ERROR: SSL Error connecting to https://default-centos65.vagrantup.com/clients, retry 1/5
ERROR: SSL Validation failure connecting to host: default-centos65.vagrantup.com - hostname "default-centos65.vagrantup.com" does not match the server
certificate
</pre>

  * fix for it was to run <b>knife ssl check</b>, where I saw the following:
  <pre>
  ➜  chef-repo git:(master) ✗ knife ssl check
Connecting to host default-centos65.vagrantup.com:443
ERROR: The SSL cert is signed by a trusted authority but is not valid for the given hostname
ERROR: You are attempting to connect to:   'default-centos65.vagrantup.com'
ERROR: The server's certificate belongs to 'default-centos65'
</pre>
  * then, I just changed the entry from default-centos65.vagrantup.com to default-centos65 in /etc/hosts
  * also, I went in knife.rb and I changed <b>chef_server_url</b> from:
  ```{ruby}
  chef_server_url "https://default-centos65.vagrantup.com:443"
  ```
  to
  ```{ruby}
  chef_server_url "https://default-centos65:443"
  ```
  * next, run another <b>knife ssl check</b>, and you should see this output:
  <pre>
  ➜  chef-repo git:(master) ✗ knife ssl check
  Connecting to host default-centos65:443
  Successfully verified certificates from `default-centos65'
  </pre>
