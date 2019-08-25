# Monkey Lang

Currently extending the Monkey programming language designed in [_Writing An Interpreter In Go_](https://interpreterbook.com/) and [_Writing a Compiler in Go_](https://compilerbook.com).

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