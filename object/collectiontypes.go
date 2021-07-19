package object

import (
	"bytes"
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
