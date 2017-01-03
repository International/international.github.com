---
layout: inner
title: golang include sha information in your binary
tags: ["go"]
---
Problem: you want to include the git SHA in your binary

Solution: Assuming the following code:

{% highlight go %}
package main

var appVersion string

func main() {
  fmt.Println("running version:", appVersion)
}
{% endhighlight %}

* you can run the binary like this: <b>go run -ldflags "-X main.appVersion=$(git rev-parse HEAD)" main.go</b>
* you can build it like this: <b> go build -ldflags "-X main.appVersion=$(git rev-parse HEAD)" main.go</b>
