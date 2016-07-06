---
layout: inner
title: bash overridable variables in scripts
tags: ["bash","shell"]
---
This script will show `default` when ran without any arguments:
{%highlight bash%}
myvar=${myvar:=default}
echo $myvar
{%endhighlight%}

and running it like this:
<pre>myvar=different ./script.sh</pre>

will show `different`
