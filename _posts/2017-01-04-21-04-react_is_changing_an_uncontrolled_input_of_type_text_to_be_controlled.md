---
layout: inner
title: react is changing an uncontrolled input of type text to be controlled
tags: ["react","javascript"]
---
Problem: received following message in dev tools: <b>MyClass is changing an uncontrolled input of type text to be controlled. Input elements should not switch from uncontrolled to controlled (or vice versa). Decide between using a controlled or uncontrolled input element for the lifetime of the component</b>

Solution: make sure the property is specified in the initial state. Most likely your initial state does not include it. More info [here](http://stackoverflow.com/questions/37427508/react-changing-an-uncontrolled-input)
