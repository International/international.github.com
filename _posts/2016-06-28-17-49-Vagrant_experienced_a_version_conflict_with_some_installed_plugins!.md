---
layout: inner
title: Vagrant experienced a version conflict with some installed plugins!
tags: ["ruby","vagrant"]
---
Tried to run `vagrant ssh-config` from within capistrano, and got the following
message:

<pre>
Vagrant experienced a version conflict with some installed plugins!
This usually happens if you recently upgraded Vagrant. As part of the
upgrade process, some existing plugins are no longer compatible with
this version of Vagrant. The recommended way to fix this is to remove
your existing plugins and reinstall them one-by-one. To remove all
plugins: ...
</pre>

Turns out, the fix was to run `ssh-config` like this:

<pre>
system("unset RUBYLIB; vagrant ssh-config")
</pre>

Found the solution [here](https://github.com/mitchellh/vagrant/issues/6158)
