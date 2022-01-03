# The Mortar Programming Language

# Introduction
This is a reference manual for the Mortar programming language.

It supports mathematical expressions, variable bindings, functions, and the application of those functions, conditionals, return statements and even advanced concepts like higher-order functions and closures.
And then there are the different data types: integers, booleans, strings, arrays, and dictionaries.

This morter interpreter is written on top of [golang](https://go.dev/).

# Variables
A variable is a storage location for holding a value.

## Variable Declaration
A variable declaration creates one variable binds corresponding identifiers to them.
```
let <identifier> = <expression>;

let x = 10;
let y = 15;

let add = fn(a, b) { 
  return a + b;
};
```