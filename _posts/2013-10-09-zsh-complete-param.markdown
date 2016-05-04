---
layout: inner
title: zsh complete param
tags: ["zsh", "autocompletion"]
---
Been struggling to have autocompletion based on the param. This is the function:

{% highlight bash %}
load_pg () {
    pg_restore --verbose --clean --no-acl --no-owner -h localhost -U $1 -d $2 $3
}
{% endhighlight %}

and here's the solution I found ( full code provided ) :

{% highlight bash %}

_load_pg() {
  _arguments "1: :->user" "2: :->db" "3: :->file"

  case $state in
    user)
      compadd $USER
    ;;
    db)
      compadd $(cat config/database.yml | grep -i database | awk '{print $2}')
    ;;  
    file)
      compadd $(ls *.dump*)
    ;;
  esac
}

load_pg() {
    pg_restore --verbose --clean --no-acl --no-owner -h localhost -U $1 -d $2 $3
}

compdef _load_pg load_pg
{% endhighlight %}
