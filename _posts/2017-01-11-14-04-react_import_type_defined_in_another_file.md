---
layout: inner
title: react import type defined in another file
tags: ["javascript", "react","flow"]
---
Problem: wanted to use a type defined in another file. Received following message:
<b>Named import from module '../mod.js' 'MyType' is a type, but not a value.In order to import it, please use 'import type'</b>

Solution:
{% highlight javascript %}
import type {MyType} from '../mod'
{% endhighlight %}
