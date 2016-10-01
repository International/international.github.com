---
layout: inner
title: brew readline upgrade issue
tags: ["osx","mac","homebrew","readline","ruby"]
---
Updated readline, and after that got this error:
<pre>
    Sorry, you can't use byebug without Readline. To solve this, you need to
    rebuild Ruby with Readline support. If using Ubuntu, try `sudo apt-get
    install libreadline-dev` and then reinstall your Ruby.
/Users/geo/.rvm/gems/ruby-2.3.0@timers/gems/activesupport-5.0.0/lib/active_support/dependencies.rb:293:in `require': dlopen(/Users/geo/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0/x86_64-darwin15/readline.bundle, 9): Library not loaded: /usr/local/opt/readline/lib/libreadline.6.dylib (LoadError)
  Referenced from: /Users/geo/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0/x86_64-darwin15/readline.bundle
  Reason: image not found - /Users/geo/.rvm/rubies/ruby-2.3.0/lib/ruby/2.3.0/x86_64-darwin15/readline.bundle
</pre>

Turns out it was caused by upgrading to readline 7.0. Fixed by:
<pre>
brew uninstall readline # this left 6.3 still in the system
brew link readline --force
</pre>
