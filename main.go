package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bradford-hamilton/monkey-lang/repl"
)

func main() {
	engine := flag.String("engine", "vm", "Engine options are 'vm' or 'eval'")
	console := flag.Bool("console", false, "Repl options are true or false")
	flag.Parse()

	if *engine != "vm" && *engine != "eval" {
		fmt.Printf("Engine must be either 'vm' or 'eval'. Got %s\n", *engine)
		return
	}

	if *console == true {
		repl.Start(os.Stdin, os.Stdout, engine)
	} else {
		fmt.Print("Add ability to pass file path and execute...\n")
	}
}
