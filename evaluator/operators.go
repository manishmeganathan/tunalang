package evaluator

import "github.com/manishmeganathan/tuna/object"

// A function that returns the result object for a
// given object with the prefix bang operator applied
func evalBangOperatorExpression(right object.Object) object.Object {
	// Check value of object
	switch right {

	// Flip true to false
	case TRUE:
		return FALSE

	// Flip false to true
	case FALSE:
		return TRUE

	// Flip null to true
	case NULL:
		return TRUE

	// Default to false
	default:
		return FALSE
	}
}

// A function that returns the result object for a
// given object with the prefix minus operator applied
func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	// Check that object is an Integer
	if right.Type() != object.INTEGER_OBJ {
		// Return null for non integer objects
		return NULL
	}

	// Retrieve the value of the Integer object
	value := right.(*object.Integer).Value
	// Return the modified Integer with the negative of the value
	return &object.Integer{Value: -value}
}
