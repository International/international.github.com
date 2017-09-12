---
layout: inner
title: convert dot file to png
tags: ["graphviz","dot"]
---
Problem: have a dot file I'd like to view as a png.

Solution: found [here](https://stackoverflow.com/a/44157868/31610)

* install graphviz
<pre>
brew install graphviz
</pre>

* generate png:
<pre>
dot -Tpng DocName.dot -o DocName.png
</pre>