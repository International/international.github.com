---
layout: inner
title: "iframe, cookies and IE"
description: ""
category: ""
tags: ["iframe","cookie","internet explorer"]
---
If you need to accept cookies from inside an iframe with IE, you need to set the P3P header:
P3P: CP="whatever P3P keywords"

or

P3P: CP="we do not support P3P"

And this will make IE happy.
