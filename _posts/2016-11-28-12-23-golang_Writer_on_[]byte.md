---
layout: inner
title: golang Writer on []byte
tags: ["go"]
---
Forgot to do a <b>Flush</b> on <b>bufio.NewWriter</b>, and the <b>Bytes()</b>
method was returning an empty slice:

{% highlight go %}
var buffer bytes.Buffer
writer := bufio.NewWriter(&buffer)
... do something with writer
// did not call <b>writer.Flush()</b>
data := writer.Bytes()
{% endhighlight %}
