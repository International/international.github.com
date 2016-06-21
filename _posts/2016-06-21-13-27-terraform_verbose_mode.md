---
layout: inner
title: terraform verbose mode
tags: ["terraform"]
---
If you're facing such an error:
<pre>
UnauthorizedOperation: You are not authorized to perform this operation.
</pre>

Run terraform like this:
<pre>
TF_LOG=1 terraform apply
</pre>

And this will most likely help you further with your debugging.
