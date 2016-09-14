---
layout: inner
title: filtering non-null values from json with jq
tags: ["jq","json"]
---
Given a json file, in the format:
<pre>
{
  "content": [
    {"name":some_name},
    {"name": null}
  ]
}
</pre>

The following command would extract only the entries where the name is not null:

<pre>
cat file.json | jq '.content | .[] | select(.name != null) | {name: .name}'
</pre>
