---
layout: inner
title: neovim view who changed expandtab
tags: ["neovim","vim"]
---
When pressing newline in a .go file, I got tabs instead of spaces (my default setting).
This helped me track where the setting was changed:
<pre>
:verb set expandtab?
</pre>
