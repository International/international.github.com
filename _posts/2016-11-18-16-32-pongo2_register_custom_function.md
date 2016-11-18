---
layout: inner
title: pongo2 register custom function
tags: ["go","pongo2", "template"]
---
Say you want to create a function that returns the value of an environment variable:

{% highlight go %}
func get_env(name *pongo2.Value) *pongo2.Value {
	return pongo2.AsValue(os.Getenv(name.String()))
}
{% endhighlight %}

To use it in the templates, register it on the pongo2.Context:
{% highlight go %}
context := pongo2.Context{}
context["get_env"] = get_env
{% endhighlight %}

Last step, is to pass that context when you render the view.
