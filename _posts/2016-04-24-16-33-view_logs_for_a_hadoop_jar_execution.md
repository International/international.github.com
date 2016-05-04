---
layout: inner
title: view logs for a hadoop jar execution
tags: ["hadoop"]
---
Assuming you started a jar with `hadoop jar your_jar some_args`, you might see 
something similar to this, when the job starts:

<pre>
16/04/24 07:33:01 INFO mapreduce.Job: The url to track the job: http://quickstart.cloudera:8088/proxy/application_1461491409387_0019/
</pre>

Once the job finishes, a way to view your logs is by issuing this command:

{% highlight bash %}
yarn logs -applicationId 1461491409387_0019
{% endhighlight %}

I found this helpful for viewing log statements that I added in the jar.
