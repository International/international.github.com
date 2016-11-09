---
layout: inner
title: hiveserver2 view detailed logging output
tags: ["hive", "hadoop"]
---
Can be useful in diagnosing problems:

<pre>hive --service hiveserver2 --hiveconf hive.server2.thrift.port=10000 --hiveconf hive.root.logger=INFO,console</pre>
