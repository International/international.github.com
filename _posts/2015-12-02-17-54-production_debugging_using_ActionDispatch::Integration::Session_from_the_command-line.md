---
layout: inner
title: production debugging using ActionDispatch::Integration::Session from the command-line
tags: ["ruby","rails","debug"]
---

It may be useful to perform some requests to your controllers, from the Rails console. 
The following solution was tested against Rails 3.2.13, but I'm sure the tweaks should not be
that big, to get it working with other versions:

{% highlight ruby %}
session = ActionDispatch::Integration::Session.new(Rails.application)
{% endhighlight %}

after this, you can call get/post/put/etc on the session object:

{% highlight ruby %}
session.get "/", {}, {}
{% endhighlight %}

The first hash represents the params to be sent to the request, and the second one, additional headers.

Most likely, your controller will require authentication. To ensure that we're logged in, we go 
through the same steps as we would from a browser. 

{% highlight ruby %}
session.get "/login"
authenticity_token = session.response.body.match(/<[^<]+authenticity_token[^>]+value="([^"]+)"[^>]+>/)[1]
session.post("/login", {"user[email]" => "your@email","user[password]" => "your_password","authenticity_token" => authenticity_token})
{% endhighlight %}

So, we request login, and then we extract the authenticity token, and then we issue a post request to perform the login.
Of course, in your case, the name of the parameters might be different, so please adapt accordingly.

Once we post to the sessions controller, we can then issue requests to controllers that need authentication.

Hope you found this useful!
