package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/manishmeganathan/tuna/lexer"
	"github.com/manishmeganathan/tuna/parser"
)

const PROMPT = ">> "
const TUNA = `
oooooooooooooooooooooooooooooooooo
  dP
  88
d8888P dP    dP 88d888b. .d8888b.
  88   88    88 88'  '88 88'  '88
  88   88.  .88 88    88 88.  .88
  '88P '88888P' db    db '8888888.
oooooooooooooooooooooooooooooooooo
`

const TUNA2 = `
  dP                              dP                            
  88                              88                            
d8888P dP    dP 88d888b. .d8888b. 88 .d8888b. 88d888b. .d8888b. 
  88   88    88 88'  '88 88'  '88 88 88'  '88 88'  '88 88'  '88 
  88   88.  .88 88    88 88.  .88 88 88.  .88 88    88 88.  .88 
  '88P '88888P' db    db '88888P8 db '88888P8 db    db '8888P88 
oooooooooooooooooooooooooooooooooooooooooooooooooooooooo~~~~.88
							d8888P 
`

// A function that starts the Tuna REPL
func StartREPL(in io.Reader, out io.Writer) {
	// Create a buffered IO scanner
	scanner := bufio.NewScanner(in)

	for {
		// Print the REPL line prompt
		fmt.Fprintf(out, PROMPT)
		// Scan the line
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Collect the scanned text
		line := scanner.Text()
		// Create a Lexer instance
		lex := lexer.NewLexer(line)
		par := parser.NewParser(lex)

		program := par.ParseProgram()
		if len(par.Errors) != 0 {
			printParserErrors(out, par.Errors)
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We has some trouble parsing!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
