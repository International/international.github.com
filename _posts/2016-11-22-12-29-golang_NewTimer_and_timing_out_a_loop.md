---
layout: inner
title: golang NewTimer and timing out a loop
tags: ["golang","timeout"]
---
{% highlight go %}
package main

import "time"
import "fmt"

func main() {

	timer := time.NewTimer(time.Second * 1)
LOOP:
	for {
		select {
		case <-timer.C:
			fmt.Println("timing out")
			break LOOP
		default:
			fmt.Println("not timing out")
		}

	}
}

{% endhighlight %}
