---
layout: inner
title: mysql access denied for user root
tags: ["mysql"]
---
Faced this error:
<b>Access denied for user 'root'@'localhost' (using password: YES) (Mysql2::Error)</b>

Found the fix [here](http://superuser.com/a/603027/11904):
<pre>
$ mysql -u root mysql
$mysql> UPDATE user SET Password=PASSWORD('my_password') where USER='root';
$mysql> FLUSH PRIVILEGES;
</pre>

but the 2nd variant was the one that worked in my case:
<pre>
update user set authentication_string=password('my_password') where user='root';
</pre>
