---
layout: inner
title: chef remote_file progress
tags: ["chef"]
---
When using <b>remote_file</b>, the <b>show_progress</b> attribute can show download progress.
Example:

<pre>
remote_file "/home/vagrant/zeppelin-0.6.2-bin-all.tgz" do
  source "http://mirrors.m247.ro/apache/zeppelin/zeppelin-0.6.2/zeppelin-0.6.2-bin-all.tgz"
  show_progress true
end
</pre>

and it will have the following effect on the output:

<pre>
==> default:  - Progress: 10%
==> default:
==> default:  - Progress: 20%
...
</pre>
