---
layout: inner
title: fix vscode flow Type aliases can be used only in a .ts file
tags: ["vscode","flow"]
---
Problem: <b>Type aliases can be used only in a .ts file</b> appears when working with flowtype.

Solution:

* found [here](https://github.com/Microsoft/vscode/issues/15171)
* go to <b>Code -> Preferences -> Settings</b>
* in your workspace make this change:

{% highlight javascript %}
{
  "javascript.validate.enable": false
}
{% endhighlight %}