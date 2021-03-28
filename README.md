# Bad Interpreter

## Getting started

```
At the current state, the interpreter is nothing more than a lexical tokenizer, it will read inputs from the source code (or in this case, console's REPL) and spit out the array of tokens it generates
```

```
Tokens are nothing more than [Token's Type, Token's Literal]
```

- Install Golang
- `cd` into the project directory
- `go run main.go`

## Reading the source code

You should probably read `token.go` first, then `lexer.go`, then `repl.go`, and lastly `main.go`. You can read my `parser.go` and `ast.go` however those have not been implemented yet.
