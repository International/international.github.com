---
layout: inner
title: terminal inside neovim
tags: ["vim","neovim","shell"]
---
The command `:terminal` will start a new command. For example:
<pre>
:term rails c
</pre>

would start `rails c` in a new buffer. To get out of terminal mode, and be 
able to switch through splits, use the following combination: `ctrl-\ ctrl-n`.
That is ctrl-backspace, followed by a ctrl-n.
