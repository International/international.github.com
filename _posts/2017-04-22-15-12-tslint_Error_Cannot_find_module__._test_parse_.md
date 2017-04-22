---
layout: inner
title: tslint Error Cannot find module './test/parse'
tags: ["typescript","tslint","yarn"]
---
Problem: Received this error: <b>Error: Cannot find module './test/parse'</b>

Solution:

* it was caused by me having a <i>.yarnclean</i>
* remove it, delete <b>node_modules</b>, run <b>yarn</b> again