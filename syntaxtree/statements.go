package syntaxtree

import (
	"bytes"

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

// A method of LetStatment that returns its string representation
func (ls *LetStatement) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Add the token literal and identifier string into buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// Check if let statement has a value
	if ls.Value != nil {
		// Add the value into the buffer
		out.WriteString(ls.Value.String())
	}

	// Add a semicolon
	out.WriteString(";")
	// Return the string of the buffer
	return out.String()
}

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

// A method of ReturnStatement that returns its string representation
func (rs *ReturnStatement) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Add the token literal into the buffer
	out.WriteString(rs.TokenLiteral() + " ")
	// Check if the return statement has a value
	if rs.ReturnValue != nil {
		// Add the value to the buffer
		out.WriteString(rs.ReturnValue.String())
	}

	// Add a semicolon
	out.WriteString(";")
	// Return the string of the buffer
	return out.String()
}

// A structure that represents a statement wrapper for an expression
type ExpressionStatement struct {
	// Represents the first token of the expression
	Token lexer.Token

	// Represents the full Expression
	Expression Expression
}

// A method of ExpressionStatement to satisfy the Statement interface
func (es *ExpressionStatement) statementNode() {}

// A method of ExpressionStatement that returns its token literal value
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// A method of ExpressionStatement that returns its string representation
func (es *ExpressionStatement) String() string {
	// Check if the expression value is set
	if es.Expression != nil {
		// Return the expresion value
		return es.Expression.String()
	}
	// Return an empty string
	return ""
}
