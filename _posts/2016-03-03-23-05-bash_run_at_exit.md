---
layout: inner
title: bash run at exit
tags: ["bash"]
---

{% highlight bash %}
#!/bin/bash -e
echo "do work here"
start_some_binary_in_background &
jobpid=$!

function cleanup {
  echo
  echo "Killing $jobpid"
  kill $jobpid
}

echo "do more work"
trap cleanup EXIT

echo "also more work"
{% endhighlight %}
