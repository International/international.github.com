---
layout: inner
title: Could not load got :badfile
tags: ["elixir"]
---
Received an error like this in the console:
<pre>
Could not load module WebAssist.Repo, got: badfile
[error] beam/beam_load.c(1278): Error loading module 'Elixir.WebAssist.Repo':
  module name in object code is Elixir.Webassist.Repo
</pre>

Turns out, the module name was Webassist, and not WebAssist.
