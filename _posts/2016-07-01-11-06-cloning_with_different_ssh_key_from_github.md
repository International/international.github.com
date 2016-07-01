---
layout: inner
title: cloning with different ssh key from github
tags: ["git","github","ssh"]
---
I needed to do a git clone using a different key than my own. This was what I had to do to get it to work:
* in ~/.ssh/config , I setup the following:

<pre>
Host github-work
  HostName github.com
  IdentityFile /path/to/the/private_key
  User the_github_username
</pre>
* add the key to the ssh-agent:
<pre>
ssh-add /path/to/private/key
</pre>

* export GIT_SSH_COMMAND:
<pre>
export GIT_SSH_COMMAND="ssh -i /path/to/the/private_key"
</pre>
* the actual clone:
<pre>
git clone git@github-work:organization_name/repo.git
</pre>

[This StackOverflow answer](http://superuser.com/questions/232373/how-to-tell-git-which-private-key-to-use) helped me with this.
