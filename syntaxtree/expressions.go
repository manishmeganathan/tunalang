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

// A structure that represents an if expression node on the syntax tree
type IfExpression struct {
	// Represents the IF token
	Token lexer.Token

	// Represents the conditional expression
	Condition Expression

	// Represents the consequent statement if conditional evaulates to true
	Consequence *BlockStatement

	// Represents the alternative consequent statement if conditional evaulates to false
	Alternative *BlockStatement
}

// A method of IfExpression to satisfy the Expression interface
func (pe *IfExpression) expressionNode() {}

// A method of IfExpression that returns its token literal value
func (pe *IfExpression) TokenLiteral() string { return pe.Token.Literal }

// A method of IfExpression that returns its string representation
func (ie *IfExpression) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Start expression with if
	out.WriteString("if")
	// Add the condition and consequence
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	// Add the else and the alternate if it exists
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	// Return the string of the buffer
	return out.String()
}
