package object

// A structure that represents an execution environment
type Environment struct {
	// Represents the memory pool of stored objects
	store map[string]Object
}

// A constructor function that generates
// and returns a new Environment
func NewEnvironment() *Environment {
	// Initialize the store hash map
	s := make(map[string]Object)
	// Return the environment
	return &Environment{store: s}
}

// A method of Environment to retrieve a value from the store
func (e *Environment) Get(name string) (Object, bool) {
	// Retrieve the value from the store
	obj, ok := e.store[name]
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
