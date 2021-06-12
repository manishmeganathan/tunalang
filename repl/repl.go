package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/manishmeganathan/tuna/lexer"
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

		// Iterate over the input characters and lex them into their tokens
		for tok := lex.NextToken(); tok.Type != lexer.EOF; tok = lex.NextToken() {
			// Print the tokens as they are lexed
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}