package evaluator

import (
	"fmt"

	"github.com/manishmeganathan/tuna/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			// List objects
			case *object.List:
				return &object.Integer{Value: int64(len(arg.Elements))}

			// String objects
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}

			// Everything else
			default:
				return object.NewError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.LIST_OBJ {
				return object.NewError("argument to `first` must be LIST, got %s",
					args[0].Type())
			}

			list := args[0].(*object.List)
			if len(list.Elements) > 0 {
				return list.Elements[0]
			}

			return NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.LIST_OBJ {
				return object.NewError("argument to `last` must be LIST, got %s",
					args[0].Type())
			}

			list := args[0].(*object.List)
			length := len(list.Elements)
			if len(list.Elements) > 0 {
				return list.Elements[length-1]
			}

			return NULL
		},
	},
	"tail": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.LIST_OBJ {
				return object.NewError("argument to `tail` must be LIST, got %s",
					args[0].Type())
			}

			list := args[0].(*object.List)
			length := len(list.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1)
				copy(newElements, list.Elements[1:length])
				return &object.List{Elements: newElements}
			}

			return NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 2 {
				return object.NewError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			if args[0].Type() != object.LIST_OBJ {
				return object.NewError("argument to `push` must be LIST, got %s",
					args[0].Type())
			}

			list := args[0].(*object.List)
			length := len(list.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, list.Elements)
			newElements[length] = args[1]

			return &object.List{Elements: newElements}
		},
	},
}
