package evaluator

import (
	"github.com/manishmeganathan/tuna/object"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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

	// Prefix Expression Node
	case *syntaxtree.PrefixExpression:
		// Evaluate the expression into an object
		right := Evaluate(node.Right)
		// Evaluate the object for the operator
		return evalPrefixExpression(node.Operator, right)

	// IntegerLiteral Node
	case *syntaxtree.IntegerLiteral:
		// Return the Integer Object
		return &object.Integer{Value: node.Value}

	// BooleanLiteral Node
	case *syntaxtree.BooleanLiteral:
		// Return the native Boolean Object for the value
		return getNativeBoolean(node.Value)
	}

	// Return nil if not evaluated
	return NULL
}

// A function that returns the native Boolean
// Object for a given boolean value
func getNativeBoolean(input bool) *object.Boolean {
	// Check the input value
	if input {
		// Return the TRUE boolean native
		return TRUE
	}

	// Return the FALSE boolean native
	return FALSE
}
