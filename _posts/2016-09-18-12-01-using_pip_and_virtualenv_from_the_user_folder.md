---
layout: inner
title: using pip and virtualenv from the user folder
tags: ["python"]
---
Doing a:
<pre>
python get-pip.py --user
</pre>

installs pip just for the current user. To run it, I use:
<pre>
~/Library/Python/2.7/bin/pip install -U virtualenv
</pre>

The same goes for virtualenv:
<pre>
~/Library/Python/2.7/bin/virtualenv someenv
</pre>
