---
layout: inner
title: decoding encoded amazon authorization messages
tags: ["aws","amazon"]
---
If you receive something like this as a result of an aws operation:
<pre>
You are not authorized to perform this operation. Encoded authorization failure message: (redacted)
</pre>

Then, you can decode it like this:
<pre>
aws sts decode-authorization-message --encoded-message redacted
</pre>

To have the above command work, make sure your user has the <b>sts:DecodeAuthorizationMessage</b> grant.
