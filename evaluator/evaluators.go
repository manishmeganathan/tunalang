package evaluator

import (
	"github.com/manishmeganathan/tuna/object"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

// A function that evaluates a set of Syntax tree
// statements into an evaluated object
func evalStatements(stmts []syntaxtree.Statement) object.Object {
	// Declare an object
	var result object.Object

	// Iterate over the program statements
	for _, statement := range stmts {
		// Update the result object
		result = Evaluate(statement)
	}

	// Return the result object
	return result
}

// A function that evaluates a prefix expression
// given a prefix operator and an object
func evalPrefixExpression(operator string, right object.Object) object.Object {
	// Check the type of operator
	switch operator {

	// Bang Operator
	case "!":
		// Evaluate the object for the bang operator
		return evalBangOperatorExpression(right)

	// Minus Operator
	case "-":
		// Evaluate the object for the minus operator
		return evalMinusPrefixOperatorExpression(right)

	// Unsupported Operator
	default:
		// Return null (for now)
		return NULL
	}
}

// A function that evaluates an infix expression given
// a infix operator and the left and right objects
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	// Check the object type combinations
	switch {

	// Both are Integers
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		// Evaluate expression for integer objects
		return evalIntegerInfixExpression(operator, left, right)

	// Unsupported combination
	default:
		// Return null
		return NULL
	}
}

// A function that evaluates an infix expression between two Integers
// given a infix operator and the left and right Integers objects
func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	// Retrieve the left and right integer values
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	// Check the type of operator
	switch operator {

	// Plus operator (Add)
	case "+":
		// Evaluate the object for addition
		return &object.Integer{Value: leftVal + rightVal}

	// Minus Operator (Subtract)
	case "-":
		// Evaluate the object for subtraction
		return &object.Integer{Value: leftVal - rightVal}

	// Asterisk Operator (Multiply)
	case "*":
		// Evaluate the object for multiplication
		return &object.Integer{Value: leftVal * rightVal}

	// Slash Operator (Divide)
	case "/":
		// Evaluate the object for division
		return &object.Integer{Value: leftVal / rightVal}

	// Unsupported Operator
	default:
		// Return null
		return NULL
	}
}
