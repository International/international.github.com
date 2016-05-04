---
layout: inner
title: connection management with node.js
tags: ["node.js","javascript","coffee","coffeescript"]
---
Strange thing. The wireless network I'm currently on seems to randomly disconnect.
To prevent having to reconnect from the NetworkManager applet, I wrote the following
simple script:

{% highlight coffeescript %}
child_process = require("child_process")
network_name = "your-wireless-network-name"

reenqueue_time = 15000
network_check = ->
  console.log("started network check")
  child_process.exec "ping -c 1 google.ro", (err, stdout, stderr) ->
    if(err)
      console.log("pinging google failed")
      console.log(err)
      console.log("stderr:")
      console.log(stderr)
      child_process.exec "nmcli con down id #{network_name}", (err, stdout, stderr) ->
        if !err
          child_process.exec "nmcli con up id #{network_name}", (err, stdout, stderr) ->
            if err
              console.log("nmcli con up failed")
            else
              console.log("nmcli up succesful")
            setTimeout(network_check, reenqueue_time)
        else
          setTimeout(network_check, reenqueue_time)
    else
      console.log("pinging succesful")
      setTimeout(network_check, reenqueue_time)

network_check()
{% endhighlight %}
