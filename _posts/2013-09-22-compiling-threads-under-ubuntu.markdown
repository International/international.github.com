---
layout: inner
title: Compiling threads under ubuntu
tags: ["c++", "compile"]
---

Command to compile:
**g++ r.cpp -o x -pthread -std=c++0x**

{% highlight cpp linenos=table %}
#include <iostream>
#include <thread>

using namespace std;
  
void do_work() {
  cout << "Doing work" << endl;
}

int main(int argc,char* argv[]) {
  std::thread t(do_work);
  t.join();
}
{% endhighlight %}
