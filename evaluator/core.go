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
		return evalProgram(node)

	// Return Statement Node
	case *syntaxtree.ReturnStatement:
		// Evaluate the Expression in the return value
		val := Evaluate(node.ReturnValue)
		// Return the evaluated return object
		return &object.ReturnValue{Value: val}

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

	// Infix Expression Node
	case *syntaxtree.InfixExpression:
		// Evaluate the left node
		left := Evaluate(node.Left)
		// Evaluate the right node
		right := Evaluate(node.Right)
		// Evaluate the expression with the objects and the operator
		return evalInfixExpression(node.Operator, left, right)

	// Block Statement Node
	case *syntaxtree.BlockStatement:
		// Evaluate the statements in the block
		return evalBlockStatement(node)

	// If Expression Node
	case *syntaxtree.IfExpression:
		// Evaluate the if expression
		return evalIfExpression(node)

	// Integer Literal Node
	case *syntaxtree.IntegerLiteral:
		// Return the Integer Object
		return &object.Integer{Value: node.Value}

	// Boolean Literal Node
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

// A function that returns whether an Object is 'truthy'.
// An object is 'truthy' if it is not false and not null.
func isTruthy(obj object.Object) bool {
	// Check object value
	switch obj {

	// Null values are not truthy
	case NULL:
		return false

	// True values are truthy
	case TRUE:
		return true

	// False values are not truthy
	case FALSE:
		return false

	// All other types are truthy
	default:
		return true
	}
}
