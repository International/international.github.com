---
layout: inner
title: java hbase obtain column families for a particular table
tags: ["java","hbase"]
---
Problem: want to obtain column families for a particular table.

Solution:

{% highlight java %}
        Connection conn = ConnectionFactory.createConnection(config);
        Admin admin = conn.getAdmin();

        TableName transactionTableName = TableName.valueOf("some_table");
        HTableDescriptor tableDesc = admin.getTableDescriptor(transactionTableName);
        // this will contain all of them
        HColumnDescriptor colDesc[] = tableDesc.getColumnFamilies();
{% endhighlight %}