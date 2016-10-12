---
layout: inner
title: access google+ API from node.js
tags: ["node", "google", "google plus"]
---
Steps:
* enable Google+ API in google console
* create credentials for `Web Application`
* here's the code snippet I used to obtain the code:
{% highlight javascript %}

const google = require('googleapis');
const dotenv = require("dotenv").config();
const exec = require('child_process').exec;
const util = require("util");

const readline = require('readline');
const rl = readline.createInterface(process.stdin, process.stdout);

let client_id = process.env.GOOGLE_APP_CLIENT_ID,
  client_secret = process.env.GOOGLE_APP_CLIENT_SECRET;

let app_google_auth_url = process.env.GOOGLE_APP_AUTH_URL;
let client = new google.auth.OAuth2(client_id, client_secret, app_google_auth_url);

console.log(`client_id ${client_id} client_secret ${client_secret} oauth_url ${app_google_auth_url}`);

let scopes = [
  'https://www.googleapis.com/auth/plus.me'
];

let url = client.generateAuthUrl({
  access_type: 'offline',
  scope: scopes
});

console.log(`URL:${url}`);

rl.setPrompt('enter authorization code here> ');
rl.prompt();
rl.on('line', function(line) {
  client.getToken(line, function(err, tokens) {
    if(!err) {
      console.log(util.inspect(tokens));
    }
    process.exit(0);
  });
}).on('close',function(){
  process.exit(0);
});

{% endhighlight %}
* I'm using `dotenv` for node in this example, and I had the values for client_id, client_secret, and the redirect url mentioned there.
* at some point, the application will show the URL in the console, copy it and visit it in your browser
* google will redirect to your registered application URL with a `code` parameter. Write that down, paste it, and press enter. For this purpose, I had a rails app running in the background,
with a `binding.pry` so that I could inspect what's sent as `params`.
* you will then see in the console the access_token, refresh_token and expiry_date. Those you can use when you make the API calls.
