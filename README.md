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

## Makefile Commands

To see all available commands, run:

```bash
make help
```

This will display a list of commands and their descriptions. Here are the available commands:

| Command         | Description                                 |
|-----------------|---------------------------------------------|
| help            | Show available commands                     |
| run             | Run the interpreter                         |
| run/file        | Run the interpreter with a specified file   |
| test            | Run tests with coverage                     |
| test/scanner    | Run tests in the scanner package            |
| linter          | Run linter (uses golangci-lint)             |
| fmt             | Format code                                 |
| vet             | Run go vet                                  |
| tidy            | Clean module dependencies                   |
| audit           | Run full quality checks (all of the above)  |

For more details on each command, use `make help`.
