# Monkey Interpreter

An implementation of the **Monkey programming language** written in **Go**, based on the book **Writing an Interpreter in Go** by Thorsten Ball.

This project builds a **tree-walking interpreter** from scratch. The interpreter parses source code, builds an Abstract Syntax Tree (AST), and evaluates it by walking the tree.

## About the Monkey language

Monkey is a small programming language designed specifically for educational purposes. It has no other implementation besides the one built in this project.

Although compact, Monkey includes many powerful language **features**:
1. C-like syntax  
2. Variable bindings  
3. Integers and booleans  
4. Arithmetic expressions  
5. Built-in functions  
6. First-class and higher-order functions  
7. Closures  
8. Strings  
9. Arrays  
10. Hash maps  

## Running the interpreter

```bash
make run
```

## Running the linter

To check code quality and style, run:

```bash
make linter
```

This uses [golangci-lint](https://golangci-lint.run/) with a 5-minute timeout.

## Running the tests

To run all tests with coverage:

```bash
make test
```

To run only the scanner package tests:

```bash
make test/scanner
```
