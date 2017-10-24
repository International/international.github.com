---
layout: inner
title: create-react-app environment variables
tags: ["node","react","environment"]
---
Problem: wanted to pass environment variables to an app created with create-react-app ( in development )

This is how I did it ( not exactly sure which version of react-scripts I'm using, since I ejected the config ):

* create a <b>.env.development</b> file, in your app's root folder and add your variables there:

{% highlight javascript %}
API_URL=http://api.local
{% endhighlight %}

* open <b>config/envs.js</b> and add your variable to the object in <b>getClientEnvironment</b> function:

{% highlight javascript %}
      {
        // Useful for determining whether we’re running in production mode.
        // Most importantly, it switches React into the correct mode.
        NODE_ENV: process.env.NODE_ENV || 'development',
        // Useful for resolving the correct path to static assets in `public`.
        // For example, <img src={process.env.PUBLIC_URL + '/img/logo.png'} />.
        // This should only be used as an escape hatch. Normally you would put
        // images into the `src` and `import` them in code to get their paths.
        PUBLIC_URL: publicUrl,
        API_URL: process.env.API_URL,
      }
{% endhighlight %}