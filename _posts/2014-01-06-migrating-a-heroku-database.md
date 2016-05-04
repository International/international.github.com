---
layout: inner
title: "migrating a heroku database"
description: ""
category: ""
tags: ["heroku","postgresql"]
---
Put app in maintenance:
{% highlight bash %}
heroku maintenance:on -a some-app
{% endhighlight %}

Create backup
{% highlight bash %}
heroku pgbackups:capture --expire -a some-app
{% endhighlight %}

Add new db addon, and check color in heroku config, and restore from backup:

{% highlight bash %}
heroku pgbackups:restore HEROKU_POSTGRESQL_COLOUR -a some-app
{% endhighlight %}

After everything's been done, promote it to be the main DB:

{% highlight bash %}
heroku pg:promote HEROKU_POSTGRESQL_COLOUR -a some-app
{% endhighlight %}

Remove maintenance, and restart app:

{% highlight bash %}
heroku maintenance:off -a some-app
heroku restart -a some-app
{% endhighlight %}
