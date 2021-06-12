package syntaxtree

import "bytes"

// An interface that represents a node
// on the Abstract Syntax Tree
type Node interface {
	TokenLiteral() string
	String() string
}

// An interface that represents a statement
// node on the Abstract Syntax Tree
type Statement interface {
	Node
	statementNode()
}

// An interface that represents an expression
// node on the Abstract Syntax Tree
type Expression interface {
	Node
	expressionNode()
}

// A structure that represents the collection
// of statements in the program
type Program struct {
	Statements []Statement
}

// A method of Program that returns the token
// literal of the first statement in the program
func (p *Program) TokenLiteral() string {
	// Check if there are any statements in the program
	if len(p.Statements) > 0 {
		// Return the token literal of the first statement
		return p.Statements[0].TokenLiteral()

	} else {
		// Return an empty string
		return ""
	}
}

// A method of Program that returns the string representation of the
// Program by accumulating the string values of each statement in the
// program into a bytes buffer and returning its string value
func (p *Program) String() string {
	// Declare a bytes buffer
	var out bytes.Buffer

	// Iterate over the program statements
	// and collect their string values
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	// Return the string value of the buffer
	return out.String()
}
