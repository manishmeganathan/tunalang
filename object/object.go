package object

const (
	NULL_OBJ = "NULL"

	RETURN_VALUE_OBJ = "RETURN_VALUE"

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
