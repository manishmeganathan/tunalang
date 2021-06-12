package syntaxtree

import (
	"github.com/manishmeganathan/tuna/lexer"
)

// A structure that represents a Let statement token
type LetStatement struct {
	// Represents the lexological token 'LET'
	Token lexer.Token

	// Represents the identifier in the let statement
	Name *Identifier

	// Represents the value in the let statement
	Value Expression
}

// A method of LetStatement to satisfy the Statement interface
func (ls *LetStatement) statementNode() {}

// A method of LetStatement that returns its token literal value
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// A structure that represents a Return statement token
type ReturnStatement struct {
	// Represents the lexological token 'RETURN'
	Token lexer.Token

	// Represents the value in the return statement
	ReturnValue Expression
}

// A method of ReturnStatement to satisfy the Statement interface
func (rs *ReturnStatement) statementNode() {}

// A method of ReturnStatement that returns its token literal value
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
