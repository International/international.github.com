---
layout: inner
title: golang blank struct when unmarshalling json
tags: ["go","golang","json"]
---
TIL: if you have a struct with unexported fields ( aka lowercase field names ),
and you try to unmarshal it from another package, the result will be a struct
with blank field names.
