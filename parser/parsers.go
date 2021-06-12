package parser

import (
	"github.com/manishmeganathan/tuna/lexer"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

// A method of Parser that parses the lexer input
// into an abstract syntax tree Program
func (p *Parser) ParseProgram() *syntaxtree.Program {
	// Create a syntax tree program
	program := &syntaxtree.Program{}
	// Initialize the syntax tree statements slice
	program.Statements = []syntaxtree.Statement{}

	// Iterate until the lexer returns an EOF token
	for !p.isCursorToken(lexer.EOF) {
		// Parse the current token into a statement
		stmt := p.parseStatement()
		// Check if the statement exists
		if stmt != nil {
			// Add the statement to the syntax tree program statements
			program.Statements = append(program.Statements, stmt)
		}

		// Advance the parse cursor
		p.NextToken()
	}

	// Return the parsed syntax tree program
	return program
}

// A method of Parser that parses the token in the
// parse cursor into an syntax tree statement node
func (p *Parser) parseStatement() syntaxtree.Statement {
	// Check the value of the token in the parse cursor
	switch p.cursorToken.Type {
	// Let Statement
	case lexer.LET:
		// Parse the statement into a 'let' statement
		return p.parseLetStatement()

	// Default Case (not a recognized statement)
	default:
		return nil
	}
}

// A method of Parser that parses the token in the parse
// cursor into a LET statement node for the syntax tree
func (p *Parser) parseLetStatement() *syntaxtree.LetStatement {
	// Create a LET statement node with the token
	stmt := &syntaxtree.LetStatement{Token: p.cursorToken}

	// Check the peek cursor for an identfier token and move to it
	if !p.expectPeek(lexer.IDENT) {
		// no identifier token detected i.e invalid let statement
		return nil
	}

	// Assign the statement identifier to the statement node
	stmt.Name = &syntaxtree.Identifier{Token: p.cursorToken, Value: p.cursorToken.Literal}

	// Check the peek cursor for assignment token and move to it
	if !p.expectPeek(lexer.ASSIGN) {
		// no assign token detcted i.e invalid let statement
		return nil
	}

	// Advance until semicolon in encountered (TODO: let statement value detection)
	for !p.isCursorToken(lexer.SEMICOLON) {
		// Advance the parse cursor
		p.NextToken()
	}

	// Return the parsed let statement
	return stmt
}
