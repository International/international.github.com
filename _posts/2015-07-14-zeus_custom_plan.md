---
layout: inner
title: zeus custom plan
tags: ["ruby","zeus","rails"]
---
I wanted that zeus would reload all my shared examples for each test. This required the following custom plan.

{% highlight ruby %}
# ./custom_plan.rb
require 'zeus/rails'

class CustomPlan < Zeus::Rails
  def test
    Dir[Rails.root.join('spec/support/**/*.rb')].each { |f| require f }
    Dir[Rails.root.join('spec/controllers/shared_specs/**/*.rb')].each { |f| puts "reloading #{f}";load f }
    super
  end
end

Zeus.plan = CustomPlan.new
{% endhighlight %}

 To get zeus to use it, add a zeus.json with the following content:

{% highlight json %}
{
  "command": "ruby -rubygems -r./path/to_your/custom_plan -eZeus.go",

  "plan": {
    "boot": {
      "default_bundle": {
        "development_environment": {
          "prerake": {"rake": []},
          "runner": ["r"],
          "console": ["c"],
          "server": ["s"],
          "generate": ["g"],
          "destroy": ["d"],
          "dbconsole": []
        },
        "test_environment": {
          "test_helper": {"test": ["test"]}
        }
      }
    }
  }
}
{% endhighlight %}
