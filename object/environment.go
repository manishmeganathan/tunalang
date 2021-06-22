package object

// A structure that represents an execution environment
type Environment struct {
	// Represents the memory pool of stored objects
	store map[string]Object
	// Represents the outer environment scope
	outer *Environment
}

// A constructor function that generates
// and returns a new Environment
func NewEnvironment() *Environment {
	// Initialize the store hash map
	s := make(map[string]Object)
	// Return the environment
	return &Environment{store: s, outer: nil}
}

// A constructor function that generates and returns
// an enclosed Environment given the outer environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	// Create a new environment
	env := NewEnvironment()
	// Assign its outer field to the given outer scope
	env.outer = outer
	// Return the new enclosed environment
	return env
}

// A method of Environment to retrieve a value from the store
func (e *Environment) Get(name string) (Object, bool) {
	// Retrieve the value from the store
	obj, ok := e.store[name]

	// Check if an outer environment exists
	// if the value is not found
	if !ok && e.outer != nil {
		// Retrive the value from the outer environment
		obj, ok = e.outer.Get(name)
	}

	// Return the object and the ok flag
	return obj, ok
}

// A method of Environment to add a value to store
func (e *Environment) Set(name string, val Object) Object {
	// Add the value to the store
	e.store[name] = val
	// Return the value as an acknowledgement
	return val
}
