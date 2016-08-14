---
layout: inner
title: hive and beeline in cloudera
tags: ["hive","beeline","cloudera"]
---
Quickstart:
* start beeline
<pre>beeline</pre>
* connect
<pre>beeline> !connect jdbc:hive2://localhost:10000 username password org.apache.hive.jdbc.HiveDriver</pre>
* to insert data from a local file, the file had to have `755` permissions on it, and be owned by the user cloudera (?)
* if it did not have that, I kept getting errors like this:
<pre>
0: jdbc:hive2://localhost:10000> LOAD DATA LOCAL INPATH '/user/cloudera/employee.txt' overwrite into table employee;
Error: Error while compiling statement: FAILED: SemanticException Line 1:23 Invalid path ''/user/cloudera/employee.txt'': No files matching path file:/user/cloudera/employee.txt (state=42000,code=40000)
</pre>
