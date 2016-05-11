---
layout: inner
title: cannot load such file -- rvm chef vagrant
tags: ["ruby","chef","vagrant","rvm"]
---
In which I discover that I should pass GEM_PATH to a runit controller service in order to have my gems loaded
<!-- exce -->

{% highlight ruby %}
runit_service 'my-service' do
  env(
    'HOME'     => "/home/vagrant",
    'GEM_PATH' => '/home/vagrant/.rvm/gems/ruby-2.3.0'
  )
  check true
end
{% endhighlight %}

Before I added the GEM_PATH change, I kept receiving this in my runit logs: <pre>home/vagrant/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0/rubygems/core_ext/kernel_require.rb:55:in `require': cannot load such file</pre>
