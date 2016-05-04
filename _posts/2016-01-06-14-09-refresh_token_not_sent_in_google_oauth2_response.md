---
layout: inner
title: refresh_token not sent in google oauth2 response
tags: ["oauth2", "google", "api"]
---

I kept receiving tokens such as:

{% highlight javascript %}
{ access_token: 'REDACTED', token_type: 'Bearer', expiry_date: 1452085069587 }
{% endhighlight %}

when making calls to google api, even though I was doing the call like this:

{% highlight javascript %}
var url = cli.generateAuthUrl({
  access_type: 'offline',
  scope: 'https://www.googleapis.com/auth/calendar',
})
{% endhighlight %}

As it turns out, you only get the refresh_token on the first authorization, so
you should save it then. To address this, go to google apps, and revoke the app's access
to your API ( calendar in my case ), and go through the authorization flow again.

This time you'll receive a refresh_token as well.
