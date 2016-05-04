---
layout: inner
title: capybara login using warden helpers
tags: ["ruby","capybara","rspec","rails"]
---
Include Warden::Test::Helpers in your test, to be able to access the login_as method.
Example:

{% highlight ruby %}
describe "Something", :js => true do
  include Warden::Test::Helpers

  it "does smth" do
    user = FactoryGirl.create(:user)
    login_as(user) # method provided by warden
  end
end
{% endhighlight %}
