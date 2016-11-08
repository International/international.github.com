---
layout: inner
title: YouCompleteMe and nvim
tags: ["nvim", "osx"]
---
Steps taken:

* install pip: <b>sudo easy_install pip</b>
* install neovim: <b>pip install --user neovim</b>
* added to init.vim the following: <b>NeoBundle 'Valloric/YouCompleteMe'</b>
* restart nvim, install YouCompleteMe
* in <b>~/.config/nvim/bundle/YouCompleteMe</b> run <b>./install.py --gocode-completer</b>
