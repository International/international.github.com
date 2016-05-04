---
layout: inner
title: anonymous controllers with rspec
tags: ["ruby","rails","rspec"]
---
If you have new routes in your anonymous controller, the docs say you should do this:

{% highlight ruby %}
require "rails_helper"

RSpec.describe ApplicationController do
  controller do
    def custom
      render :text => "custom called"
    end
  end

  specify "a custom action can be requested if routes are drawn manually" do
    routes.draw { get "custom" => "anonymous#custom" }

    get :custom
    expect(response.body).to eq "custom called"
  end
end
{% endhighlight %}

However, in my case, the following change helped:

{% highlight ruby %}

specify "a custom action can be requested if routes are drawn manually" do
  routes.draw { get "custom", controller: 'application_controller' }

  get :custom
  expect(response.body).to eq "custom called"
end
{% endhighlight %}
