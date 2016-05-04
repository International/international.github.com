---
layout: inner
title: elixir get information about a module
tags: ["elixir"]
---
{% highlight elixir %}
iex(2)> String.module_info
[module: String,
 exports: [__info__: 1, downcase: 1, jaro_distance: 2, splitter: 2, strip: 2,
  to_float: 1, rjust: 2, ends_with?: 2, contains?: 2, upcase: 1, at: 2,
  equivalent?: 2, valid?: 1, reverse: 1, normalize: 2, splitter: 3,
  replace_trailing: 3, graphemes: 1, last: 1, strip: 1, chunk: 2, ljust: 2,
  ljust: 3, next_codepoint: 1, rstrip: 2, capitalize: 1, replace_suffix: 3,
  printable?: 1, first: 1, match?: 2, lstrip: 2, replace_leading: 3,
  split_at: 2, split: 2, length: 1, replace: 4, to_existing_atom: 1,
  next_grapheme_size: 1, duplicate: 2, starts_with?: 2, rstrip: 1, slice: 3,
  split: 1, valid_character?: 1, lstrip: 1, rjust: 3, slice: 2, to_atom: 1,
  ...], attributes: [vsn: [272593462254668022266866328062939769355]],
 compile: [options: [:debug_info], version: '6.0.2',
  time: {2016, 1, 14, 19, 20, 22},
  source: '/home/geo/code/elixir/lib/elixir/lib/string.ex'], native: false,
 md5: <<205, 19, 151, 86, 67, 198, 236, 82, 86, 152, 206, 221, 94, 34, 42, 11>>]

{% endhighlight %}
