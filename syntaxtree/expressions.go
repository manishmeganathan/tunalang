package syntaxtree

import (
	"bytes"
	"strings"

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
func (ie *InfixExpression) expressionNode() {}

// A method of InfixExpression that returns its token literal value
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

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
func (ie *IfExpression) expressionNode() {}

// A method of IfExpression that returns its token literal value
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

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

// A structure that represents an call expression node on the syntax tree
type CallExpression struct {
	// Represents the ( token
	Token lexer.Token

	// Represents the function identifier
	Function Expression

	// Represents the function arguments
	Arguments []Expression
}

// A method of CallExpression to satisfy the Expression interface
func (ce *CallExpression) expressionNode() {}

// A method of CallExpression that returns its token literal value
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }

// A method of CallExpression that returns its string representation
func (ce *CallExpression) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer
	// Initialize the args slice
	args := []string{}

	// Iterate over the expression arguments
	for _, a := range ce.Arguments {
		// Add them to the arg slice
		args = append(args, a.String())
	}
	// Add the function to the buffer
	out.WriteString(ce.Function.String())
	// Add the function arguments
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	// Return the string of the buffer
	return out.String()
}

// A structure that represents an index expression node on the syntax tree
type IndexExpression struct {
	// Represents the [ token
	Token lexer.Token

	// Represents the indexable expression
	Left Expression

	// Represents the index of the expression
	Index Expression
}

// A method of IndexExpression to satisfy the Expression interface
func (ie *IndexExpression) expressionNode() {}

// A method of IndexExpression that returns its token literal value
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }

// A method of IndexExpression that returns its string representation
func (ie *IndexExpression) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Start expression with parenthesis
	out.WriteString("(")
	// Add the left expression
	out.WriteString(ie.Left.String())
	// Add the index
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	// Return the string of the buffer
	return out.String()
}
