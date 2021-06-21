package object

import "fmt"

// A structure that represents a Null object
type Null struct{}

// A method of Null that returns the Null value type
func (n *Null) Type() ObjectType { return NULL_OBJ }

// A method of Null that returns the string value of the Null
func (n *Null) Inspect() string { return "null" }

// A structure that represents an Integer object
type Integer struct {
	// Represents the value of the Integer
	Value int64
}

// A method of Integer that returns the Integer value type
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// A method of Integer that returns the string value of the Integer
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// A structure that represents a Boolean object
type Boolean struct {
	// Represents the value of the Boolean
	Value bool
}

// A method of Boolean that returns the Boolean value type
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// A method of Boolean that returns the string value of the Boolean
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
