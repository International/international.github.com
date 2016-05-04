---
layout: inner
title: Compiling cpp code using a rakefile
tags: ["c++", "compile", "ruby", "rakefile", "rake"]
---

Been trying to build smth with libcurl. These is my Rakefile:

{% highlight ruby linenos %}
PROG  = "x"                                                                                                                                               
files = FileList["**/*.cpp"]                                                    
linked_libs_folders  = ["/home/john/code/cpp1/curl-7.32.0/installed/lib"].map {|e| "-L#{e}"}.join(" ")
libs_to_link_against = ["curl"].map {|e| "-l#{e}"}.join(" ")                    
file PROG => files do                                                           
  puts "Compiling #{PROG}"                                                      
  sh "g++ #{files} -o #{PROG} #{linked_libs_folders} #{libs_to_link_against}"   
end                                                                             

task :default => [PROG] 
{% endhighlight %}
