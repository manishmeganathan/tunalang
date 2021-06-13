package syntaxtree

import (
	"bytes"
	"strings"

	"github.com/manishmeganathan/tuna/lexer"
)

// A structure that represents an Identifier literal
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

// A method of Identifier that returns its string representation
func (i *Identifier) String() string { return i.Value }

// A structure that represents an Integer literal
type IntegerLiteral struct {
	// Represents the lexological token 'INT'
	Token lexer.Token

	// Represents the integer value
	Value int64
}

// A method of IntegerLiteral to satisfy the Expression interface
func (il *IntegerLiteral) expressionNode() {}

// A method of IntegerLiteral that returns its token literal value
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// A method of IntegerLiteral that returns its string representation
func (il *IntegerLiteral) String() string { return il.Token.Literal }

// A structure that represents a Boolean literal
type BooleanLiteral struct {
	// Represents the lexological token 'TRUE'/'FALSE'
	Token lexer.Token

	// Represents the boolean value
	Value bool
}

// A method of BooleanLiteral to satisfy the Expression interface
func (b *BooleanLiteral) expressionNode() {}

// A method of BooleanLiteral that returns its token literal value
func (b *BooleanLiteral) TokenLiteral() string { return b.Token.Literal }

// A method of BooleanLiteral that returns its string representation
func (b *BooleanLiteral) String() string { return b.Token.Literal }

// A structure that represents a Function literal
type FunctionLiteral struct {
	// Represents the lexological token 'FN'
	Token lexer.Token

	// Represent the list of function parameters
	Parameters []*Identifier

	// Represents the the block of statements in the function
	Body *BlockStatement
}

// A method of FunctionLiteral to satisfy the Expression interface
func (b *FunctionLiteral) expressionNode() {}

// A method of FunctionLiteral that returns its token literal value
func (b *FunctionLiteral) TokenLiteral() string { return b.Token.Literal }

// A method of FunctionLiteral that returns its string representation
func (fl *FunctionLiteral) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer
	// Initialize the parameter list
	params := []string{}

	// Iterate over the parameters of the fn literal
	for _, p := range fl.Parameters {
		// Add parameter to the list
		params = append(params, p.String())
	}

	// Start function with the 'FN' token
	out.WriteString(fl.TokenLiteral())
	// Add the function parameters
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	// Add the function block of code
	out.WriteString(fl.Body.String())

	// Return the string from the buffer
	return out.String()
}
