---
layout: inner
title: simple sorting of csv lines via python and pandas
tags: ["python","pandas"]
---
Problem: have a file containing a lot of csv entries that I'd like to sort via a specific column ( published, in my case )

Solution: use pandas:

{% highlight python %}
import pandas as pd
df = pd.read_csv("input.csv")
df.sort_values("published",ascending=False).to_csv('input-sorted.csv', index = False)
{% endhighlight %}