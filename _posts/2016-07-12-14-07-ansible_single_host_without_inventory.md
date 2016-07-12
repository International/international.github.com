---
layout: inner
title: ansible single host without inventory
tags: ["ansible","vagrant"]
---
Example of provisioning a vagrant machine, without using Vagrantfile:

<pre>
ansible-playbook --private-key=.vagrant/machines/default/virtualbox/private_key -u vagrant -i "127.0.0.1:2200," playbooks/first.yml
</pre>
