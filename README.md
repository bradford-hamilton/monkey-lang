# Monkey Lang

Currently extending original tutorial/src code:
- https://interpreterbook.com
- https://compilerbook.com

## Running monkey
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