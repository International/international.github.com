---
layout: inner
title: terminal inside neovim
tags: ["vim","neovim","shell"]
---
The command <b>:terminal</b> will start a new command. For example:
<pre>
:term rails c
</pre>

would start <b>rails c</b> in a new buffer. To get out of terminal mode, and be
able to switch through splits, use the following combination: <b>ctrl-\ ctrl-n</b>.
That is ctrl-backspace, followed by a ctrl-n.
