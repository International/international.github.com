---
layout: inner
title: sinatra routes inside a class
tags: ["ruby", "sinatra"]
---
Working example of sinatra routes with routes inside a class:

{% highlight ruby %}
module MyMod
  class Application < ::Sinatra::Base
    get "/hello" do
      "potato"
    end

    get "/movies/:movie_name/actors" do
      {
        actors: [
          {
            name: "Actor 1",
            character_played: "Character 1"
          },
          {
            name: "Actor 2",
            character_played: "Character 2"
          }
        ]
      }.to_json
    end
  end
end

MyMod::Application.run!
{% endhighlight %}
