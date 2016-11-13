---
layout: inner
title: check non-standard extensions with facebook flow
tags: ["node","javascript","flow"]
---
By default, <b>.js</b> and <b>.json</b> are analyzed. Had some files with <b>.es6</b>
extension. To get them to show up when running flow, I had to add this to <b>.flowconfig</b>:
<pre>
[options]
module.file_ext=.es6
module.file_ext=.js
module.file_ext=.json
</pre>
