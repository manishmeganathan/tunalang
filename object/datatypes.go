package object

import (
	"fmt"
	"hash/fnv"
)

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

// A method of Integer that return the HashKey of the object
func (i *Integer) HashKey() HashKey {
	// Create and return the HashKey object from the integer value
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// A structure that represents a Boolean object
type Boolean struct {
	// Represents the value of the Boolean
	Value bool
}

// A method of Boolean that returns the Boolean value type
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// A method of Boolean that returns the string value of the Boolean
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// A method of Boolean that returns the HashKey of the object
func (b *Boolean) HashKey() HashKey {
	// Declare an unsigned int64
	var value uint64

	if b.Value {
		// If the Boolean is true, set the value to 1
		value = 1

	} else {
		// Else set the value to 0
		value = 0
	}

	// Create and return the HashKey object
	return HashKey{Type: b.Type(), Value: value}
}

// A structure that represents a String object
type String struct {
	// Represents the value of the String
	Value string
}

// A method of String that returns the String value type
func (s *String) Type() ObjectType { return STRING_OBJ }

// A method of String that returns the string value of the String
func (s *String) Inspect() string { return s.Value }

// A method of String that returns the HashKey of the object
func (s *String) HashKey() HashKey {
	// Create new 64bit FNV hasher
	h := fnv.New64a()
	// Write the string value to the hasher
	h.Write([]byte(s.Value))

	// Return the HashKey object
	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
