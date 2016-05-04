---
layout: inner
title: view version of a file from a certain git branch
tags: ["git"]
---
For example to view how `app/models/user.rb` looks on branch `name_of_branch`:

{% highlight bash %}
git show name_of_branch:app/models/user.rb
{% endhighlight %}
