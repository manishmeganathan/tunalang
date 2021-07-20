package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/manishmeganathan/tunalang/evaluator"
	"github.com/manishmeganathan/tunalang/lexer"
	"github.com/manishmeganathan/tunalang/object"
	"github.com/manishmeganathan/tunalang/parser"
)

const PROMPT = ">> "

// A function that starts the Tuna REPL
func StartREPL(in io.Reader, out io.Writer) {
	// Create a buffered IO scanner
	scanner := bufio.NewScanner(in)
	// Create a new execution environment
	env := object.NewEnvironment()

	for {
		// Print the REPL line prompt
		fmt.Fprint(out, PROMPT)
		// Scan the line
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		// Collect the scanned text
		line := scanner.Text()
		// Create a Lexer instance
		lex := lexer.NewLexer(line)
		// Create a Parser instance
		par := parser.NewParser(lex)

		// Parse the input into a Program
		program := par.ParseProgram()
		// Check for parser errors
		if len(par.Errors) != 0 {
			printParserErrors(out, par.Errors)
			continue
		}

		// Evaluate the Program
		evaluated := evaluator.Evaluate(program, env)
		// Print the evaluated values if they exist
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	// Print some error header
	io.WriteString(out, "Whoops! We had some trouble parsing!\n")
	io.WriteString(out, "parser errors:\n")

	// Iterate over the parser errors and print them out
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
