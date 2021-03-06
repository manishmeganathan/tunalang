package parser

import (
	"fmt"
	"strconv"

	"github.com/manishmeganathan/tunalang/lexer"
	"github.com/manishmeganathan/tunalang/syntaxtree"
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
	INDEX       // myList[X]
)

var precedences = map[lexer.TokenType]int{
	lexer.EQ:       EQUALS,
	lexer.NOT_EQ:   EQUALS,
	lexer.LT:       LESSGREATER,
	lexer.GT:       LESSGREATER,
	lexer.PLUS:     SUM,
	lexer.MINUS:    SUM,
	lexer.SLASH:    PRODUCT,
	lexer.ASTERISK: PRODUCT,
	lexer.LPAREN:   CALL,
	lexer.LBRACK:   INDEX,
}

var traceON = false

// A function that returns the precedence value
// for the given lexilogical token type
func GetPrecedence(tokentype lexer.TokenType) int {
	// Check the precedence table for the token type
	if p, ok := precedences[tokentype]; ok {
		// Return the precedence value
		return p
	}
	// Return the lowest precedence if not found in the map
	return LOWEST
}

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
		// Add the statement to the syntax tree program statements
		program.Statements = append(program.Statements, stmt)
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

	// Advance the parse cursor
	p.NextToken()
	// Assign the parsed let value expression
	stmt.Value = p.parseExpression(LOWEST)

	// Advance until semicolon in encountered
	if p.isPeekToken(lexer.SEMICOLON) {
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

	// Assign the parsed return value
	stmt.ReturnValue = p.parseExpression(LOWEST)

	// Advance until semicolon in encountered
	if p.isPeekToken(lexer.SEMICOLON) {
		p.NextToken()
	}

	// Return the parsed return statement
	return stmt
}

// A method of Parser that parses the token in the parse
// cursor into an expression statement node for the syntax tree
func (p *Parser) parseExpressionStatement() *syntaxtree.ExpressionStatement {
	if traceON {
		// Print parser trace
		defer untrace(trace("parseExpressionStatement"))
	}

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
	if traceON {
		// Print parser trace
		defer untrace(trace("parseExpression"))
	}

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

	// Iterate on the token and check if the next token is neither a
	// semicolon and hash a precedence below the the given precedence
	for !p.isPeekToken(lexer.SEMICOLON) && precedence < GetPrecedence(p.peekToken.Type) {
		// Retrive the infix parser function
		infix := p.infixParseFns[p.peekToken.Type]
		// Check if the infix parser is null
		if infix == nil {
			// Return the left expression as is
			return leftExp
		}

		// Advance the parse cursor
		p.NextToken()
		// Call the infix parser on left expression and update it
		leftExp = infix(leftExp)
	}

	// Return the left expression
	return leftExp
}

// A method of Parser that parses an Identifier literal
func (p *Parser) parseIdentifier() syntaxtree.Expression {
	return &syntaxtree.Identifier{Token: p.cursorToken, Value: p.cursorToken.Literal}
}

