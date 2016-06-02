---
layout: inner
title: ssh agent forwarding for github deploys
tags: ["ssh","forward","github", "capistrano"]
---
Kept receiving these errors when trying to execute commands on the remote server, ( with capistrano )
while forwarding my SSH keys:

<pre>
Permission denied (publickey).
fatal: Could not read from remote repository.

Please make sure you have the correct access rights
and the repository exists.
</pre>

[This link](http://stackoverflow.com/a/23171517/31610) offered the solution. I ran:

{% highlight bash %}
$ ssh-add -L
{% endhighlight %}

and the output was:
<pre>
The agent has no identities.
</pre>

Then, I've added the key:

{% highlight bash %}
$ ssh-add ~/.ssh/id_rsa
{% endhighlight %}

Confirmed that it appears in ssh-add -L:

{% highlight bash %}
$ ssh-add -L
ssh-rsa ....
{% endhighlight %}

And after this step, the ssh key was properly forwarded.
