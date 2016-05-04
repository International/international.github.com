---
layout: inner
title: automating the commandline using ruby
tags: ["ruby","expect","gem","greenletters"]
---
Lately, I've been thinking about how to automate some commandline stuff.
I knew about expect, so I started to dig into similar solutions for ruby.
Turns out, there's [expect](http://ruby-doc.org/stdlib-2.1.1/libdoc/pty/rdoc/IO.html), but ... 
it looks kind of complicated. And that's how the search for gems with a nicer API started.
I fire up github, and I write **ruby expect** in the searchbox. As expected, a ton of results,
but some of them look promising. Until I see that there've been no commits in quite some time.

So, I order them by **Recently Updated**. And I stumble upon [greenletters](https://github.com/avdi/greenletters).

Seeing that it's from [Avdi Grimm](http://about.avdi.org/), it kind of instantly gets a lot of credibility from me.

I read [the introductory blog post](http://devblog.avdi.org/2010/07/19/greenletters-painless-automation-and-testing-for-command-line-applications/), and I'm hooked.

So, I've decided to write a simple script, to automate running **heroku run bash**, because why not, right?

The first try failed really fast

{% highlight bash %}
home/geo/.rvm/gems/ruby-2.2.1@russh/gems/greenletters-0.3.0/lib/greenletters.rb:648:in `process_interruption': Interrupted (timeout) while waiting for output matching /~ $/. (Greenletters
::SystemError)
{% endhighlight %}

Turns out, greenletters seems to have a pretty small timeout by default. After a bit of browsing around in the issues section,
I stumble upon the **:timeout** option, which looks to be exactly what I need. After that, 
things just started to fall into place, so without further ado, I present to you the most useful
script, which runs **ls -l** and cats the **Gemfile**:

{% highlight ruby %}
require "bundler"

Bundler.require(:default)

adv = Greenletters::Process.new("heroku run bash -a myherokuapp", :timeout => 100, :transcript => $stdout)
adv.start!

adv.wait_for(:output, /run.\d+.*?\$/m)
adv << "ls -l\n"
adv.wait_for(:output, /\$/)
adv << "cat Gemfile\n"
adv.wait_for(:output, /\$/)
{% endhighlight %}

Low and behold, it works. The run pattern is in that way, because when heroku runs a command,
it's output has the following format:

{% highlight bash %}
Running `bash` attached to terminal... up, run.5710
{% endhighlight %}

Also, the heroku shell command prompt looks like this **~ $**. Hopefully, it makes sense why I used those regexes now.

I can see plenty of uses for such a gem, and it's API is pretty nice to work with.
Hopefully you can put it to great use. I'm interested to see what you'll come up with :)
