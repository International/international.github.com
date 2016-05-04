---
layout: inner
title: recursively replace with ack-grep and a little help from ruby and xargs
tags: ["ack"]
---

ack-grep is a great tool for finding files containing some strings. It's easy to combine it's power, with the in-place editing capabilities of ruby. Here's how we'd replace the string alert with # alert in all the files within a directory:

{% highlight bash %}
ack-grep -lia alert app | xargs ruby -pi -e '$_.gsub!("alert","#alert")'
{% endhighlight %}

What's important to notice here, is that we're using the l switch for ack, which means "give me only the name of the files containing that string", and we're piping each file to ruby's in place editing. If you'd like to create a backup of the file, run the command like this:

{% highlight bash %}
ack-grep -lia alert app | xargs ruby -pi.bak -e '$_.gsub!("alert","#alert")'
{% endhighlight %}

This way, you'd also have the old file, in case your replace did not go so well.
