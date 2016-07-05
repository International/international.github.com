---
layout: inner
title: bash process substitution while loop
tags: ["bash","shell","process","pipe"]
---
I recently stumbled upon a problem in a bash script. I was doing something like this:

{% highlight bash %}
cat $1 | while read -r line; do
  line_with_new_value=$line
  echo $line_with_new_value | grep -vE '^#|^$' | grep -o -E '\${.*?}' | while read -r dollar_var; do
	  # do something with line_with_new_value here
		#
  done
	echo $line_with_new_value
done
{% endhighlight %}

Basically, when I was echoing, after the while, I was expecting the modified value to be shown.
However, piping into something will create a new subshell, and thus the modification is only
visible inside that shell. [This link](http://tldp.org/LDP/abs/html/process-sub.html) gives some useful examples
on process substitution. The fix was to rewrite that loop like this:

{% highlight bash %}
cat $1 | while read -r line; do
  line_with_new_value=$line
  while read -r interpolation; do
	  # modify line_with_new_value here
  done < <(echo $interpolated_line | grep -vE '^#|^$' | grep -o -E '\${.*?}')
  echo $line_with_new_value
done
{% endhighlight %}
