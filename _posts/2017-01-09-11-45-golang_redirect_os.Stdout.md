---
layout: inner
title: golang redirect os.Stdout
tags: ["go","golang","stdout"]
---
Problem: want to redirect <b>os.Stdout</b> and access it's contents from a string variable

Solution: [this answer](http://stackoverflow.com/a/10476304/31610) proved very useful:

{% highlight go %}
package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
)

func print() {
    fmt.Println("output")
}

func main() {
    old := os.Stdout // keep backup of the real stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    print()

    outC := make(chan string)
    // copy the output in a separate goroutine so printing can't block indefinitely
    go func() {
        var buf bytes.Buffer
        io.Copy(&buf, r)
        outC <- buf.String()
    }()

    // back to normal state
    w.Close()
    os.Stdout = old // restoring the real stdout
    out := <-outC

    // reading our temp stdout
    fmt.Println("previous output:")
    fmt.Print(out)
}
{% endhighlight %}
