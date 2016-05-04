---
layout: inner
title: triggering remote builds from jenkins
tags: ["jenkins"]
---
If you receive this error when triggering builds from your scripts:

<pre>
Authentication required

<-- You are authenticated as: anonymous
Groups that you are in:

Permission you need to have (but didn't): hudson.model.Hudson.Read
... which is implied by: hudson.security.Permission.GenericRead
... which is implied by: hudson.model.Hudson.Administer
->
</pre>


Then install [this plugin](https://wiki.jenkins-ci.org/display/JENKINS/Build+Token+Root+Plugin), and then call the following url:

{% highlight bash %}
curl http://JENKINS_URL/buildByToken/build?job=JOB_NAME&token=TOKEN_NAME
{% endhighlight %}
