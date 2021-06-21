package evaluator

import (
	"github.com/manishmeganathan/tuna/object"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

// A function that evaluates a Syntax Tree given a node
// on it and returns an evaluated object
func Evaluate(node syntaxtree.Node) object.Object {
	// Check the type of Syntax Tree Node
	switch node := node.(type) {
	// Program Node (Tree Root)
	case *syntaxtree.Program:
		// Evaluate the statements in the program
		return evalStatements(node.Statements)

	// Expression Node
	case *syntaxtree.ExpressionStatement:
		// Recursive evaluation
		return Evaluate(node.Expression)

	// IntegerLiteral Node
	case *syntaxtree.IntegerLiteral:
		// Return the Integer Object
		return &object.Integer{Value: node.Value}
	}

	// Return nil if not evaluated
	return nil
}
