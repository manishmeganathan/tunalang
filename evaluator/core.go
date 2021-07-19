package evaluator

import (
	"github.com/manishmeganathan/tunalang/object"
	"github.com/manishmeganathan/tunalang/syntaxtree"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// A function that evaluates a Syntax Tree given a node
// on it and returns an evaluated object
func Evaluate(node syntaxtree.Node, env *object.Environment) object.Object {
	// Check the type of Syntax Tree Node
	switch node := node.(type) {
	// Program Node (Tree Root)
	case *syntaxtree.Program:
		// Evaluate the statements in the program
		return evalProgram(node, env)

	// Return Statement Node
	case *syntaxtree.ReturnStatement:
		// Evaluate the Expression in the return statement
		val := Evaluate(node.ReturnValue, env)
		// Check if evaluated value is an error
		if isError(val) {
			// Return the error
			return val
		}

		// Return the evaluated return object
		return &object.ReturnValue{Value: val}

		// Let Statement Node
	case *syntaxtree.LetStatement:
		// Evaluate the Expression in the let statement
		val := Evaluate(node.Value, env)
		// Check if evaluated value is an error
		if isError(val) {
			// Return the error
			return val
		}

		// Set the evaluated object and the literal
		// name to the environment store
		env.Set(node.Name.Value, val)

	// Expression Node
	case *syntaxtree.ExpressionStatement:
		// Recursive evaluation
		return Evaluate(node.Expression, env)

	// Prefix Expression Node
	case *syntaxtree.PrefixExpression:
		// Evaluate the expression into an object
		right := Evaluate(node.Right, env)
		// Check if evaluated value is an error
		if isError(right) {
			// Return the error
			return right
		}

		// Evaluate the object for the operator
		return evalPrefixExpression(node.Operator, right)

	// Infix Expression Node
	case *syntaxtree.InfixExpression:
		// Evaluate the left node
		left := Evaluate(node.Left, env)
		// Check if evaluated left value is an error
		if isError(left) {
			// Return the error
			return left
		}

		// Evaluate the right node
		right := Evaluate(node.Right, env)
		// Check if evaluated right value is an error
		if isError(right) {
			// Return the error
			return right
		}

		// Evaluate the expression with the objects and the operator
		return evalInfixExpression(node.Operator, left, right)

	// Block Statement Node
	case *syntaxtree.BlockStatement:
		// Evaluate the statements in the block
		return evalBlockStatement(node, env)

	// If Expression Node
	case *syntaxtree.IfExpression:
		// Evaluate the if expression
		return evalIfExpression(node, env)

	// Call Expression Node
	case *syntaxtree.CallExpression:
		// Evaluate the function
		function := Evaluate(node.Function, env)
		// Check if the evaluated value is an error
		if isError(function) {
			// Return the error
			return function
		}

		// Evaluate the function arguments
		args := evalExpressions(node.Arguments, env)
		// Check for errors
		if len(args) == 1 && isError(args[0]) {
			// Return the error
			return args[0]
		}

		// Evaluate the function call
		return applyFunction(function, args)

	// Identifier Expression Node
	case *syntaxtree.IndexExpression:
		// Evaluate the left expression
		left := Evaluate(node.Left, env)
		// Check if evaluated value is an error
		if isError(left) {
			// Return the error
			return left
		}

		// Evaluate the index expression
		index := Evaluate(node.Index, env)
		// Check if evaluated value is an error
		if isError(index) {
			// Return the error
			return index
		}

		// Evaluate the index expression
		return evalIndexExpression(left, index)

	// List Literal Node
	case *syntaxtree.ListLiteral:
		// Evaluate the list literal elements
		elements := evalExpressions(node.Elements, env)
		// Check for errors
		if len(elements) == 1 && isError(elements[0]) {
			// Return the error
			return elements[0]
		}

		// Return the List Object
		return &object.List{Elements: elements}

	// Map Literal Node
	case *syntaxtree.MapLiteral:
		// Evaluate the map literal
		return evalMapLiteral(node, env)

	// Function Literal Node
	case *syntaxtree.FunctionLiteral:
		// Return the Function Object
		return &object.Function{Parameters: node.Parameters, Env: env, Body: node.Body}

	// Identifier Literal Node
	case *syntaxtree.Identifier:
		// Evaluate the identifier
		return evalIdentifier(node, env)

	// Integer Literal Node
	case *syntaxtree.IntegerLiteral:
		// Return the Integer Object
		return &object.Integer{Value: node.Value}

	// Boolean Literal Node
	case *syntaxtree.BooleanLiteral:
		// Return the native Boolean Object for the value
		return getNativeBoolean(node.Value)

	// String Literal Node
	case *syntaxtree.StringLiteral:
		// Return the String Object
		return &object.String{Value: node.Value}
	}

	// Return nil if not evaluated
	return nil
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

// A function that returns whether an Object is an Error
func isError(obj object.Object) bool {
	// Check if object is non null
	if obj != nil {
		// Check the object type for Error
		return obj.Type() == object.ERROR_OBJ
	}

	// Return false (null object)
	return false
}

// A function that applies a given function object on a slice of object arguments
func applyFunction(fn object.Object, args []object.Object) object.Object {

	switch fn := fn.(type) {

	case *object.Function:
		// Create the function's extended environment
		extendedEnv := extendFunctionEnv(fn, args)
		// Evaluate the function body
		evaluated := Evaluate(fn.Body, extendedEnv)
		// Return the unwrapped value
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		// Call the built-in function with the args
		return fn.Fn(args...)

	default:
		// Return an Error
		return object.NewError("not a function: %s", fn.Type())
	}
}

// A function that creates an extended environment for a function
func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	// Create a new enclosed enivronment
	env := object.NewEnclosedEnvironment(fn.Env)
	// Iterate over the function args
	for paramIdx, param := range fn.Parameters {
		// Add the function arg to the enclosed environment
		env.Set(param.Value, args[paramIdx])
	}

	// Return the extended environment
	return env
}

// A function that unwraps an object into its value if it is a Return Object
func unwrapReturnValue(obj object.Object) object.Object {
	// Check if the given object is a Return Object
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		// Return the return value
		return returnValue.Value
	}

	// Return the object back
	return obj
}
