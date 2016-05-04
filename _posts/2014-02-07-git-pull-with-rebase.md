---
layout: inner
title: "git pull with rebase"
description: ""
category: ""
tags: ["git", "rebase"]
---
To configure git to do a pull with rebase everytime:

{% highlight bash %}
git config --global branch.autosetuprebase always
{% endhighlight %}

And to setup an existing branch in the same way:

{% highlight bash %}
git config branch.ACTUAL_BRANCH_NAME.rebase true
{% endhighlight %}
