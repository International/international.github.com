---
layout: inner
title: redux-form conditional validation
tags: ["javascript","react","redux","redux-form"]
---
Problem: want to conditionally validate some fields in a form

Solution:

* Found [here](http://mycodesmells.com/post/conditional-validation-in-redux-form)

{% highlight javascript %}
// dispatch when you know which field you want to "set as required"
dispatch(change('formName', 'nameOfFieldRequired', required));

...
// and after, in your validation function:
validate: (values) => {
  const errors = {}
  const { nameOfFieldRequired, nameOfField } = values;
  if (nameOfFieldRequired && !nameOfField) {
    errors.nameOfField = 'Field is required!';
  }
  return errors;
}
{% endhighlight %}