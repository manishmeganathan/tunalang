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
