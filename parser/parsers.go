package parser

import (
	"fmt"
	"strconv"

	"github.com/manishmeganathan/tuna/lexer"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
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

	// Return Statement
	case lexer.RETURN:
		// Parse the statement into a 'return' statement
		return p.parseReturnStatement()

	// Expression Statement
	default:
		// Parse the statement into an 'expression' statement
		return p.parseExpressionStatement()
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

// A method of Parser that parses the token in the parse
// cursor into a RETURN statement node for the syntax tree
func (p *Parser) parseReturnStatement() *syntaxtree.ReturnStatement {
	// Create a RETURN statement node with the token
	stmt := &syntaxtree.ReturnStatement{Token: p.cursorToken}
	// Advance the parse cursor
	p.NextToken()

	// Advance until semicolon in encountered (TODO: let statement value detection)
	for !p.isCursorToken(lexer.SEMICOLON) {
		p.NextToken()
	}

	// Return the parsed return statement
	return stmt
}

// A method of Parser that parses the token in the parse
// cursor into an expression statement node for the syntax tree
func (p *Parser) parseExpressionStatement() *syntaxtree.ExpressionStatement {
	// Create an expression statement node with the token
	stmt := &syntaxtree.ExpressionStatement{Token: p.cursorToken}
	// Parse the full expression
	stmt.Expression = p.parseExpression(LOWEST)

	// Check if the next token is a semicolon
	// (expressions do not have to end with a semicolon)
	if p.isPeekToken(lexer.SEMICOLON) {
		// Advance the parse cursor
		p.NextToken()
	}

	// Returned the parsed expression statement
	return stmt
}

// A method of Parser that parses a full expression given a precedence value
func (p *Parser) parseExpression(precedence int) syntaxtree.Expression {
	// Retrive the prefix parser function
	prefix := p.prefixParseFns[p.cursorToken.Type]

	// Check if the prefix parser is null
	if prefix == nil {
		p.noPrefixParseFnError(p.cursorToken.Type)
		// Return a nil
		return nil
	}

	// Call the prefix parser
	leftExp := prefix()
	// Return the left over expression
	return leftExp
}

// A method of Parser that parses an Identifier literal
func (p *Parser) parseIdentifier() syntaxtree.Expression {
	return &syntaxtree.Identifier{Token: p.cursorToken, Value: p.cursorToken.Literal}
}

// A method of Parser that parses an Integer literal
func (p *Parser) parseIntegerLiteral() syntaxtree.Expression {
	// Create an integer literal node with the token
	lit := &syntaxtree.IntegerLiteral{Token: p.cursorToken}

	// Parse the literal to int64
	value, err := strconv.ParseInt(p.cursorToken.Literal, 0, 64)
	// Check the error
	if err != nil {
		// Construct an error message
		msg := fmt.Sprintf("could not parse %q as integer", p.cursorToken.Literal)
		// Add the error to parser's errors
		p.Errors = append(p.Errors, msg)
		// Return a nil
		return nil
	}

	// Assign the integer literal node's value
	lit.Value = value
	// Return the integer literal node
	return lit
}

func (p *Parser) parsePrefixExpression() syntaxtree.Expression {
	// Create an prefix expression node with the token and operator literal
	expression := &syntaxtree.PrefixExpression{
		Token:    p.cursorToken,
		Operator: p.cursorToken.Literal,
	}

	// Advance the parse cursor
	p.NextToken()
	// Assign the prefix expression node's expression value
	expression.Right = p.parseExpression(PREFIX)
	// Return the prefix expression node
	return expression
}
