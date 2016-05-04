---
layout: inner
title: scala spark - No TypeTag available
tags: ["scala","spark"]
---
Received the error <b>No TypeTag available</b> when trying to package a spark
sql app. Turns out, it was caused by having a case class defined inside a method.
Moving it outside of the method fixes the compilation;
