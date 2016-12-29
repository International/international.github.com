---
layout: inner
title: xvfb run something with custom screen size
tags: []
---
Problem: ran selenium integration tests. They pass locally, but fail that some element is not visible.
Took a browser screenshot, turned out that the browser window was very small, and the element was not in view.

Solution: run xvfb-run with custom args: <pre>xvfb-run --server-args="-screen 0, 1920x1080x24" npm run test</pre>. Solution found [here](http://stackoverflow.com/questions/6356169/resizing-an-xvfb-display)
