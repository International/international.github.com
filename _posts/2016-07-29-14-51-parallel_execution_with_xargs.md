---
layout: inner
title: parallel execution with xargs
tags: ["xargs","parallel"]
---
Useful trick seen [in the vagrant-cassandra repo](https://github.com/bcantoni/vagrant-cassandra/blob/master/2.MultiNode/up-parallel.sh)
{% highlight bash %}
$ cat <<EOF | xargs -P3 -I"BOXNAME" sh -c "vagrant up --provision BOXNAME"
node0
node1
node2
EOF
{% endhighlight %}

This can also be used with bash functions, but first you have to export them:
{% highlight bash %}

export -f some_function

cat <<EOF | xargs -P3 -I"BOXNAME" sh -c "some_function BOXNAME"
node0
node1
node2
EOF
{% endhighlight %}
