# Break down the io.uwu example to make it easier to understand

Original file:
```
UwU? is this IO?

nuzzels("What's your name?")
name iws wisten()
nuzzels("*notices* " pwus name pwus " OwO")
```

End result should be something like this:

```
keyword("nuzzels") lparen() string("What's your name?") rparen()
keyword("name") assignment() keyword("wisten") lparen() rparen()
keyword("nuzzels") lparen() string("*notices* ") addition() keyword("name") addition() string(" OwO") rparen()

```
Every token I need:

comment
iden
string
assignment
addition
lparen
rparen

This will go to the interpreter which will then interpret the file.