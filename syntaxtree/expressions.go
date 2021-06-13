package syntaxtree

import (
	"bytes"

	"github.com/manishmeganathan/tuna/lexer"
)

// A structure that represents a prefix expression node on the syntax tree
type PrefixExpression struct {
	// Represents the prefix operator token
	Token lexer.Token

	// Represents the operator literal string
	Operator string

	// Represents the expression after the operator
	Right Expression
}

// A method of PrefixExpression to satisfy the Expression interface
func (pe *PrefixExpression) expressionNode() {}

// A method of PrefixExpression that returns its token literal value
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// A method of PrefixExpression that returns its string representation
func (pe *PrefixExpression) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Start expression with parenthesis
	out.WriteString("(")
	// Add the operator and expression string values
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	// End expression with parenthesis
	out.WriteString(")")

	// Return the string of the buffer
	return out.String()
}

// A structure that represents an infix expression node on the syntax tree
type InfixExpression struct {
	// Represents the infix operator token
	Token lexer.Token

	// Represents the expression before the operator
	Left Expression

	// Represents the operator literal string
	Operator string

	// Represents the expression after the operator
	Right Expression
}

// A method of InfixExpression to satisfy the Expression interface
func (pe *InfixExpression) expressionNode() {}

// A method of InfixExpression that returns its token literal value
func (pe *InfixExpression) TokenLiteral() string { return pe.Token.Literal }

// A method of InfixExpression that returns its string representation
func (ie *InfixExpression) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Start expression with parenthesis
	out.WriteString("(")
	// Add left expression, the operator and the right expression values
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	// End expression with parenthesis
	out.WriteString(")")

	// Return the string of the buffer
	return out.String()
}
