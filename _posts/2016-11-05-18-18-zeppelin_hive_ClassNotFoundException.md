---
layout: inner
title: zeppelin hive ClassNotFoundException
tags: ["zeppelin", "hive", "hadoop"]
---
Got a `ClassNotFoundException` when trying to create a snippet in a zeppelin notebook.
Luckily I found [this answer](http://stackoverflow.com/a/38726719/31610), which suggested this fix:

* download hive-jdbc standalone.jar
* download hadoop-common.jar
* place them in `interpreter/jdbc` of the zeppelin installation:

<pre>
cp ~/Dev/Hadoop/apache-hive-1.2.1-bin/lib/hive-jdbc-1.2.1-standalone.jar ./interpreter/jdbc/
cp ~/Dev/Hadoop/hadoop-2.6.4/share/hadoop/common/hadoop-common-2.6.4.jar ./interpreter/jdbc/
</pre>
