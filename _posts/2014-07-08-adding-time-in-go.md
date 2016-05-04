---
layout: inner
title: "adding time in go"
description: ""
category: ""
tags: ["go","date"]
---
{%highlight go%}
import "time"

now := time.Now()
end := now.Add(1 * time.Hour)
// formatting as iso8601
now_string := now.Format(time.RFC3339)
{%endhighlight%}
