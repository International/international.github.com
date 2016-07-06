---
layout: inner
title: git error cannot lock ref
tags: ["git"]
---
Had this error:
<pre>
error: cannot lock ref 'refs/remotes/origin/feature/SomeBranch': ref refs/remotes/origin/feature/SomeBranch is at SOME_SHA b
ut expected SOME_OTHER_SHA
</pre>

This fixed it:
<pre>
rm .git/refs/remotes/origin/feature/SomeBranch
</pre>
