---
layout: inner
title: monitor files for changes and run command with nodemon
tags: ["node","javascript"]
---
Used nodemon to restart a simple go web server, when changing any of the files from a folder called web:
<pre>nodemon --watch web --exec "go run main.go"</pre>
