---
layout: inner
title: node create readable stream from file contents
tags: ["nodejs", "javascript"]
---
Wanted to pipe the contents of a file through to <b>unzip</b> module. Achieved it like this:
{% highlight javascript %}
let unzip = require("unzip");
let fs = require("fs");

var Readable = require('stream').Readable

fs.readFile("x.zip", (err, data) => {
  if(err)
    throw err;

  var s = new Readable
  s.push(data)
  s.push(null)

  s.pipe(unzip.Parse()).on('entry', (entry) => {
      let fileName = entry.path;
      let type = entry.type; // 'Directory' or 'File'
      let size = entry.size;
      if(fileName.toLowerCase().endsWith(".srt")) {
        entry.pipe(fs.createWriteStream(fileName));
      } else {
        entry.autodrain()
      }
    });
});

{% endhighlight %}
