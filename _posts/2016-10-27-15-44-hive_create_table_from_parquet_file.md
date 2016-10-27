---
layout: inner
title: hive create table from parquet file
tags: ["hive","hdfs"]
---
Example of creating a hive table from a parquer file:
<pre>
create external table mytable (name string,id int) STORED AS PARQUET LOCATION '/user/me/file.parquet'
</pre>
