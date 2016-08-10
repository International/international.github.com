---
layout: inner
title: parallelize with space separated args with xargs
tags: ["xargs","parallel"]
---
<pre>
echo "node0 node1 node2" | xargs -n1 -P3 -IBOX sh -c 'vagrant halt BOX'
</pre>
