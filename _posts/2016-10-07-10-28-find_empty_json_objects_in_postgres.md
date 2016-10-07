---
layout: inner
title: find empty json objects in postgres
tags: ["postgres","json"]
---
Say you have a content column of type json, which holds arrays, and you want to delete the empty ones:

<pre>
select * from events where content::text = '[]'::text;
</pre>
