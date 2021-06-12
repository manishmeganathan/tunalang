package syntaxtree

// An interface that represents a node
// on the Abstract Syntax Tree
type Node interface {
	TokenLiteral() string
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
