package syntaxtree

import "github.com/manishmeganathan/tuna/lexer"

// A structure that represents an Identifier token
type Identifier struct {
	// Represents the lexological token 'IDENT'
	Token lexer.Token

	// Represents the identifier name
	Value string
}

// A method of Identifier to satisfy the Expression interface
func (i *Identifier) expressionNode() {}

// A method of Identifier that returns its token literal value
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
