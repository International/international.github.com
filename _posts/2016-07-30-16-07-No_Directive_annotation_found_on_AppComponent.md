---
layout: inner
title: No Directive annotation found on AppComponent
tags: ["angular","javascript"]
---
Received the following exception when following the "Tour of heroes" app quickstart:

<pre>
No Directive annotation found on AppComponent
</pre>

Turned out, my code was like this:
{% highlight javascript %}

@Component({
    selector: 'my-app',
    template: '<h1>{{title}}</h1><h2>{{hero.name}} details!</h2>'
})

export class Hero {
  id: number;
  name: string;
}

export class AppComponent {

{% endhighlight %}

when it should have been like this:
{% highlight javascript %}


export class Hero {
  id: number;
  name: string;
}

@Component({
    selector: 'my-app',
    template: '<h1>{{title}}</h1><h2>{{hero.name}} details!</h2>'
})
export class AppComponent {
{% endhighlight %}

notice the annotation has to be on the AppComponent class.
