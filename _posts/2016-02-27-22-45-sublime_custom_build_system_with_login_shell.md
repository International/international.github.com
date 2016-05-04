---
layout: inner
title: sublime custom build system with login shell
tags: ["sublime", "build"]
---
Needed to run a node-cli tool. Build that enabled me to do so:

{% highlight json %}
{
  "cmd": "/bin/zsh -c 'source ~/.zshrc; resume export resume.html; sensible-browser resume.html'",
  "shell": true
}

{% endhighlight %}
