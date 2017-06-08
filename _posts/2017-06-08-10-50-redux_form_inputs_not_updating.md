---
layout: inner
title: redux form inputs not updating
tags: ["javascript","react","redux"]
---
Problem: had some fields in a form, and they were not updating.

Solution: The name of the redux reducer was passed to combineReducers in a wrong way. Problem is also discussed [here](https://stackoverflow.com/a/40923003/31610)