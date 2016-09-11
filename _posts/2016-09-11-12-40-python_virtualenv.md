---
layout: inner
title: python virtualenv
tags: ["python","virtualenv"]
---
Kind of similar to what rvm does:
<pre>
pip install virtualenv
cd my_project_folder
virtualenv nameofenv
source nameofenv/bin/activate
</pre>

Now you should see a prefix on the shell prompt, with the name of the environment. Afterwards, do:
<pre>
pip install somepackage
</pre>

and it will install it for a specific package. To get something akin to a `Gemfile.lock` do:
<pre>
pip freeze > requirements.txt
</pre>

And to do a `bundle install`:
<pre>
pip install -r requirements.txt
</pre>
