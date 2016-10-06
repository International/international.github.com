---
layout: inner
title: ember check if model settled before displaying
tags: ["javascript","ember"]
---
If you have a model with relationships, the `isSettled` property could be of some use, to avoid drawing only partial state, in case of ajax,
or even screen flickering:
<pre>
{{#if someModel.isSettled}}
  Display some data here
{{else}}
  Loading ...
{{/if}}
</pre>
