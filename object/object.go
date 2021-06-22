package object

import "fmt"

const (
	NULL_OBJ = "NULL"

	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"

	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

// A type alias that represents the type of an object
type ObjectType string

// A structure that represents an evaluated object
type Object interface {
	Type() ObjectType
	Inspect() string
}

// A structure that represents a Returned object
type ReturnValue struct {
	// Represents the returned object
	Value Object
}

// A method of ReturnValue that returns the Return value type
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// A method of ReturnValue that returns the string value of the Returned object
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// A structure that represents an Error object
type Error struct {
	// Represents the error message
	Message string
}

// A method of Error that returns the Error value type
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// A method of Error that returns the string value of the Error object
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// A constructor function that generates and returns a new Error
// for a given message and some variadic interface
func NewError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