// A method of Parser that parses an Integer literal
func (p *Parser) parseIntegerLiteral() syntaxtree.Expression {
	if traceON {
		// Print parser trace
		defer untrace(trace("parseIntegerLiteral"))
	}

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

// A method of Parser that parses a Prefix Expression
func (p *Parser) parsePrefixExpression() syntaxtree.Expression {
	if traceON {
		// Print parser trace
		defer untrace(trace("parsePrefixExpression"))
	}

	// Create a prefix expression node with the token and operator literal
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

// A method of Parser that parses an Infix Expression
func (p *Parser) parseInfixExpression(left syntaxtree.Expression) syntaxtree.Expression {
	if traceON {
		// Print parser trace
		defer untrace(trace("parseInfixExpression"))
	}

	// Create an infix expression node with the token, operator literal and left expression
	expression := &syntaxtree.InfixExpression{
		Token:    p.cursorToken,
		Operator: p.cursorToken.Literal,
		Left:     left,
	}

	// Determine the precedence of the cursor token
	precedence := GetPrecedence(p.cursorToken.Type)
	// Advance the parse cursor
	p.NextToken()
	// Assign the right expression to the parsed value of
	// the right expression with the given precedence
	expression.Right = p.parseExpression(precedence)
	// Return the infix expression node
	return expression
}

// A method of Parser that parses a Boolean Literal
func (p *Parser) parseBooleanLiteral() syntaxtree.Expression {
	return &syntaxtree.BooleanLiteral{Token: p.cursorToken, Value: p.isCursorToken(lexer.TRUE)}
}

// A method of Parser that parses a String Literal
func (p *Parser) parseStringLiteral() syntaxtree.Expression {
	return &syntaxtree.StringLiteral{Token: p.cursorToken, Value: p.cursorToken.Literal}
}

// A method of Parser that parses Grouped Expressions
func (p *Parser) parseGroupedExpression() syntaxtree.Expression {
	// Advance the parse cursor
	p.NextToken()
	// Parse the expression in the parentheses
	exp := p.parseExpression(LOWEST)

	// Check for closing parentheses
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Return the parsed parentheses
	return exp
}

// A method of Parser that parses Block Statements
func (p *Parser) parseBlockStatement() *syntaxtree.BlockStatement {
	// Create a block statement node for the syntax tree
	block := &syntaxtree.BlockStatement{Token: p.cursorToken}
	// Initialze the block statements slice
	block.Statements = []syntaxtree.Statement{}
	// Advance the parse cursor
	p.NextToken()

	// Iterate until an } or EOF token is encountered
	for !p.isCursorToken(lexer.RBRACE) && !p.isCursorToken(lexer.EOF) {
		// Parse the statement
		stmt := p.parseStatement()
		// Add it the to block statements
		block.Statements = append(block.Statements, stmt)
		// Advance the parse cursor
		p.NextToken()
	}

	// Return the parsed block statement
	return block
}

// A method of Parser that parses If Expressions
func (p *Parser) parseIfExpression() syntaxtree.Expression {
	// Create an if expression node for the syntax tree
	expression := &syntaxtree.IfExpression{Token: p.cursorToken}
	// Check for the conditional opening ( token
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	// Advance the parse cursor
	p.NextToken()
	// Parse the condition expression
	expression.Condition = p.parseExpression(LOWEST)

	// Check for the conditional ending ) token
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Check for the block opening { token
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	// Parse the consequence block statement
	expression.Consequence = p.parseBlockStatement()

	// Check the ELSE token
	if p.isPeekToken(lexer.ELSE) {
		// Advance the parse cursor
		p.NextToken()

		// Check for the block opening { token
		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}
		// Parse the alternate consequence block statement
		expression.Alternative = p.parseBlockStatement()
	}

	// Return the parsed if expression
	return expression
}

// A method of Parser that parses Function parameters
func (p *Parser) parseFunctionParameters() []*syntaxtree.Identifier {
	// Initialize a slice of identifier nodes
	identifiers := []*syntaxtree.Identifier{}

	// Check if the next token is ) token
	if p.isPeekToken(lexer.RPAREN) {
		// Advance the parse cursor
		p.NextToken()
		// Return the empty list of identifiers
		return identifiers
	}

	// Advance the parse cursor
	p.NextToken()

	// Create an identifer node for the syntax tree and add it to the list
	ident := &syntaxtree.Identifier{Token: p.cursorToken, Value: p.cursorToken.Literal}
	identifiers = append(identifiers, ident)

	// Iterate as long as the next token is comma
	for p.isPeekToken(lexer.COMMA) {
		// Advance the parse cursor twice (skip the over the comma)
		p.NextToken()
		p.NextToken()

		// Create an identifer node for the syntax tree and add it to the list
		ident := &syntaxtree.Identifier{Token: p.cursorToken, Value: p.cursorToken.Literal}
		identifiers = append(identifiers, ident)
	}

	// Check for the ) token
	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	// Return the list of idenfier nodes for the function parameters
	return identifiers
}

// A method of Parser that parses Function literals
func (p *Parser) parseFunctionLiteral() syntaxtree.Expression {
	// Create a function literal node
	lit := &syntaxtree.FunctionLiteral{Token: p.cursorToken}

	// Check for the function parameter begin ( token
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}
	// Assign the fn paramters after parsing them
	lit.Parameters = p.parseFunctionParameters()

	// Check for the function block begin { token
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	// Assign the fn body after parsing it
	lit.Body = p.parseBlockStatement()

	// Return the parsed function literal node
	return lit
}

// A method of Parser that parses a list of expressions
func (p *Parser) parseExpressionList(end lexer.TokenType) []syntaxtree.Expression {
	// Initialize a slice of expression nodes
	list := []syntaxtree.Expression{}

	// Check if the next token is the end token
	if p.isPeekToken(end) {
		// Advance the parse cursor
		p.NextToken()
		// Return the empty list of expressions
		return list
	}

	// Advance the parse cursor
	p.NextToken()
	// Append the parsed expression to the list
	list = append(list, p.parseExpression(LOWEST))

	// Iterate as long as the next token is not the end token
	for p.isPeekToken(lexer.COMMA) {
		// Advance the parse cursor twice (skip over the comma)
		p.NextToken()
		p.NextToken()

		// Append the parsed expression to the list
		list = append(list, p.parseExpression(LOWEST))
	}

	// Check if the next token is the end token
	if !p.expectPeek(end) {
		return nil
	}

	// Return the parsed list of expressions
	return list
}

// A method of Parser that parses Call expressions
func (p *Parser) parseCallExpression(function syntaxtree.Expression) syntaxtree.Expression {
	// Create a call expression node for the syntax tree
	exp := &syntaxtree.CallExpression{Token: p.cursorToken, Function: function}
	// Parse the call arguments
	exp.Arguments = p.parseExpressionList(lexer.RPAREN)

	// Return the parsed call expression
	return exp
}

// A method of Parser that parses List literals
func (p *Parser) parseListLiteral() syntaxtree.Expression {
	// Create a list literal node for the syntax tree
	list := &syntaxtree.ListLiteral{Token: p.cursorToken}
	// Parse the expression for the list elements
	list.Elements = p.parseExpressionList(lexer.RBRACK)

	// Return the parsed list literal
	return list
}

// A method of Parser that parses Index expressions
func (p *Parser) parseIndexExpression(left syntaxtree.Expression) syntaxtree.Expression {
	// Create an index expression node for the syntax tree
	exp := &syntaxtree.IndexExpression{Token: p.cursorToken, Left: left}

	// Advance the parse cursor
	p.NextToken()
	// Parse the expression for the index
	exp.Index = p.parseExpression(LOWEST)
	// Check for the ] token
	if !p.expectPeek(lexer.RBRACK) {
		return nil
	}

	// Return the parsed index expression
	return exp
}

// A method of Parser that parses Map literals
func (p *Parser) parseMapLiteral() syntaxtree.Expression {
	// Create a map literal node for the syntax tree
	hash := &syntaxtree.MapLiteral{Token: p.cursorToken}
	// Initialize a map for key value pairs
	hash.Pairs = make(map[syntaxtree.Expression]syntaxtree.Expression)

	// Iterate until the next token is the } token
	for !p.isPeekToken(lexer.RBRACE) {
		// Advance the parse cursor
		p.NextToken()
		// Parse the key expression
		key := p.parseExpression(LOWEST)
		// Assert the colon token expected
		if !p.expectPeek(lexer.COLON) {
			return nil
		}

		// Advance the parse cursor
		p.NextToken()
		// Parse the value expression
		value := p.parseExpression(LOWEST)
		// Assign the key value pair
		hash.Pairs[key] = value

		// Check if the map either continues or ends
		if !p.isPeekToken(lexer.RBRACE) && !p.expectPeek(lexer.COMMA) {
			return nil
		}
	}

	// Check if the next token is the } token
	if !p.expectPeek(lexer.RBRACE) {
		return nil
	}

	// Return the parsed map literal
	return hash
}
