---
layout: inner
title: angular4 No provider for ControlContainer
tags: ["javascript","angular","angular4"]
---
Problem: Received error: <b>No provider for ControlContainer!</b>

Solution:

* forgot to include <b>FormsModule</b> besides <b>ReactiveFormsModule</b> in my app.module.ts :)