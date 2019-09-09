[![Build Status](https://travis-ci.org/bradford-hamilton/monkey-lang.svg?branch=master)](https://travis-ci.org/bradford-hamilton/monkey-lang)
[![Go Report Card](https://goreportcard.com/badge/github.com/bradford-hamilton/monkey-lang)](https://goreportcard.com/report/github.com/bradford-hamilton/monkey-lang)
[![codecov](https://codecov.io/gh/bradford-hamilton/monkey-lang/branch/master/graph/badge.svg)](https://codecov.io/gh/bradford-hamilton/monkey-lang)
[![GoDoc](https://godoc.org/github.com/bradford-hamilton/monkey-lang?status.svg)](https://godoc.org/github.com/bradford-hamilton/monkey-lang)
[![Go 1.12.9](https://img.shields.io/badge/go-1.12.9-9cf.svg)](https://golang.org/dl/)

# Monkey Lang

Currently extending the Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com/) and [_Writing a Compiler in Go_](https://compilerbook.com).

I will formally document the language and it's features at some point, but for now I'm keeping a list of the additional functionality I've added on top of original design:

1. Ability to execute Monkey files (.mo file ext) in addition to the interactive console. This is now the default behavior. Add `--console` flag when executing to drop into the REPL instead.
2. Both file execution and console usage respond to an `--engine=` flag where you can choose to use the evaluator or the VM.
3. Logical operators `&&` and `||`
4. Single line comments starting with `//`
5. Multi line comments using `/* */`
6. `const` variable declaration (although it only mocks let at this point until I add variable reassignment)
7. Modulo operator `%`
8. Postfix operators `++` and `--`
9. Comparison operators `>=` and `<=`
10. String comparisons using `!=` and `==`
11. Line numbers throughout the tokens/lexer/parsing/evaluator used for better errors.

## Installation
Install Monkey using `go get`:

```
go get -u github.com/bradford-hamilton/monkey-lang
```

## Usage
Build
```
go build -o monkey main.go
```

Run
```
./monkey [option...] filePath
```

## Examples

Running with vm
```
./monkey --engine=vm examples/program.mn
```

Running with evaluator
```
./monkey --engine=eval examples/program.mn
```

Run interactive console
```
./monkey --console
```