---
layout: inner
title: css selector excluding class
tags: ["jquery","css","selector"]
---
This would find you all the elements that are `.some.class` but not `.some_other_class`
<pre>
$(".some.class").not(".some_other_class")
</pre>
