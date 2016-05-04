---
layout: inner
title: "long lines in csv reports"
description: ""
category: ""
tags: ["ruby","rails","csv","encoding"]
---
If you have some CSV files with lines longer than 990 chars, you may notice that
the lines get splitted. The fix is described [here](http://stackoverflow.com/questions/10401863/how-to-send-a-csv-attachment-with-lines-longer-than-990-characters)

{% highlight ruby %}
attachments['file.csv']= { :data=> ActiveSupport::Base64.encode64(@string), :encoding => 'base64' }
{% endhighlight %}
