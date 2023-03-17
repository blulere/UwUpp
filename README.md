UwU++ Documentation (uwupp because it's a lot faster to type for me)

# uwupp, the next generation esoteric language
Created by 
[Deltaphish](https://github.com/Deltaphish/UwUpp),
documentation re-written by
[blulere](https://github.com/blulere)

Keep in mind I don't know how to use markdown

# Stay UwUthonic

Every object in uwupp *must* follow the naming conventions of UwUspeak in order to stay UwUthonic!

TODO: some other UwUthonic stuff lol

# Basics

```
UwU blah blah blahhhh
```

```
UwU
UwU This is a bunch of comments
UwU blah blah blahhhh
UwU
```
Write comments. Block comments are for the weak

`nuzzels("Hello World")` - Prints "Hello World" to the stdout.

`nuzzels(1234)` - Prints 1234 to the terminal. The type paradigm is dynamic by default but can be changed to static by making a type annotation, more on that later...

# Variables

Variables are dynamic and allocated with garbage collection by default.

All variables are deleted when a function reaches the end of its loop. This includes the global scope. If you would like to persist a variable, store it somewhere, I guess

`a iws 23` - Define the variable `a` and set it to the value 23. 

`b iws a` - Define the variable `b` and set it to the variable `a`, which is 23.
- This won't work if `a` doesn't exist.

# Arithmetic

## Addition
`c iws a pwus b` - Create a variable named `c`, which is equal to the result of `a` plus `b`. No, you don't get to use `+` like the rest of the languages. You made this commitment

## Subtraction
`c iws a minuws b` - Now `c` is equal to the result of `a` minus `b`.

## Multiplication

`c iws a twimes b` - `c` is `a` multiplied by `b` now.

## Division

`c iws a diwide b` - `c` is `a` divided by `b` now.

## Modulus

- As far as I know, the original uwupp documentation has no involvement of modulus. Please make an isswue/puww weqwest if the syntax is wrong in any way.

Not commonly known outside the medium-level programming scope, Modulus calculates the remainder given when a number is divided by another number. This can be useful when calculating overshoots, or other things I can't think of. But it's useful.

`c iws a modwuwo b` - `c` will become `a` modulo `b`.

Still confused? I've got you:

`nuzzels(12 modwuwo 2)` - Prints the result of 12 modulo 2. 12 divided by 2 is 6, and there is nothing left over, so the result is 0.

`nuzzels(10 modwuwo 7)` - Prints the result of 10 modulo 7. 10 divided by 7 is... a number, and if you were to take off the remainder of that, you'd get 3 left over. Bad example, once more...

`nuzzels(100 modwuwo 30)` - This does 100 modulo 30. 100 divided by 30 lets you put 30 into 100 3 times, and you get 10 left over. The 10 isn't 30, so that's the remainder, which just so happens to be the result; the result is 10.

Still confused? I don't got you, [but GeeksforGeeks does](https://www.geeksforgeeks.org/difference-between-modulo-and-modulus/).

# Arithmetic: Constants

In proper uwupp fashion, as a glow-up of the predecessor uwu-lang, hard-coded constants are a must-have to prevent enum madness or extra space for imports.

`twue` - True. Yes. 1. " ". []. Put this in an `OwO` loop to have some fun.

`fawse` - False. No. 0. I don't know enough JSFuck for this.

`nuww` - None, NULL, nil, nul. Kinda like false, but more dangerous.

# Loops n Conditions

Alike Go, (for the reason that I don't know how to implement other loops) uwupp doesn't have while loops, only for loops.
Here is a basic implementation of a while loop:
```
OwO *notices a gweatew than b*
    UwU do something?
    b iws b pwus 1
stawp
```
For as long as `a` is greater than `b`, the code inside will run. In this instance, `b` will add 1 to itself every iteration until `a` is no longer greater than `b`. Now, here's a for loop:
```
OwO *notices i iws 0, i gweatew than b, i iws pwus 1*
    UwU do something?
stawp
```
In this mammoth of a for loop, the variable `i` is assigned `0`. For as long as `i` is greater than `b`, the code inside the block will execute, and then `i` will increment by `1`.

```
OwO *notices a gweatew than b*
    UwU do something?
    nuzzels(a pwus b)
stawp
```
This is an if condition. If `a` is greater than `b`, the code will run. In this instance, the result of `a` plus `b` will be printed to the output.

# Functions

Functions give the functional programming paradigm to uwupp, which lets the programmer use the same code without having to copy and paste it. Functions also support parameters, allowing the programmer to function the function of functions by functioning the function functionally.

`nyaa *func(a, b)*` - Declare a function named `func`. It has the parameters `a` and `b`.

`func(12, 20)` - Use the declared function `func`, and give it the values 12 and 20.

`nuzzels(func(12, 20))` - Print the output of `func`.

`nyaa *new_pewson(name iws "Jown Doe", age iws 18)*` - Declare a function named `new_pewson`, and give it the default values of `name` equalling `"Jown Dow"`, and `age` equalling `18`.

`new_pewson("Jane Dow")` - Use `new_pewson` and override the default value `name` with `"Jane Dow"`, still keeping the default `age` value.

`new_pewson(name iws "Jane Dow", age iws nuw)` - Use `new_pewson` and override the default values with custom values completely. This works regardless of order, *however*, function parameters in non-order must be listed after function parameters in the right order.

# Classes

- Original implement of uwupp seems like it was designed for Haskell, a purely functional language. However, object oriented languages are cool. So, uwupp is object oriented now.

Coming soon :3

# To-Do:

* Move this stupidly long README to a readthedocs.io to make it easier to parse for data and stuff
* :3