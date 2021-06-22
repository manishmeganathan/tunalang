package evaluator

import (
	"github.com/manishmeganathan/tuna/object"
	"github.com/manishmeganathan/tuna/syntaxtree"
)

// A function that evaluates a Syntax tree program into an evaluated object
func evalProgram(program *syntaxtree.Program, env *object.Environment) object.Object {
	// Declare an object
	var result object.Object

	// Iterate over the program statements
	for _, statement := range program.Statements {
		// Update the result object
		result = Evaluate(statement, env)

		// Check the type of evaluated object
		switch result := result.(type) {

		// Return Object
		case *object.ReturnValue:
			// Return the return value
			return result.Value

		// Error Object
		case *object.Error:
			// Return the error object
			return result
		}
	}

	// Return the result object
	return result
}

// A function that evaluates a Syntax tree block into an evaluated object
func evalBlockStatement(block *syntaxtree.BlockStatement, env *object.Environment) object.Object {
	// Declare an object
	var result object.Object

	// Iterate over the block statements
	for _, statement := range block.Statements {
		// Update the result object
		result = Evaluate(statement, env)

		// Check if result has evaluated object
		if result != nil {
			// Retrieve the object type
			rt := result.Type()

			// Check if the object type is either a Return or an Error
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				// Return the object
				return result
			}
		}
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
		// Return Error
		return object.NewError("unsupported operator: %s%s", operator, right.Type())
	}
}

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
		// Return Error for non integer objects
		return object.NewError("unsupported operator: -%s", right.Type())

	}

	// Retrieve the value of the Integer object
	value := right.(*object.Integer).Value
	// Return the modified Integer with the negative of the value
	return &object.Integer{Value: -value}
}

// A function that evaluates an infix expression given
// a infix operator and the left and right objects
func evalInfixExpression(operator string, left, right object.Object) object.Object {
	// Check Parameters
	switch {

	// If both are objects are Integers
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		// Evaluate expression for integer objects
		return evalIntegerInfixExpression(operator, left, right)

	// If both objects are not Integers but the operator is '=='
	case operator == "==":
		// Evaluate the objects for '=='
		return getNativeBoolean(left == right)

	// If both objects are not Integers but the operator is '!='
	case operator == "!=":
		// Evaluate the objects for '!='
		return getNativeBoolean(left != right)

	// If both objects are not of the same type
	case left.Type() != right.Type():
		// Return Error
		return object.NewError("type mismatch: %s %s %s", left.Type(), operator, right.Type())

	// Unsupported combination
	default:
		// Return Error
		return object.NewError("unsupported operator: %s %s %s", left.Type(), operator, right.Type())
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
		// Evaluate the objects for addition
		return &object.Integer{Value: leftVal + rightVal}

	// Minus Operator (Subtract)
	case "-":
		// Evaluate the objects for subtraction
		return &object.Integer{Value: leftVal - rightVal}

	// Asterisk Operator (Multiply)
	case "*":
		// Evaluate the objects for multiplication
		return &object.Integer{Value: leftVal * rightVal}

	// Slash Operator (Divide)
	case "/":
		// Evaluate the objects for division
		return &object.Integer{Value: leftVal / rightVal}

	// Less Than Operator
	case "<":
		// Evaluate the objects for '<'
		return getNativeBoolean(leftVal < rightVal)

	// Greater Than Operator
	case ">":
		// Evaluate the objects for '>'
		return getNativeBoolean(leftVal > rightVal)

	// Equal To Operator
	case "==":
		// Evaluate the objects for '=='
		return getNativeBoolean(leftVal == rightVal)

	// Not Equal To Operator
	case "!=":
		// Evaluate the objects for '!='
		return getNativeBoolean(leftVal != rightVal)

	// Unsupported Operator
	default:
		// Return Error
		return object.NewError("unsupported operator: %s %s %s", left.Type(), operator, right.Type())
	}
}

// A function that evaluates an if expression given an IfExpression syntax tree node
func evalIfExpression(ie *syntaxtree.IfExpression, env *object.Environment) object.Object {
	// Evaluate the conditional statement
	condition := Evaluate(ie.Condition, env)
	// Check if evaluated condition is an error
	if isError(condition) {
		// Return the error
		return condition
	}

	// Check if the condition is truthy
	if isTruthy(condition) {
		// Evaluate the consequence block
		return Evaluate(ie.Consequence, env)

		// Check if alternate exists
	} else if ie.Alternative != nil {
		// Evaluate the alternate consequence block
		return Evaluate(ie.Alternative, env)

	} else {
		// Return null
		return NULL
	}
}
