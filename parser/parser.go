package parser

import (
	"fmt"

	"github.com/manishmeganathan/tuna/lexer"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

// A structure that represents a Parser
type Parser struct {
	// Represents the lexer used by the parser
	Lexer *lexer.Lexer

	// Represent the errors collected by the parser
	Errors []string

	// Represents the current token on the parser queue
	cursorToken lexer.Token

	// Represents the next token on the parser queue
	peekToken lexer.Token

	// Represnts the mapping of token type to its prefix parser function
	prefixParseFns map[lexer.TokenType]PrefixParseFn

	// Represnts the mapping of token type to its infix parser function
	infixParseFns map[lexer.TokenType]InfixParseFn
}

// Represents an alias for a prefix parser function
type PrefixParseFn func() syntaxtree.Expression

// Represents an alias for an infix parser function
type InfixParseFn func(syntaxtree.Expression) syntaxtree.Expression

// A constructor function that generates and returns a Parser after
// initializing it with the given lexer and advancing parser queue
// such that both the cursor and peek tokens are set
func NewParser(l *lexer.Lexer) *Parser {
	// Construct a parser with the lexer
	p := &Parser{Lexer: l, Errors: make([]string, 0)}

	// Initialize the prefix parser function map
	p.prefixParseFns = make(map[lexer.TokenType]PrefixParseFn)
	// Register the prefix parser functions
	p.registerPrefix(lexer.IDENT, p.parseIdentifier)
	p.registerPrefix(lexer.INT, p.parseIntegerLiteral)
	p.registerPrefix(lexer.BANG, p.parsePrefixExpression)
	p.registerPrefix(lexer.MINUS, p.parsePrefixExpression)
	p.registerPrefix(lexer.TRUE, p.parseBooleanLiteral)
	p.registerPrefix(lexer.FALSE, p.parseBooleanLiteral)
	p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(lexer.IF, p.parseIfExpression)

	// Initialize the infix parser function map
	p.infixParseFns = make(map[lexer.TokenType]InfixParseFn)
	// Register the infix parser functions
	p.registerInfix(lexer.PLUS, p.parseInfixExpression)
	p.registerInfix(lexer.MINUS, p.parseInfixExpression)
	p.registerInfix(lexer.SLASH, p.parseInfixExpression)
	p.registerInfix(lexer.ASTERISK, p.parseInfixExpression)
	p.registerInfix(lexer.EQ, p.parseInfixExpression)
	p.registerInfix(lexer.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(lexer.LT, p.parseInfixExpression)
	p.registerInfix(lexer.GT, p.parseInfixExpression)

	// Advance two tokens such that cursorToken
	// and peekToken are both set
	p.NextToken()
	p.NextToken()

	// Return the parser
	return p
}

// A method of Parser that advances the parser cursor and peek tokens
func (p *Parser) NextToken() {
	// Move cursor to peek position
	p.cursorToken = p.peekToken
	// Advance the peek to the next lexer token
	p.peekToken = p.Lexer.NextToken()
}

// A method of Parser that checks if the parse
// cursor is on the specified type of token
func (p *Parser) isCursorToken(t lexer.TokenType) bool {
	return p.cursorToken.Type == t
}

// A method of Parser that checks if the peek
// cursor is on the specified type of token
func (p *Parser) isPeekToken(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

// A method of Parser that check if the peek cursor
// on the specified type of token and advances to it
func (p *Parser) expectPeek(t lexer.TokenType) bool {
	// Check if peek token matches
	if p.isPeekToken(t) {
		// Advance the parse cursor
		p.NextToken()
		// Return true
		return true
	} else {
		// Add a peer error to the parser
		p.peekError(t)
		// Return false
		return false
	}
}

// A method of Parser that adds a Peek Error to the
// list of parse errors given the token type
func (p *Parser) peekError(t lexer.TokenType) {
	// Construct the error message
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	// Add the error message to the parser's errors
	p.Errors = append(p.Errors, msg)
}

// A method of Parser that adds a NoPrefixParseFn Error to
// the list of parse errors given the token type
func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	// Construct the error message
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	// Add the error message to the parser's errors
	p.Errors = append(p.Errors, msg)
}

// A method of Parser that registers a prefix parser
// given the token type and a prefix parser function
func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn PrefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// A method of Parser that registers a infix parser
// given the token type and a infix parser function
func (p *Parser) registerInfix(tokenType lexer.TokenType, fn InfixParseFn) {
	p.infixParseFns[tokenType] = fn
}
