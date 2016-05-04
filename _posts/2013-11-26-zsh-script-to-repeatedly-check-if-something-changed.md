---
layout: inner
title: "zsh script to repeatedly check if something changed"
description: ""
category: ""
tags: ["zsh"]
---
{% highlight bash %}
#!/bin/zsh
command_to_run=$1
expect_change_from=$2
notification_to_show=$3
default_sleep_time=$4

if [[ -z $command_to_run ]] || [[ -z $expect_change_from ]] || [[ -z $notification_to_show ]]; then
  echo "Command to run , expected_change or notification are blank"
  exit 1
fi

if [[ -z $default_sleep_time ]]; then
  default_sleep_time=30
fi

echo "Running $command_to_run ; expecting output different than $expect_change_from ; every $default_sleep_time"

while true; do
  res=$(eval ${command_to_run})
  if [[ "$res" == $expect_change_from ]]; then
    echo "Value didn't change"
    sleep $default_sleep_time
  else
    notify-send -u critical $notification_to_show
    break
  fi
done
{% endhighlight %}

I wanted to be notified when CDN Sumo finished provisioning, so I wrote the script above.
It can be ran the following way:

zsh q.sh "heroku config:get CDN_SUMO_URL -a my_app" "my_app.herokuapp.com" "CDN_Sumo finished provisioning"
