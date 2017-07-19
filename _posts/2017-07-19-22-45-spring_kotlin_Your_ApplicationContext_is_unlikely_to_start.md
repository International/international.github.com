---
layout: inner
title: spring kotlin  Your ApplicationContext is unlikely to start
tags: ["spring","kotlin"]
---
Problem: started a kotlin spring app. Received this message: <pre>Your ApplicationContext is unlikely to start due to a `@ComponentScan` of the default package</pre>. Problem is described [here](https://stackoverflow.com/q/41729712/31610) as well.

Solution:

* move files out of the default package, to a named one.