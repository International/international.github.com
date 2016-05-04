---
layout: inner
title: rails listen/notify with postgres
tags: ["rails","pg","activerecord"]
---

Example of using listen/notify functionality, taken from [here](http://stackoverflow.com/questions/16405520/postgres-listen-notify-rails)  

{% highlight ruby %}
# Be sure to check out a connection, so we stay thread-safe.
ActiveRecord::Base.connection_pool.with_connection do |connection|
  # connection is the ActiveRecord::ConnectionAdapters::PostgreSQLAdapter object
  conn = connection.instance_variable_get(:@connection)
  # conn is the underlying PG::Connection object, and exposes #wait_for_notify

  begin
    conn.async_exec "LISTEN channel1"
    conn.async_exec "LISTEN channel2"

    # This will block until a NOTIFY is issued on one of these two channels.
    conn.wait_for_notify do |channel, pid, payload|
      puts "Received a NOTIFY on channel #{channel}"
      puts "from PG backend #{pid}"
      puts "saying #{payload}"
    end

    # Note that you'll need to call wait_for_notify again if you want to pick
    # up further notifications. This time, bail out if we don't get a
    # notification within half a second.
    conn.wait_for_notify(0.5) do |channel, pid, payload|
      puts "Received a second NOTIFY on channel #{channel}"
      puts "from PG backend #{pid}"
      puts "saying #{payload}"
    end
  ensure
    # Don't want the connection to still be listening once we return
    # it to the pool - could result in weird behavior for the next
    # thread to check it out.
    conn.async_exec "UNLISTEN *"
  end
end
{% endhighlight %}

Under ruby 1.9, async_exec is just an alias for exec. Example notify call:

{% highlight ruby %}
ActiveRecord::Base.connection.execute "notify channel1,'some payload'"
{% endhighlight %}
