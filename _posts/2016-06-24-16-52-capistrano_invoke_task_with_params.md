---
layout: inner
title: capistrano invoke task with params
tags: ["capistrano","ruby"]
---
Imagine the following code:
{% highlight ruby %}
desc "callerz"
task :callerz do
  set :branch, "potato"
  called_with
end

desc "called with param"
task :called_with do
  br = variables[:branch]
  puts "called with branch #{br}"
end
{% endhighlight %}

In the above example, the task <b>called_with</b> can be passed a param form the command line:

<pre>
cap called_with -s branch=master
</pre>

And this allows us to send a parameter, from the commandline. The output will be:

<pre>
called with branch master
</pre>

 If you need to call a parameterized task from another task,
the above task, <b>callerz</b> illustrates one way of getting that result. You invoke it like this:

<pre>
cap callerz
</pre>

And inside, it sets the variable branch to a value. The output will be:

<pre>
called with branch potato
</pre>

The above code was tested in capistrano 2.15.8
