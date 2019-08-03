package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/bradford-hamilton/monkey-lang/lexer"
	"github.com/bradford-hamilton/monkey-lang/token"
)

const prompt = "> "

// Start - starts REPL, passes input into lexer line by line
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}
