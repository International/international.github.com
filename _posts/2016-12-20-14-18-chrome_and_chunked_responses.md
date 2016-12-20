---
layout: inner
title: chrome and chunked responses
tags: ["chrome","http","chunked","ajax"]
---
Had an issue when trying to react to a chunked response from a server. The <b>onprogress</b>
function was being called only at the end, and not for intermediate values. Found the answer [here](http://stackoverflow.com/questions/26164705/chrome-not-handling-chunked-responses-like-firefox-safari):

* turns out, you need to set the header <b>X-Content-Type-Options</b> to <b>nosniff</b>
