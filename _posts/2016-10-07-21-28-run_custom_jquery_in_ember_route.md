---
layout: inner
title: run custom jquery in ember route
tags: ["javascript","ember"]
---
Needed to initialize some jquery elements in a route. [This](http://stackoverflow.com/a/33707561/31610) provided the answer:
{% highlight javascript %}
export default Ember.Route.extend({
  actions: {
    didTransition() {
      Ember.run.next(this, 'initTooltip');
    }
  },
  initTooltip() {
    Ember.$('[data-toggle="tooltip"]').tooltip();
  }
});
{% endhighlight %}
