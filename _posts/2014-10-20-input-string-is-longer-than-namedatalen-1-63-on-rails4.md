---
layout: inner
title: "Input string is longer than NAMEDATALEN 1 (63) on rails4"
description: ""
category: ""
tags: ["rails", "rails4"]
---
This error occurs because when you rename a table, the indexes are also renamed,
and most likely one of your indexes auto-generated name is too long.
