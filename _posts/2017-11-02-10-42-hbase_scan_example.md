---
layout: inner
title: hbase scan example
tags: ["java","hbase"]
---
Problem: want to perform a scan on a table, and list the values.

Solution:

{% highlight java %}
        Scan scan = new Scan();
        ResultScanner scanner = someTable.getScanner(scan);

        for(Result row : scanner) {

            List<Cell> cells = row.listCells();

            for(Cell c : cells) {
                byte columnFamily[] = CellUtil.cloneFamily(c);
                byte qualifier[] = CellUtil.cloneQualifier(c);
                byte value[] = CellUtil.cloneValue(c);
                System.out.println("columnFamily:" + Bytes.toString(columnFamily) +
                    " qualifier:" + Bytes.toString(qualifier) + " value:" + Bytes.toString(value));
            }
        }
{% endhighlight %}