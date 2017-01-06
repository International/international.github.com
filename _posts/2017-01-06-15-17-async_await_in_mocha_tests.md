---
layout: inner
title: async await in mocha tests
tags: ["javascript", "async","await","mocha"]
---
Problem: use async-await in mocha tests. Had an async test, and received this error: <b>Error: Resolution method is overspecified. Specify a callback *or* return a Promise; not both.</b>

Solution: good example [here](https://github.com/mochajs/mocha/issues/2407#issuecomment-253577679):

{% highlight javascript %}
it("should work", done => {
  (async () => {
    await something;
    done();
  })();
});
{% endhighlight %}

Also, I had to require <b>babel-polyfill</b> in my test instruction:

<b>./node_modules/mocha/bin/mocha --compilers js:babel-core/register --require babel-polyfill test</b>
