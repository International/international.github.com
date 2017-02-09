---
layout: inner
title: webpack ignore plugin
tags: ["webpack","ignore","javascript", "optimization"]
---
The [ignore](https://webpack.github.io/docs/list-of-plugins.html#ignoreplugin) plugin was useful to
reduce the size of moment.js. Found [this answer](https://github.com/moment/moment/issues/2373#issuecomment-233293900),
which gave this useful snippet:

{% highlight javascript %}
new webpack.IgnorePlugin(/^\.\/locale$/, /moment$/) // Ignore all optional deps of moment.js
{% endhighlight %}

Couple that with the BundleAnalyzerPlugin ( config has to be done in webpack.config.js ) :

{% highlight javascript %}
new BundleAnalyzerPlugin({
  reportFilename: 'report.html',
  // Automatically open report in default browser
  openAnalyzer: false,
  // If `true`, Webpack Stats JSON file will be generated in bundles output directory
  generateStatsFile: false,
  // Name of Webpack Stats JSON file that will be generated if `generateStatsFile` is `true`.
  // Relative to bundles output directory.
  statsFilename: 'stats.json',
  // Options for `stats.toJson()` method.
  // For example you can exclude sources of your modules from stats file with `source: false` option.
  // See more options here: https://github.com/webpack/webpack/blob/webpack-1/lib/Stats.js#L21
  statsOptions: null,
  // Log level. Can be 'info', 'warn', 'error' or 'silent'.
  logLevel: 'info'
}),

{% endhighlight %}

and you have a really nice way of seeing what causes your js file to be so big.
