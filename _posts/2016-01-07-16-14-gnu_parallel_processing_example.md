---
layout: inner
title: gnu parallel processing example
tags: ["ruby","parallel"]
---
Here's a simple script that *does* work :)

{% highlight ruby %}
def do_work(argument)
  puts "#{Time.now} Processing #{argument}"
  sleep 5
  puts "#{Time.now} Done processing #{argument}"
end

do_work(ARGV[0])
{% endhighlight %}

When running this, I get the following output on my machine:

{% highlight bash %}
2016-01-07 16:14:18 +0100 Processing 1
2016-01-07 16:14:23 +0100 Done processing 1
{% endhighlight %}

If we would like to run the script for more than one argument, it would take us some time.
Luckily, we can use [parallel](http://www.gnu.org/software/parallel/) to speed things up:

{% highlight bash %}
$ (echo 1;echo 2) | parallel -j+0 --eta 'ruby initial.rb {}'
When using programs that use GNU Parallel to process data for publication please cite:

  O. Tange (2011): GNU Parallel - The Command-Line Power Tool,
  ;login: The USENIX Magazine, February 2011:42-47.

This helps funding further development; and it won't cost you a cent.
Or you can get GNU Parallel without this requirement by paying 10000 EUR.

To silence this citation notice run 'parallel --bibtex' once or use '--no-notice'.


Computers / CPU cores / Max jobs to run
1:local / 4 / 2

Computer:jobs running/jobs completed/%of started jobs/Average seconds to complete
ETA: 0s Left: 2 AVG: 0.00s  local:2/0/100%/0.0s 2016-01-07 16:12:02 +0100 Processing 1
2016-01-07 16:12:07 +0100 Done processing 1
2016-01-07 16:12:02 +0100 Processing 2
2016-01-07 16:12:07 +0100 Done processing 2
ETA: 0s Left: 0 AVG: 0.00s  local:0/2/100%/2.5s
{% endhighlight %}

As we can see from the timestamps, both scripts started at the same time, and were executed in parallel.
The -j+0 flag tells parallel to use as many cores as possible to complete the jobs.

Alternatively, this has a simpler syntax:

{% highlight bash %}
parallel --eta -j+0 'ruby initial.rb {}' ::: 1 2
{% endhighlight %}

or, with a shell glob:

{% highlight bash %}
parallel --eta -j+0 'ruby initial.rb {}' ::: *.txt
{% endhighlight %}

More examples [here](http://www.shakthimaan.com/posts/2014/11/27/gnu-parallel/news.html)

GNU citation:

```
O. Tange (2011): GNU Parallel - The Command-Line Power Tool, The USENIX Magazine, February 2011:42-47.
```
