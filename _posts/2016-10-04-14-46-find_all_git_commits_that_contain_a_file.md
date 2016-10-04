---
layout: inner
title: find all git commits that contain a file
tags: ["git"]
---
If you want to find commits that included a now deleted file, you can run:
<pre>
git log --follow -- path/to/file
</pre>
