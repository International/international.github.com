---
layout: inner
title: handling cookies with Mechanize
tags: ['ruby', 'mechanize', 'scraper']
---
{% highlight ruby %}
cookie_file = File.join(File.dirname(__FILE__), "cookies.yml")

w = Mechanize.new
if File.exists?(cookie_file)
  w.cookie_jar.load cookie_file
end

w.get("http://some_url")

# or whatever logic to verify whether you need to authenticate or not
if w.page.forms.count != 0
  form = w.page.forms.first

  email_field = form.field("admin_user[email]")
  pwd = form.field("admin_user[password]")

  email_field.value = "your_username@your_domain.com"
  pwd.value = "password"

  form.submit
end

w.cookie_jar.save(cookie_file, session: true)
{% endhighlight %}
