# Monkey Interpreter
This project is a fully functional interpreter for the Monkey programming language, featuring its own lexer, AST-producing parser, and evaluator. The monkey language, as well as the interpreter's structure, is based on the book [Writing An Interpreter in Go](https://interpreterbook.com/). 

## Demo
https://github.com/user-attachments/assets/4909d89e-e96e-4222-8c1c-4069eb46f832

## Features
- Lexer, parser, AST, and tree-walking evaluator
- Data types: integers, strings, booleans, arrays, hashes
- First-class functions and closures
- Interactive REPL

## How It Works

### Lexer 
The lexer reads raw source text character by character and emits a flat stream of tokens (ex, `IDENT`, `INT`, `PLUS`, `Let`). It strips whitespace and flags illegal characters so errors surface early.

### Parser
The parser consumes the token stream and builds an **Abstract Syntax Tree (AST)** - a tree of nodes representing the program's structure. Expressions are parsed using a **Pratt parser** (top-down operator precedence), which handles precedence and associativity cleanly without a grammar table.

### Evaluator
The evaluator walks the AST recursively, visiting each node and producing an **Object** — the runtime representation of a value. Every value in Monkey implements an `Object` interface, with concrete types like `Integer`, `String`, `Boolean`, `Array`, `Hash`, `Function`, and `Error`.

### Environment & Closures
Variable bindings live in an **Environment** — a map from name to Object. Environments can enclose an outer one, which is how closures work: a function captures a reference to the environment it was defined in and can still access those bindings when called later.

## Installation

1. Clone the repository
```bash
git clone https://github.com/tablecrasher/monkey-interpreter.git
```

2. Run the interpreter
```bash
go run main.go
```
