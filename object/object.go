package object

const (
	NULL_OBJ = "NULL"

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
