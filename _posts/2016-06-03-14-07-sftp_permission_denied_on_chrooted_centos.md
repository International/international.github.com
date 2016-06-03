---
layout: inner
title: sftp permission denied on chrooted centos
tags: ["chroot","sftp","ftp","centos","linux"]
---
I had the following in my sshd_config:

<pre>
Match Group myftpgroup
  Allowtcpforwarding no
  Chrootdirectory %h
  Forcecommand internal-sftp
</pre>

Imagining a user called john, this was the setup:

* /home was owned by root
* /home/john was owned by root 
* the files inside john were john's own files

With all these in place, this was what I received when sftp-ing:

<pre>
sftp> put README.md
Uploading README.md to /new/README.md
remote open("/new/README.md"): Permission denied
</pre>

Turns out, it was caused by selinux, and the fix was to run this cmd:
{% highlight bash %}
setsebool -P ssh_chroot_rw_homedirs on
{% endhiglight %}
