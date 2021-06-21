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

// A function that evaluates prefix expression
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
