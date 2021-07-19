package object

import (
	"bytes"
	"fmt"
	"strings"
)

// A structure that represents a List object
type List struct {
	// Represents the elements of the list
	Elements []Object
}

// A method of List that returns the List value type
func (l *List) Type() ObjectType { return LIST_OBJ }

// A method of List that returns the string value of the List
func (l *List) Inspect() string {
	// Create a string buffer
	var out bytes.Buffer

	// Declare a slice to accumulate the elements
	elements := []string{}
	// Iterate through the elements
	for _, e := range l.Elements {
		// Append the string representation of the element to the slice
		elements = append(elements, e.Inspect())
	}

	// Join the elements with a comma
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	// Return the string representation
	return out.String()
}

// A structure that represents a Map key-value pair
type MapPair struct {
	// Represents the key of the key-value pair
	Key Object
	// Represents the value of the key-value pair
	Value Object
}

// A structure that represents a Map object
type Map struct {
	// Represents the key-value pairs of the map
	Pairs map[HashKey]MapPair
}

// A method of Map that returns the Map value type
func (h *Map) Type() ObjectType { return MAP_OBJ }

// A method of Map that returns the string value of the Map
func (h *Map) Inspect() string {
	// Create a string buffer
	var out bytes.Buffer

	// Declare a slice to accumulate the key value pairs
	pairs := []string{}
	// Iterate through the key value pairs
	for _, pair := range h.Pairs {
		// Append the string representation of the pair to the slice
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	// Join the key value pairs with a comma
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	// Return the string representation
	return out.String()
}
