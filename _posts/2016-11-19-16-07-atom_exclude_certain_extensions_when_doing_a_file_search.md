---
layout: inner
title: atom exclude certain extensions when doing a file search
tags: ["atom"]
---
You can specify more than one pattern in the file/directory pattern, with <b>!</b>
where you can do something like this:

<pre>
!vendor,!*.go
</pre>

This will have the effect of not searching in the <b>vendor</b> folder, or in any
file with <b>.go</b> extension.
