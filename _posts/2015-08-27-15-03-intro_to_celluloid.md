---
layout: inner
title: intro to celluloid
tags: ["ruby","celluloid"]
---
I started to use Celluloid to speed up some tasks, and I like the simplicity it offers.

For my specific case, I wanted to speed up some downloads/uploads. Here's how you do it:

{% highlight ruby %}
class YourClass
  include Celluloid

  def do_something
    # code that takes a long time goes here
  end
end
{% endhighlight %}

So, if we do this:

{% highlight ruby %}
YourClass.new.do_something
{% endhighlight %}

<del>will not use Celluloid</del> the Ruby main thread becomes an Actor.
When you write <b>YourClass.new.do_something</b> you are actually doing a blocking call and invoking <b>do_something</b>
on the YourClass instance (which is contained within an actor). **Thanks for the correction Sam Williams**

{% highlight ruby %}
YourClass.new.future.do_something
{% endhighlight %}

This returns a future, which may or may not be completed at the time you need it.
Each future exposes a method called **ready?** which you can call to see if the future finished.

So, you could do something like this:

{% highlight ruby %}
ftr = YourClass.new.future.do_something

loop do
  break if ftr.ready?
  sleep 0.1
end

puts ftr.value
{% endhighlight %}

Now, imagine that you want to run your processing using multiple threads at the same time.
You could create a number of futures inside a loop, and then check them all with ready? and append
to them once some of them have finished. I did that in the beginning as well :)
Or, you could just use the **pool** method that Celluloid offers. So, to parallelize the class above, we can do this:

{% highlight ruby %}
pool = YourClass.pool(size: 4) # number of threads in the pool

4.times {|e| pool.future.do_something(e) }

sleep 60 # or whatever time you estimate you need to complete all jobs
{% endhighlight %}

And the code above will queue 4 jobs to your pool. Notice that we're sleeping for
some seconds, to allow the jobs to finish. Yeah, I'm not a big fan of this approach either,
because your jobs may take much longer, or much less than 60 seconds.

Luckily, there's another class from Celluloid that comes to our rescue. Enter Celluloid::Condition:

{% highlight ruby %}
pool = YourClass.pool(size: 4)

number_of_jobs = 100 # or however many you may have

condition = Celluloid::Condition.new

# do work
# do more work
# when you are done, call condition.signal

condition.wait # waits until someone calls signal on it
{% endhighlight %}

I know this was quite a whirlwind introduction to Celluloid. I invite you to experiment with it for a bit.
I really like the power it offers you, and I'm quite sure you will find it useful too. As I learn more about it, I'll try to post more updates/tips/tricks.
