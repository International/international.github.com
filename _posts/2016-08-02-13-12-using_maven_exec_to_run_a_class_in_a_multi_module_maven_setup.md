---
layout: inner
title: using maven exec to run a class in a multi module maven setup
tags: ["maven","java"]
---
Had the following structure:
<pre>
project/
  tests/
    pom.xml
    integration-tests/
      pom.xml
      src
        test
          java
            Potato.java
...
</pre>

and I wanted to run the Potato class using maven exec. The following command
did the trick:
<pre>
mvn -Dexec.mainClass="my.pkg.Potato" -Dexec.classpathScope=test exec:java -pl tests/integration-tests
</pre>


