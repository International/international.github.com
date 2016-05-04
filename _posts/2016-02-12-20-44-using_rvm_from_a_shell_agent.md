---
layout: inner
title: using rvm from a shell agent
tags: ["ruby","huginn","rvm"]
---
The following worked for me:

{% highlight json %}
{
  "path": "/home/geo/code/capy-extractor",
  "command": "BUNDLE_GEMFILE=/home/geo/code/capy-extractor/Gemfile /bin/bash -c '/home/geo/code/capy-extractor/wrapper.sh'",
  "suppress_on_failure": false,
  "suppress_on_empty_output": false,
  "expected_update_period_in_days": 1
}
{% endhighlight %}

and wrapper.sh:

{% highlight bash %}
#!/bin/bash -l
set -e

source $HOME/.bashrc
cd ~/code/capy-extractor
rvm use ruby-2.3.0@polty
ruby your-script.rb args

{% endhighlight %}
