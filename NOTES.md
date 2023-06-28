
- https://github.com/antlr/antlr4/blob/master/doc/getting-started.md

```
cd /usr/local/lib
sudo curl -O https://www.antlr.org/download/antlr-4.13.0-complete.jar
export CLASSPATH=".:/usr/local/lib/antlr-4.13.0-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "/usr/local/lib/antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'
```