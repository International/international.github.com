---
layout: inner
title: ignore commits by an author in git log
tags: ["git"]
---
This would show all commits where the author does not contain George:

<pre>git log --perl-regexp --author='^((?!George).*)$'</pre>
